package data

import (
	"path/filepath"

	"github.com/superNWHG/doman/cmd"
)

func Sync(path string, message string, push bool, auth bool) error {
	dataPath := filepath.Join(path, "dotfiles.json")
	err, files, _, _ := ReadDataFile(dataPath)
	if err != nil {
		return err
	}

	files = append(files, "dotfiles.json")

	var name, pass, mail string
	if auth {
		err, name, mail, pass = cmd.GetGitCredentials()
		if err != nil {
			return err
		}
	} else {
		name = ""
		mail = ""
		pass = ""
	}

	if err := cmd.Add(path, files); err != nil {
		return err
	}

	if err := cmd.Commit(path, message, name, mail); err != nil {
		return err
	}

	if push {
		if err := cmd.Push(path, name, pass); err != nil {
			return err
		}
	}

	return nil
}
