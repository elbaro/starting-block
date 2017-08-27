package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"

	shell "github.com/progrium/go-shell"
)

var usr = func() *user.User {
	ret, _ := user.Current()
	return ret
}()

var (
	sh = shell.Run
)

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

func brewInstall(recipe string) error {
	// "brew install " + recipe
	return nil
}
