package components

import (
	"errors"
	"runtime"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var zsh = func() Component {

	check := func() (bool, string) {
		if binaryExist("zsh") {
			return true, ""
		}
		return false, ""
	}

	install := func() error {
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

	return Component{"zsh", []string{linux, mac}, check, install}
}()
