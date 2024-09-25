package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneRepo(path string, url string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}
