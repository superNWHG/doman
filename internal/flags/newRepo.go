package flags

import "github.com/superNWHG/doman/cmd"

func newCloneRepo(path string, url string) error {
	if err := cmd.CloneRepo(path, url); err != nil {
		return err
	}

	return nil
}

func newInitRepo(path string, url string) error {
	if err := cmd.InitRepo(path, url); err != nil {
		return err
	}

	return nil
}
