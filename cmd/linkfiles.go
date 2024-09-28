package cmd

import (
	"fmt"
	"os"
)

func NewLink(oldPath []string, newPath []string) error {
	fmt.Println("Old path:", oldPath)
	fmt.Println("New path:", newPath)
	for i := range oldPath {
		if err := os.Symlink(oldPath[i], newPath[i]); err != nil {
			return err
		}
	}

	return nil
}
