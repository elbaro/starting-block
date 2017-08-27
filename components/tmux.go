package components

// Tmux installs `tmux`.
var tmux = func() Component {
	check := func() (bool, string) {
		if binaryExist("tmux") {
			return true, ""
		}
		return false, ""
	}

	install := func() error {
		packageInstall("tmux")
		return nil
	}

	return Component{"tmux", []string{linux, mac}, check, install}
}()
