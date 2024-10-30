package flags

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/superNWHG/doman/internal/data"
	"github.com/superNWHG/doman/internal/packages"
)

func install(path string, installNames []string, os string) error {
	path = filepath.Join(path, "dotfiles.json")
	names, _, _, err := data.ReadDataFile(path)
	if err != nil {
		return err
	}

	if installNames == nil {
		installNames = names

		for i := range installNames {
			if strings.Contains(installNames[i], "/") {
				for x, character := range installNames[i] {
					if string(character) == "/" {
						installNames[i] = installNames[i][:x]
					}
				}
			}
		}
	}

	for i := range installNames {
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

	return nil
}
