package flags

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/superNWHG/doman/internal/data"
	"github.com/superNWHG/doman/internal/packages"
)

func install(path string, installNames []string, os string, lastPathPart bool) error {
	path = filepath.Join(path, "dotfiles.json")
	names, _, _, err := data.ReadDataFile(path)
	if err != nil {
		return err
	}

	if installNames == nil {
		installNames = names

		for i := range installNames {
			if strings.Contains(installNames[i], "/") {
				if lastPathPart {
					for x := len(installNames[i]) - 1; x >= 0; x-- {
						if string(installNames[i][x]) == "/" {
							installNames[i] = installNames[i][x+1:]
							break
						}
					}
				} else {
					for x, character := range installNames[i] {
						if string(character) == "/" {
							installNames[i] = installNames[i][:x]
						}
					}
				}
			}
			fmt.Print("Optionally, you can specify a custom package name for ", installNames[i], " (leave empty to use default):")
			var input string
			fmt.Scanln(&input)
			if input != "" {
				installNames[i] = input
			}
			pkg, err := packages.Search(os, installNames[i])
			if err != nil {
				return err
			}
			if err := packages.Install(os, pkg); err != nil {
				return err
			}
		}
	}

	return nil
}
