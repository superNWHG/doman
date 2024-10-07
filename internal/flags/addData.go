package flags

import (
	"path/filepath"

	"github.com/superNWHG/doman/internal/data"
	"github.com/superNWHG/doman/pkg/symlink"
)

func addData(path string, name string, newPath string) error {
	if name == "" {
		name = filepath.Base(newPath)
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
	var name string
	name, err := filepath.Rel(path, oldPath)
	if err != nil {
		return err
	}

	path = filepath.Join(path, "dotfiles.json")

	oldPath, err = filepath.Abs(oldPath)
	if err != nil {
		return err
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
