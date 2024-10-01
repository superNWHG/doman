package flags

import (
	"fmt"

	"github.com/superNWHG/doman/internal/data"
)

func readData(path string) error {
	path = checkForSlash(path) + "/dotfiles.json"
	values, _, entries, err := data.ReadDataFile(path)
	if err != nil {
		return err
	}

	for i := range values {
		fmt.Println(values[i]+":", entries[values[i]])
	}

	return nil
}
