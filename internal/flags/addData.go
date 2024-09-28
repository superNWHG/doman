package flags

import (
	"github.com/superNWHG/doman/cmd"
	"github.com/superNWHG/doman/internal/data"
)

func addData(path string, name string, newPath string) error {
	path = checkForSlash(path)

	if name == "" {
		for i := len(newPath); i > 0; i-- {
			if string(newPath[i-1]) == "/" {
				name = newPath[i:]
				i = 0
			}
		}

	}

	namePath := path + "/" + name

	path = path + "/dotfiles.json"

	nameSlice := []string{name}
	namePathSlice := []string{namePath}
	newPathSlice := []string{newPath}

	if err := data.NewData(path, nameSlice, newPathSlice); err != nil {
		return err
	}

	if err := cmd.NewLink(newPathSlice, namePathSlice); err != nil {
		return err
	}

	return nil
}

func checkForSlash(slashString string) string {
	if slashString[len(slashString)-1:] == "/" {
		slashString = slashString[:len(slashString)-1]
	}

	return slashString
}
