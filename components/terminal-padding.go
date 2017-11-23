package components

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var terminalPadding = func() Component {
	gtkCss := `vte-terminal {
    padding: 30px 30px 30px 30px;
}
	`
	check := func() (bool, string) {
		switch runtime.GOOS {
		case linux:
			path := filepath.Join(usr.HomeDir, ".config/gtk-3.0/gtk.css")
			return pathExist(path), ""
		case mac:
			return false, "use iTerm2 nightly build"
		}
		return false, "You shouldn't see this"
	}

	install := func() error {
		switch runtime.GOOS {
		case linux:
			path := filepath.Join(usr.HomeDir, ".config/gtk-3.0/gtk.css")
			err := ioutil.WriteFile(path, []byte(gtkCss), 0644)
			return err
		case mac:
			// mdls -name kMDItemVersion /Applications/iTerm.app
			// kMDItemVersion = "3.1.20170825-nightly"
		}
		return errors.New("You shouldn't see this")
	}

	return Component{"terminal-padding", []string{linux, mac}, check, install}
}()
