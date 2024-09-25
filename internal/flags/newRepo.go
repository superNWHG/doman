package flags

import git "github.com/superNWHG/doman/cmd"

func newCloneRepo(path string, url string) error {
	if err := git.CloneRepo(path, url); err != nil {
		return err
	}

	return nil
}

func newInitRepo(path string, url string) error {
	if err := git.InitRepo(path, url); err != nil {
		return err
	}

	return nil
}
