package components

import (
	"io/ioutil"
	"path/filepath"

	"github.com/elbaro/starting-block/blob"
)

var tmuxconf = func() Component {
	check := func() (bool, string) {
		if !binaryExist("tmux") {
			return false, "no tmux"
		}
		path := filepath.Join(usr.HomeDir, ".tmux.conf")
		if grepCheck(path, "cpu_percentage") {
			return true, ""
		}
		return false, ""
	}

	install := func() error {
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
	return Component{".tmux.conf", []string{linux, mac}, check, install}
}()
