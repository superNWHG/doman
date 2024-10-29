package flags

import (
	"errors"
	"os"

	"github.com/spf13/pflag"
	"github.com/superNWHG/doman/internal/config"
	"github.com/superNWHG/doman/internal/data"
)

type (
	Defaults struct {
		NewRepo `toml:"NewRepo"`
		Init    `toml:"Init"`
		Add     `toml:"Add"`
		Read    `toml:"Read"`
		Sync    `toml:"Sync"`
		Link    `toml:"Link"`
		Edit    `toml:"Edit"`
		Config  `toml:"Config"`
	}

	NewRepo struct {
		NewRepoClone    bool   `default:"false" toml:"newRepoClone"`
		NewRepoDataFile bool   `default:"true" toml:"newRepoDataFile"`
		NewRepoUrl      string `default:"" toml:"newRepoUrl"`
	}

	Init struct {
	}

	Add struct {
		AddName     string `default:"" toml:"addName"`
		AddEntry    string `default:"" toml:"addEntry"`
		AddExisting bool   `default:"false" toml:"addExisting"`
		AddFormat   bool   `default:"true" toml:"addFormat"`
	}

	Read struct {
	}

	Sync struct {
		SyncMessage string   `default:"New changes" toml:"syncMessage"`
		SyncFiles   []string `default:"[]string{}" toml:"syncFiles"`
		SyncAuth    bool     `default:"true" toml:"syncAuth"`
		SyncPush    bool     `default:"false" toml:"syncPush"`
	}

	Link struct {
	}

	Edit struct {
		EditName   string `default:"" toml:"editName"`
		EditEditor string `default:"" toml:"editEditor"`
		EditFormat bool   `default:"true" toml:"editFormat"`
	}

	Config struct {
		ConfigNew  bool `default:"false" toml:"configNew"`
		ConfigRead bool `default:"false" toml:"configRead"`
	}
)

func getDefaults(path string) (*Defaults, error) {
	userDefaults, err := config.ReadConfigAny(path, &Defaults{})
	if err != nil {
		if err.Error() == "Config file does not exist" {
			return &Defaults{}, nil
		}
		return nil, err
	}

	if userDefaults == nil {
		userDefaults = &Defaults{}
		return userDefaults.(*Defaults), err
	}

	return userDefaults.(*Defaults), nil
}

func SetSubcommands() error {
	path := pflag.String("path", "./", "Path to the repo")

	defaults, err := getDefaults(*path)
	if err != nil {
		return err
	}

	newRepoCmd := pflag.NewFlagSet("new", pflag.ExitOnError)
	newRepoClone := newRepoCmd.Bool("clone", defaults.NewRepoClone, "Set to true to clone a repo instead of initializing a new one")
	newRepoDataFile := newRepoCmd.Bool("datafile", defaults.NewRepoDataFile, "Set to false if you don't wat to create a data file to keep track of your dotfiles")
	newRepoUrl := newRepoCmd.String("url", defaults.NewRepoUrl, "URL of the repo")

	initCmd := pflag.NewFlagSet("init", pflag.ExitOnError)

	addCmd := pflag.NewFlagSet("add", pflag.ExitOnError)
	addName := addCmd.String("name", defaults.AddName, "Name of the dotfile entry")
	addEntry := addCmd.String("entry", defaults.AddEntry, "Path to the new dotfile entry")
	addExisting := addCmd.Bool("existing", defaults.AddExisting, "Set to true to add an existing file in your dotfiles directory")
	addFormat := addCmd.Bool("format", defaults.AddFormat, "Automatically format dotfiles.json")

	readCmd := pflag.NewFlagSet("read", pflag.ExitOnError)

	syncCmd := pflag.NewFlagSet("sync", pflag.ExitOnError)
	syncMessage := syncCmd.String("message", defaults.SyncMessage, "Custom commit message")
	syncFiles := syncCmd.StringSlice("files", defaults.SyncFiles, "Files you want to sync (leave empty to sync all)")
	syncAuth := syncCmd.Bool("authentication", defaults.SyncAuth, "Set to false to not ask for username and password")
	syncPush := syncCmd.Bool("push", defaults.SyncPush, "Set to true to automatically push to the remote repository")

	linkCmd := pflag.NewFlagSet("link", pflag.ExitOnError)

	editCmd := pflag.NewFlagSet("edit", pflag.ExitOnError)
	editName := editCmd.String("name", defaults.EditName, "Name of the dotfile entry to edit")
	editEditor := editCmd.String("editor", defaults.EditEditor, "Editor you want to use (leave empty to use default)")
	editFormat := editCmd.Bool("format", defaults.EditFormat, "Automatically format dotfiles.json")

	configCmd := pflag.NewFlagSet("config", pflag.ExitOnError)
	configNew := configCmd.Bool("new", defaults.ConfigNew, "Create a new config file")
	configRead := configCmd.Bool("read", defaults.ConfigRead, "Read the config file")

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
			if err := newCloneRepo(*path, *newRepoUrl); err != nil {
				return err
			}
		} else {
			if err := newInitRepo(*path, *newRepoUrl); err != nil {
				return err
			}
		}

		if *newRepoDataFile {
			if err := data.NewDataFile(*path); err != nil {
				return err
			}
		}
		return nil

	case "init":
		if err := data.NewDataFile(*path); err != nil {
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
			if err := addExistingData(*path, *addName, *addEntry, *addFormat); err != nil {
				return err
			}
		} else {
			if err := addData(*path, *addName, *addEntry, *addFormat); err != nil {
				return err
			}
		}

		return nil

	case "read":
		if err := readCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := readData(*path); err != nil {
			return err
		}

		return nil

	case "sync":
		if err := syncCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := data.Sync(*path, *syncMessage, *syncPush, *syncAuth, *syncFiles); err != nil {
			return err
		}

		return nil

	case "link":
		if err := linkCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := data.LinkData(*path); err != nil {
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

		if err := data.EditData(*path, *editName, *editEditor, *editFormat); err != nil {
			return err
		}

		return nil

	case "config":
		if err := configCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if *configNew && *configRead {
			err := errors.New("Only one flag allowed")
			return err
		}

		if !*configNew && !*configRead {
			err := errors.New("Expected flag")
			return err
		}

		if *configNew {
			if err := config.NewConfig(*path, Defaults{}); err != nil {
				return err
			}
		}

		if *configRead {
			if err := readconfig(*path, &Defaults{}); err != nil {
				return err
			}
		}

		return nil
	}

	getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd, *editCmd, *configCmd)
	err = errors.New("Invalid subcommand")
	return err
}
