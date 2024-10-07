package symlink

import (
	"errors"
	"os"
)

func NewLink(oldPath []string, newPath []string, method string) error {
	if method != "deleteOld" && method != "deleteNew" && method != "deleteOldDelete" {
		err := errors.New("Invalid method")
		return err
	}
	switch method {
	case "deleteOld":
		for i := range oldPath {
			if err := os.Rename(oldPath[i], newPath[i]); err != nil {
				return err
			}
		}
	case "deleteOldDelete":
		for i := range oldPath {
			if err := os.RemoveAll(oldPath[i]); err != nil {
				return err
			}
		}
	case "deleteNew":
		for i := range newPath {
			if err := os.RemoveAll(newPath[i]); err != nil {
				return err
			}
		}
	}

	for i := range oldPath {
		if err := os.Symlink(newPath[i], oldPath[i]); err != nil {
			return err
		}
	}

	return nil
}
