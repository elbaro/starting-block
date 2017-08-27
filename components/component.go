package components

// Component is a struct of (platforms, checker, install)
type Component struct {
	Name      string
	Platforms []string
	Checker   (func() (bool, string)) // return (installed?, hindrance)
	Install   (func() error)
}

func init() {

}

const mac = "darwin"
const win = "windows"
const linux = "linux"

// Components is a list of all available components
var Components = []Component{atom, gitconfig, sshkey, tmux, tmuxconf, zsh, zshrc, terminalColor, terminalPadding}
