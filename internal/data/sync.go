package data

import (
	"path/filepath"

	"github.com/superNWHG/doman/cmd"
)

func Sync(path string, message string, push bool) error {
	dataPath := filepath.Join(path, "dotfiles.json")
	err, files, _, _ := ReadDataFile(dataPath)
	if err != nil {
		return err
	}

	files = append(files, "dotfiles.json")

	if err := cmd.Add(path, files); err != nil {
		return err
	}

	if err := cmd.Commit(path, message); err != nil {
		return err
	}

	if push {
		if err := cmd.Push(path); err != nil {
			return err
		}
	}

	return nil
}
