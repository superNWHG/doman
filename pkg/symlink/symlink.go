package symlink

import (
	"errors"
	"os"
)

type Method string

func NewLink(oldPath []string, newPath []string, method Method) error {
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
	default:
		return errors.New("Invalid method")
	}

	for i := range oldPath {
		if err := os.Symlink(newPath[i], oldPath[i]); err != nil {
			return err
		}
	}

	return nil
}
