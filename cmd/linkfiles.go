package cmd

import "os"

func NewLink(oldPath []string, newPath []string) error {
	for i := range oldPath {
		if err := os.Symlink(oldPath[i], newPath[i]); err != nil {
			return err
		}
	}

	return nil
}
