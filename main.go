package main

import (
	"fmt"
	"os"
	"runtime"

	zoo "github.com/elbaro/starting-block/components"
	"github.com/progrium/go-shell"
)

func main() {
	defer shell.ErrExit()
	shell.Trace = true // like set +x
	// usr, _ := user.Current()
	// user.Username
	// user.HomeDir
	// latestRelease("atom/atom")

	components := []zoo.Component{}
	for _, component := range zoo.Components {
		for _, platform := range component.Platforms {
			if platform == runtime.GOOS {
				components = append(components, component)
			}
		}
	}
	installed := []string{}
	available := []string{}
	unavailable := [][2]string{}
	for _, component := range components {
		is_installed, reason := component.Checker()
		if is_installed {
			installed = append(installed, component.Name)
		} else if reason == "" {
			available = append(available, component.Name)
		} else {
			unavailable = append(unavailable, [2]string{component.Name, reason})
		}
	}

	var mode string
	if len(os.Args) < 2 {
		mode = ""
	} else {
		mode = os.Args[1]
	}

	switch mode {
	case "install":
		for _, name := range os.Args[2:] {
			found := false
			for _, c := range append(append([]string{}, installed...), available...) {
				if name == c {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("설치가 불가능합니다: " + name)
				continue
			}
			for _, c := range zoo.Components {
				if name == c.Name {
					err := c.Install()
					if err != nil {
						fmt.Println("설치에 실패했습니다.: " + name)
					} else {
						fmt.Println("설치에 성공했습니다.: " + name)
					}
					break
				}

			}
		}
	case "remove":
		for _, name := range os.Args[2:] {
			found := false
			for _, c := range append(append([]string{}, installed...), available...) {
				if name == c {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("제거가 불가능합니다: " + name)
				continue
			}
			for _, c := range zoo.Components {
				if name == c.Name {
					// err := c.Install()
					// if err != nil {
					// 	fmt.Println("설치에 실패했습니다.: " + name)
					// } else {
					// 	fmt.Println("설치에 성공했습니다.: " + name)
					// }
					fmt.Println("어떻게 제거하는지 모릅니다: " + name)
					break
				}
			}
		}
	default: // list
		fmt.Printf("%s [install|remove] component1 component2 ..\n\n", os.Args[0])
		fmt.Println("Installed:")
		for _, c := range installed {
			fmt.Printf("    %s\n", c)
		}
		fmt.Println()
		fmt.Println("Available:")
		for _, c := range available {
			fmt.Printf("    %s\n", c)
		}
		fmt.Println()
		fmt.Println("Unavailable:")
		for _, c := range unavailable {
			fmt.Printf("    %-10s   %s\n", c[0], c[1])
		}
	}
}
