package flags

import (
	"errors"
	"flag"
	"os"

	"github.com/superNWHG/doman/internal/data"
)

func SetSubcommands() error {
	newRepoCmd := flag.NewFlagSet("new", flag.ExitOnError)
	newRepoClone := newRepoCmd.Bool("clone", false, "Set to true to clone a repo instead of initializing a new one")
	newRepoDataFile := newRepoCmd.Bool("datafile", true, "Set to false if you don't wat to create a data file to keep track of your dotfiles")
	newRepoPath := newRepoCmd.String("path", "./", "Path to the new repo")
	newRepoUrl := newRepoCmd.String("url", "", "URL of the repo")

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	initPath := initCmd.String("path", "./", "Path to the repo")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addPath := addCmd.String("path", "./", "Path to the repo")
	addName := addCmd.String("name", "", "Name of the dotfile entry")
	addEntry := addCmd.String("entry", "", "Path to the new dotfile entry")

	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	readPath := readCmd.String("path", "./", "Path to the repo")

	syncCmd := flag.NewFlagSet("sync", flag.ExitOnError)
	syncPath := syncCmd.String("path", "./", "Path to the repo")
	syncMessage := syncCmd.String("message", "New changes", "Custom commit message")
	syncAuth := syncCmd.Bool("authentication", true, "Set to false to not ask for username and password")
	syncPush := syncCmd.Bool("push", false, "Set to true to automatically push to the remote repository")

	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkPath := linkCmd.String("path", "./", "Path to the repo")

	if len(os.Args) < 2 {
		getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd)
		err := errors.New("Expected subcommand")
		return err
	}

	switch os.Args[1] {
	case "new":
		if err := newRepoCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if *newRepoUrl == "" {
			err := errors.New("Url flag is required")
			return err
		}

		if *newRepoClone {
			if err := newCloneRepo(*newRepoPath, *newRepoUrl); err != nil {
				return err
			}
		} else {
			if err := newInitRepo(*newRepoPath, *newRepoUrl); err != nil {
				return err
			}
		}

		if *newRepoDataFile {
			if err := data.NewDataFile(*newRepoPath); err != nil {
				return err
			}
		}
		return nil

	case "init":
		if err := data.NewDataFile(*initPath); err != nil {
			return err
		}

		return nil

	case "add":
		if err := addCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if *addEntry == "" {
			err := errors.New("Entry flag is required")
			return err
		}

		if err := addData(*addPath, *addName, *addEntry); err != nil {
			return err
		}

		return nil

	case "read":
		if err := readCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := readData(*readPath); err != nil {
			return err
		}

		return nil

	case "sync":
		if err := syncCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := data.Sync(*syncPath, *syncMessage, *syncPush, *syncAuth); err != nil {
			return err
		}

		return nil

	case "link":
		if err := linkCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := data.LinkData(*linkPath); err != nil {
			return err
		}

		return nil
	}

	err := errors.New("Invalid subcommand")
	return err
}
