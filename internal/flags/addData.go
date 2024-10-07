package flags

import (
	"path/filepath"

	"github.com/superNWHG/doman/internal/data"
	"github.com/superNWHG/doman/pkg/symlink"
)

func addData(path string, name string, newPath string) error {
	if name == "" {
		for i := len(newPath); i > 0; i-- {
			if string(newPath[i-1]) == "/" {
				name = newPath[i:]
				i = 0
			}
		}

	}

	namePath := filepath.Join(path, name)

	path = filepath.Join(path, "/dotfiles.json")

	nameSlice := []string{name}
	namePathSlice := []string{namePath}
	newPathSlice := []string{newPath}

	if err := data.NewData(path, nameSlice, newPathSlice); err != nil {
		return err
	}

	if err := symlink.NewLink(newPathSlice, namePathSlice, "deleteOld"); err != nil {
		return err
	}

	return nil
}

func addExistingData(path string, oldPath string, newPath string) error {
	path = filepath.Join(path, "dotfiles.json")

	var name string

	for i := len(newPath); i > 0; i-- {
		if string(newPath[i-1]) == "/" {
			name = newPath[i:]
			i = 0
		}
	}

	nameSlice := []string{name}
	oldPathSlice := []string{oldPath}
	newPathSlice := []string{newPath}

	if err := data.NewData(path, nameSlice, newPathSlice); err != nil {
		return err
	}

	if err := symlink.NewLink(newPathSlice, oldPathSlice, "deleteOldDelete"); err != nil {
		return err
	}

	return nil
}
