package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"

	"main/blob"

	"github.com/andlabs/ui"
	"github.com/progrium/go-shell"
)

type checker func() (bool, string)
type installer func() error
type component struct {
	checker
	installer
}

var (
	sh = shell.Run
)

var usr *user.User

func isGPU() bool {
	return binaryExist("nvidia-smi")
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func binaryExist(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func grepCheck(path string, key string) bool {
	// fmt.Println(path, key)
	err := exec.Command("grep", key, path).Run()
	return err == nil
}

func sudoCheck() bool {
	return os.Getenv("SUDO_USER") != ""
}

var aptUpdated = false

func packageInstall(name string) {
	if !aptUpdated {
		sh("apt-get update")
		aptUpdated = true
	}
	sh("apt-get install -y " + name)
}

func openURL(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// https://api.github.com/repos/atom/atom/releases/latest
func latestRelease(repo string) error {
	urlPattern := map[string]string{
		"linux":   "https[^\"]*atom-amd64\\.deb",
		"darwin":  "https[^\"]*atom-mac\\.zip",
		"windows": "https[^\"]*AtomSetup-x64\\.exe",
	}
	res, err := http.Get("https://api.github.com/repos/" + repo + "/releases/latest")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	url := regexp.MustCompile(urlPattern[runtime.GOOS]).FindString(string(body))
	openURL(url)
	return nil
}

func gitconfigCheck() (bool, string) {
	path := filepath.Join(usr.HomeDir, ".gitconfig")
	if binaryExist("git") && grepCheck(path, "lg1") {
		return true, "installed"
	}
	return true, "install"
}

func gitconfigInstall() error {
	path := filepath.Join(usr.HomeDir, ".gitconfig")
	err := ioutil.WriteFile(path, []byte(blob.Gitconfig), 0644)
	return err
}

var gitconfig = component{gitconfigCheck, gitconfigInstall}

func sshkeyCheck() (bool, string) {
	path := filepath.Join(usr.HomeDir, ".ssh", "id_rsa.pub")
	fmt.Println(path)
	if pathExist(path) {
		return false, "installed"
	}
	return true, "install"
}

func sshkeyInstall() error {
	err := shell.Cmd("yes", "y").Pipe("ssh-keygen").ErrFn()()
	return err
}

var sshkey = component{sshkeyCheck, sshkeyInstall}

func atom() {
	switch runtime.GOOS {
	case "linux":
		openURL("https://atom.io/download/deb")
	case "windows":
		openURL("https://atom.io/download/windows_x64")
	case "darwin":
		openURL("https://atom.io/download/mac")
	default:
		_ = fmt.Errorf("unsupported platform")
	}
}

func slimzsh() error {
	dir, err := ioutil.TempDir("", "slimzsh")
	if err != nil {
		return err
	}
	path := filepath.Join(usr.HomeDir, ".slimzsh")
	if !pathExist(path) {
		sh("git clone --recursive https://github.com/changs/slimzsh.git " + dir)
		sh("mv " + dir + " ~/.slimzsh")
		sh("chmod g-w ~/.slimzsh")
	}
	return nil
}

func zshrcCheck() (bool, string) {
	path := filepath.Join(usr.HomeDir, ".zshrc")
	if !binaryExist("zsh") {
		return false, "no zsh"
	}
	if grepCheck(path, "alias t=") {
		return true, "installed"
	}
	return true, "install"
}

func zshrcInstall() error {
	slimzsh()
	path := filepath.Join(usr.HomeDir, ".zshrc")
	_err := ioutil.WriteFile(path, []byte(blob.Zshrc), 0644)
	return _err
}

var zshrc = component{zshrcCheck, zshrcInstall}

func tmuxCheck() (bool, string) {
	if binaryExist("tmux") {
		return false, "installed"
	}
	return true, "install"
}

func tmuxInstall() error {
	packageInstall("tmux")
	return nil
}

var tmux = component{tmuxCheck, tmuxInstall}

func tmuxconfCheck() (bool, string) {
	if !binaryExist("tmux") {
		return false, "no tmux"
	}
	path := filepath.Join(usr.HomeDir, ".tmux.conf")
	if grepCheck(path, "cpu_percentage") {
		return true, "installed"
	}
	return true, "install"
}

func tmuxconfInstall() error {
	path := filepath.Join(usr.HomeDir, ".tmux.conf")

	sh("git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm")
	sh("git clone https://github.com/tmux-plugins/tmux-sensible ~/.tmux/plugins/tmux-sensible")
	sh("git clone https://github.com/tmux-plugins/tmux-cpu ~/.tmux/plugins/tmux-cpu")

	if isGPU() {
		sh("git clone https://gitlab.com/elbaro/tmux-gpu.git ~/.tmux/plugins/tmux-gpu")
		err := ioutil.WriteFile(path, []byte(blob.TmuxGPU), 0644)
		return err
	} else {
		sh("git clone https://gitlab.com/elbaro/tmux-gpu.git ~/.tmux/plugins/tmux-gpu")
		err := ioutil.WriteFile(path, []byte(blob.Tmux), 0644)
		return err
	}
}

var tmuxconf = component{tmuxconfCheck, tmuxconfInstall}

func zshCheck() (bool, string) {
	if binaryExist("zsh") {
		return false, "installed"
	}
	return true, "install"
}

func zshInstall() error {
	switch runtime.GOOS {
	case "linux":
		sh("apt-get install -y zsh")
		// sh("chsh -s ") TODO: run as user
	case "windows":
		openURL("https://atom.io/download/windows_x64")
	case "darwin":
		sh("yes | brew install zsh")
	default:
		return errors.New("unsupported platform")
	}
	return nil
}

var zsh = component{zshCheck, zshInstall}

var components = map[string](map[string]component){
	"darwin": map[string]component{
		".gitconfig": gitconfig,
		"sshkey":     sshkey,
		".zshrc":     zshrc,
		"zsh":        zsh,
		"tmux":       tmux,
		".tmux.conf": tmuxconf,
	},
	"linux": map[string]component{
		".gitconfig": gitconfig,
		"sshkey":     sshkey,
		".zshrc":     zshrc,
		"zsh":        zsh,
		"tmux":       tmux,
		".tmux.conf": tmuxconf,
	},
}

func lazyButton(comp component) *ui.Button {
	button := ui.NewButton("")
	go func() {
		installable, msg := comp.checker()
		button.SetText(msg)
		if installable {
			button.OnClicked(func(*ui.Button) {
				comp.installer()
				// panic(err)
			})
			button.Enable()
		} else {
			button.Disable()
		}
	}()

	return button
}

func keys(_map map[string]component) []string {
	keys := make([]string, len(_map))
	i := 0
	for key := range _map {
		keys[i] = key
		i++
	}
	return keys
}

func main() {
	defer shell.ErrExit()
	shell.Trace = true // like set +x
	usr, _ = user.Current()
	// user.Username
	// user.HomeDir
	// latestRelease("atom/atom")
	//

	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		_components := components[runtime.GOOS]
		for _, name := range keys(_components) {
			box.Append(ui.NewLabel(name), false)
			box.Append(lazyButton(_components[name]), false)
		}

		window := ui.NewWindow("Hello", 500, 600, false)
		window.SetChild(box)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
