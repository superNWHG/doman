package flags

import (
	"errors"
	"os"

	"github.com/spf13/pflag"
	"github.com/superNWHG/doman/internal/config"
	"github.com/superNWHG/doman/internal/data"
)

func SetSubcommands() error {
	newRepoCmd := pflag.NewFlagSet("new", pflag.ExitOnError)
	newRepoClone := newRepoCmd.Bool("clone", false, "Set to true to clone a repo instead of initializing a new one")
	newRepoDataFile := newRepoCmd.Bool("datafile", true, "Set to false if you don't wat to create a data file to keep track of your dotfiles")
	newRepoPath := newRepoCmd.String("path", "./", "Path to the new repo")
	newRepoUrl := newRepoCmd.String("url", "", "URL of the repo")

	initCmd := pflag.NewFlagSet("init", pflag.ExitOnError)
	initPath := initCmd.String("path", "./", "Path to the repo")

	addCmd := pflag.NewFlagSet("add", pflag.ExitOnError)
	addPath := addCmd.String("path", "./", "Path to the repo")
	addName := addCmd.String("name", "", "Name of the dotfile entry")
	addEntry := addCmd.String("entry", "", "Path to the new dotfile entry")
	addExisting := addCmd.Bool("existing", false, "Set to true to add an existing file in your dotfiles directory")

	readCmd := pflag.NewFlagSet("read", pflag.ExitOnError)
	readPath := readCmd.String("path", "./", "Path to the repo")

	syncCmd := pflag.NewFlagSet("sync", pflag.ExitOnError)
	syncPath := syncCmd.String("path", "./", "Path to the repo")
	syncMessage := syncCmd.String("message", "New changes", "Custom commit message")
	syncFiles := syncCmd.StringSlice("files", []string{}, "Files you want to sync (leave empty to sync all)")
	syncAuth := syncCmd.Bool("authentication", true, "Set to false to not ask for username and password")
	syncPush := syncCmd.Bool("push", false, "Set to true to automatically push to the remote repository")

	linkCmd := pflag.NewFlagSet("link", pflag.ExitOnError)
	linkPath := linkCmd.String("path", "./", "Path to the repo")

	editCmd := pflag.NewFlagSet("edit", pflag.ExitOnError)
	editPath := editCmd.String("path", "./", "Path to the repo")
	editName := editCmd.String("name", "", "Name of the dotfile entry to edit")
	editEditor := editCmd.String("editor", "", "Editor you want to use (leave empty to use default)")

	configCmd := pflag.NewFlagSet("config", pflag.ExitOnError)
	configPath := configCmd.String("path", "./", "Path to the repo")
	configNew := configCmd.Bool("new", false, "Create a new config file")

	if len(os.Args) < 2 {
		getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd, *editCmd, *configCmd)
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

		if *addExisting {
			if err := addExistingData(*addPath, *addName, *addEntry); err != nil {
				return err
			}
		} else {
			if err := addData(*addPath, *addName, *addEntry); err != nil {
				return err
			}
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

		if err := data.Sync(*syncPath, *syncMessage, *syncPush, *syncAuth, *syncFiles); err != nil {
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

	case "edit":
		if err := editCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if *editName == "" {
			err := errors.New("Name flag is required")
			return err
		}

		if err := data.EditData(*editPath, *editName, *editEditor); err != nil {
			return err
		}

		return nil

	case "config":
		if err := configCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if *configNew {
			if err := config.NewConfig(*configPath); err != nil {
				return err
			}
		}

		return nil
	}

	getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd, *editCmd, *configCmd)
	err := errors.New("Invalid subcommand")
	return err
}
