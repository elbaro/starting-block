package components

import (
	"io/ioutil"
	"path/filepath"

	"github.com/elbaro/starting-block/blob"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var gitconfig = func() Component {
	check := func() (bool, string) {
		path := filepath.Join(usr.HomeDir, ".gitconfig")
		if binaryExist("git") && grepCheck(path, "lg1") {
			return true, ""
		}
		return false, ""
	}

	install := func() error {
		path := filepath.Join(usr.HomeDir, ".gitconfig")
		err := ioutil.WriteFile(path, []byte(blob.Gitconfig), 0644)
		return err
	}

	return Component{".gitconfig", []string{linux, mac, win}, check, install}
}()
