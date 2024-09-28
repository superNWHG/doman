package cmd

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
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

func InitRepo(path string, url string) error {
	repo, err := git.PlainInit(path, false)
	if err != nil {
		return err
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})
	if err != nil {
		return err
	}

	return nil
}

func Add(path string, files []string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}

	for i := range files {
		_, err := workTree.Add(files[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func Commit(path string, commitMessage string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}
	_, err = workTree.Commit(commitMessage, &git.CommitOptions{})
	if err != nil {
		return err
	}

	return nil
}

func Push(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	if err := repo.Push(&git.PushOptions{}); err != nil {
		return err
	}

	return nil
}
