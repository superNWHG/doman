package data

import (
	"path/filepath"

	"github.com/superNWHG/doman/internal/git"
	"github.com/superNWHG/doman/pkg/gitcredentials"
)

func Sync(path string, message string, push bool, auth bool) error {
	dataPath := filepath.Join(path, "dotfiles.json")
	files, _, _, err := ReadDataFile(dataPath)
	if err != nil {
		return err
	}

	files = append(files, "dotfiles.json")

	var name, pass, mail string
	if auth {
		err, name, mail, pass = gitcredentials.GetGitCredentials()
		if err != nil {
			return err
		}
	} else {
		name = ""
		mail = ""
		pass = ""
	}

	if err := git.Add(path, files); err != nil {
		return err
	}

	if err := git.Commit(path, message, name, mail); err != nil {
		return err
	}

	if push {
		if err := git.Push(path, name, pass); err != nil {
			return err
		}
	}

	return nil
}
