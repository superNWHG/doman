package flags

import (
	"errors"
	"flag"
	"fmt"
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

	if len(os.Args) < 2 {
		err := errors.New("Expected subcommand")
		return err
	}

	switch os.Args[1] {
	case "new":
		err := newRepoCmd.Parse(os.Args[2:])
		fmt.Println(string(*newRepoPath))
		if err != nil {
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
	}

	err := errors.New("Invalid subcommand")
	return err
}