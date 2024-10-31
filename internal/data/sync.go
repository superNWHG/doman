package data

import (
	"path/filepath"
	"slices"
	"strings"

	"github.com/superNWHG/doman/internal/git"
	"github.com/superNWHG/doman/pkg/gitcredentials"
)

func Sync(
	path string, message string, push bool, auth bool, gitAuth bool, filesToSync []string,
) error {
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
		name, mail, pass, err = gitcredentials.AskGitCredentials()
		if err != nil {
			return err
		}
	} else if gitAuth {
		remote, err := git.GetRemote(path)
		if err != nil {
			return err
		}

		remote = strings.Split(remote, "https://")[1]
		remote = strings.Split(remote, "/")[0]
		name, mail, pass, err = gitcredentials.GetGitCredentials(remote)
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
