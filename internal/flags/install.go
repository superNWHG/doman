package flags

import (
	"fmt"
	"path/filepath"

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
	}

	for i := range installNames {
		fmt.Print("Please select a package to install for ", installNames[i], "\nPress [ENTER] to continue")
		var input string
		fmt.Scanln(&input)
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
