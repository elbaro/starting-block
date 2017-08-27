package components

import (
	"fmt"
	"runtime"
)

// Zshrc installs `~/.zshrc` including slimzsh.
var atom = func() Component {
	check := func() (bool, string) {
		return false, ""
	}

	install := func() error {
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
		return nil
	}

	return Component{"atom", []string{linux, mac, win}, check, install}
}()
