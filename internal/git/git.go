package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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

func Commit(path string, commitMessage string, name string, email string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}

	if name == "" || email == "" {
		_, err = workTree.Commit(commitMessage, &git.CommitOptions{
			AllowEmptyCommits: false,
			Author: &object.Signature{
				Name:  name,
				Email: email,
			},
		})
	} else {
		_, err = workTree.Commit(commitMessage, &git.CommitOptions{
			AllowEmptyCommits: false,
		})
	}
	if err != nil {
		return err
	}

	return nil
}

func Push(path string, name string, password string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	if err := repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: name,
			Password: password,
		},
	}); err != nil {
		return err
	}

	return nil
}

func GetRemote(path string) (string, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	remotes, err := repo.Remotes()
	if err != nil {
		return "", err
	}

	return remotes[0].String(), nil
}

func Status(path string) (*git.Status, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	workTree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, err
	}

	return &status, nil
}
