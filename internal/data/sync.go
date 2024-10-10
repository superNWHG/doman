package data

import (
	"path/filepath"
	"slices"

	"github.com/superNWHG/doman/internal/git"
	"github.com/superNWHG/doman/pkg/gitcredentials"
)

func Sync(path string, message string, push bool, auth bool, filesToSync []string) error {
	dataPath := filepath.Join(path, "dotfiles.json")
	files, _, _, err := ReadDataFile(dataPath)
	if err != nil {
		return err
	}

	if len(filesToSync) == 0 {
		files = append(files, "dotfiles.json")
	} else {
		for i := range filesToSync {
			if !slices.Contains(files, filesToSync[i]) {
				filesToSync = slices.Delete(filesToSync, i, i)
			}
		}
		files = filesToSync
	}

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
