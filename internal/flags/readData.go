package flags

import (
	"fmt"
	"path/filepath"

	"github.com/superNWHG/doman/internal/data"
)

func readData(path string) error {
	path = filepath.Join(path, "dotfiles.json")
	values, _, entries, err := data.ReadDataFile(path)
	if err != nil {
		return err
	}

	for i := range values {
		fmt.Println(values[i]+":", entries[values[i]])
	}

	return nil
}
