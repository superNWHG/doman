package flags

import "github.com/superNWHG/doman/internal/data"

func addData(path string, name string, newPath string) error {
	path = checkForSlash(path)

	path = path + "/dotfiles.json"

	if name == "" {
		found := false
		for i := len(newPath); !found; i-- {
			if string(newPath[i-1]) == "/" {
				name = newPath[i:]
				found = true
			}
		}

	}

	nameSlice := []string{name}
	newPathSlice := []string{newPath}

	if err := data.NewData(path, nameSlice, newPathSlice); err != nil {
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
