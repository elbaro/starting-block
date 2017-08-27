package components

import (
	"errors"
	"runtime"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var terminalColor = func() Component {
	check := func() (bool, string) {
		// defaults read com.googlecode.iterm2 Custom\ Color\ Presets | grep Snazzy
		return false, ""
	}

	install := func() error {
		switch runtime.GOOS {
		case linux:
			// source snazzy.gnome
		case mac:
			sh("https://github.com/sindresorhus/iterm2-snazzy/raw/master/Snazzy.itermcolors")
			sh("open Snazzy.itermcolors")
		}
		return errors.New("not implemented yet")
	}

	return Component{"terminal-color", []string{linux, mac}, check, install}
}()
