package cmd

import "os"

func NewLink(oldPath []string, newPath []string) error {
	for i := range oldPath {
		if err := os.Rename(oldPath[i], newPath[i]); err != nil {
			return err
		}

		if err := os.Symlink(newPath[i], oldPath[i]); err != nil {
			return err
		}
	}

	return nil
}
