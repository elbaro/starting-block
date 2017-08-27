package components

import (
	"fmt"
	"path/filepath"

	shell "github.com/progrium/go-shell"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var sshkey = func() Component {
	check := func() (bool, string) {
		path := filepath.Join(usr.HomeDir, ".ssh", "id_rsa.pub")
		fmt.Println(path)
		if pathExist(path) {
			return true, ""
		}
		return false, ""
	}

	install := func() error {
		err := shell.Cmd("yes", "y").Pipe("ssh-keygen").ErrFn()()
		return err
	}

	return Component{"sshkey", []string{linux, mac, win}, check, install}
}()
