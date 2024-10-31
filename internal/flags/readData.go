package flags

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/superNWHG/doman/internal/data"
)

func readData(path string) error {
	path = filepath.Join(path, "dotfiles.json")
	values, _, entries, err := data.ReadDataFile(path)
	if err != nil {
		return err
	}

	longest := 0
	for _, v := range values {
		if len(v) > longest {
			longest = len(v)
		}
	}

	for i, v := range values {
		diff := longest - len(v)
		fmt.Println(values[i]+strings.Repeat(" ", diff)+":", entries[values[i]])
	}

	return nil
}
