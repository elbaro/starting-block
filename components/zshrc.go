package components

import (
	"io/ioutil"
	"path/filepath"

	"github.com/elbaro/starting-block/blob"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var zshrc = func() Component {
	slimzsh := func() error {
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

	check := func() (bool, string) {
		path := filepath.Join(usr.HomeDir, ".zshrc")
		if !binaryExist("zsh") {
			return false, "no zsh"
		}
		if grepCheck(path, "alias t=") {
			return true, ""
		}
		return false, ""
	}

	install := func() error {
		slimzsh()
		path := filepath.Join(usr.HomeDir, ".zshrc")
		_err := ioutil.WriteFile(path, []byte(blob.Zshrc), 0644)
		return _err
	}

	return Component{".zshrc", []string{linux, mac}, check, install}
}()
