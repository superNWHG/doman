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
		Install `toml:"Install"`
	}

	NewRepo struct {
		NewRepoClone    bool   `default:"false" toml:"newRepoClone"`
		NewRepoDataFile bool   `default:"true" toml:"newRepoDataFile"`
		NewRepoUrl      string `default:"" toml:"newRepoUrl"`
		NewRepoPath     string `default:"./" toml:"path"`
	}

	Init struct {
		InitPath string `default:"./" toml:"path"`
	}

	Add struct {
		AddName     string `default:"" toml:"addName"`
		AddEntry    string `default:"" toml:"addEntry"`
		AddPath     string `default:"./" toml:"addPath"`
		AddExisting bool   `default:"false" toml:"addExisting"`
		AddFormat   bool   `default:"true" toml:"addFormat"`
	}

	Read struct {
		ReadPath string `default:"./" toml:"readPath"`
	}

	Sync struct {
		SyncMessage string   `default:"New changes" toml:"syncMessage"`
		SyncPath    string   `default:"./" toml:"syncPath"`
		SyncFiles   []string `default:"[]string{}" toml:"syncFiles"`
		SyncAuth    bool     `default:"false" toml:"syncAuth"`
		SyncGitAuth bool     `default:"false" toml:"syncGitAuth"`
		SyncPush    bool     `default:"false" toml:"syncPush"`
	}

	Link struct {
		LinkPath string `default:"./" toml:"linkPath"`
	}

	Edit struct {
		EditName   string `default:"" toml:"editName"`
		EditEditor string `default:"" toml:"editEditor"`
		EditPath   string `default:"./" toml:"editPath"`
		EditFormat bool   `default:"true" toml:"editFormat"`
	}

	Config struct {
		configPath string `default:"./" toml:"configPath"`
		ConfigNew  bool   `default:"false" toml:"configNew"`
		ConfigRead bool   `default:"false" toml:"configRead"`
	}

	Install struct {
		InstallOs    string   `default:"" toml:"os"`
		InstallPath  string   `default:"./" toml:"installPath"`
		InstallNames []string `default:"[]string{}" toml:"installNames"`
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

	return userDefaults.(*Defaults), nil
}

func SetSubcommands() error {
	defaults, err := getDefaults("./")
	if err != nil {
		return err
	}

	newRepoCmd := pflag.NewFlagSet("new", pflag.ExitOnError)
	newRepoUrl := newRepoCmd.String("url", defaults.NewRepoUrl, "URL to the repository")
	newRepoPath := newRepoCmd.String("path", defaults.NewRepoPath, "Path to the dotfiles directory")
	newRepoClone := newRepoCmd.Bool("clone", defaults.NewRepoClone, "Clone a repository instead of initializing a new one")
	newRepoDataFile := newRepoCmd.Bool("datafile", defaults.NewRepoDataFile, "Don't create a data file to keep track of your dotfiles")

	initCmd := pflag.NewFlagSet("init", pflag.ExitOnError)
	initPath := initCmd.String("path", defaults.InitPath, "Path to the dotfiles directory")

	addCmd := pflag.NewFlagSet("add", pflag.ExitOnError)
	addName := addCmd.String("name", defaults.AddName, "Name of the dotfile entry")
	addEntry := addCmd.String("entry", defaults.AddEntry, "Path to the new dotfile entry")
	addPath := addCmd.String("path", defaults.AddPath, "Path to the dotfiles directory")
	addExisting := addCmd.Bool("existing", defaults.AddExisting, "Add an existing file in your dotfiles directory")
	addFormat := addCmd.Bool("format", defaults.AddFormat, "Automatically format dotfiles.json")

	readCmd := pflag.NewFlagSet("read", pflag.ExitOnError)
	readPath := readCmd.String("path", defaults.ReadPath, "Path to the dotfiles directory")

	syncCmd := pflag.NewFlagSet("sync", pflag.ExitOnError)
	syncMessage := syncCmd.String("message", defaults.SyncMessage, "Custom commit message")
	syncPath := syncCmd.String("path", defaults.SyncPath, "Path to the dotfiles directory")
	syncFiles := syncCmd.StringSlice("files", defaults.SyncFiles, "Files you want to sync (leave empty to sync all)")
	syncAuth := syncCmd.Bool("authentication", defaults.SyncAuth, "Ask for username and password")
	syncGitAuth := syncCmd.Bool("gitauthentication", defaults.SyncGitAuth, "Use local git credentials")
	syncPush := syncCmd.Bool("push", defaults.SyncPush, "Automatically push to the remote repository")

	linkCmd := pflag.NewFlagSet("link", pflag.ExitOnError)
	linkPath := linkCmd.String("path", defaults.LinkPath, "Path to the dotfiles directory")

	editCmd := pflag.NewFlagSet("edit", pflag.ExitOnError)
	editName := editCmd.String("name", defaults.EditName, "Name of the dotfile entry to edit")
	editEditor := editCmd.String("editor", defaults.EditEditor, "Editor you want to use (leave empty to use $EDITOR)")
	editPath := editCmd.String("path", defaults.EditPath, "Path to the dotfiles directory")
	editFormat := editCmd.Bool("format", defaults.EditFormat, "Automatically format dotfiles.json")

	configCmd := pflag.NewFlagSet("config", pflag.ExitOnError)
	configPath := configCmd.String("path", defaults.configPath, "Path to the dotfiles directory")
	configNew := configCmd.Bool("new", defaults.ConfigNew, "Create a new configuration file")
	configRead := configCmd.Bool("read", defaults.ConfigRead, "Read the configuration file")

	installCmd := pflag.NewFlagSet("install", pflag.ExitOnError)
	installOs := installCmd.String("os", defaults.InstallOs, "Specify the os")
	installPath := installCmd.String("path", defaults.InstallPath, "Path to the dotfiles directory")
	installNames := installCmd.StringSlice("names", defaults.InstallNames, "Specify the names of the packages you want to install")

	if len(os.Args) < 2 {
		getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd, *editCmd, *configCmd, *installCmd)
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
			if err := addExistingData(*addPath, *addName, *addEntry, *addFormat); err != nil {
				return err
			}
		} else {
			if err := addData(*addPath, *addName, *addEntry, *addFormat); err != nil {
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

		if *syncAuth && *syncGitAuth {
			err := errors.New("Only one authentication flag allowed")
			return err
		}

		if err := data.Sync(*syncPath, *syncMessage, *syncPush, *syncAuth, *syncGitAuth, *syncFiles); err != nil {
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

		if err := data.EditData(*editPath, *editName, *editEditor, *editFormat); err != nil {
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
			if err := config.NewConfig(*configPath, Defaults{}); err != nil {
				return err
			}
		}

		if *configRead {
			if err := readconfig(*configPath, &Defaults{}); err != nil {
				return err
			}
		}

		return nil

	case "install":
		if err := installCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := install(*installPath, *installNames, *installOs); err != nil {
			return err
		}

		return nil
	}

	getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd, *editCmd, *configCmd, *installCmd)
	err = errors.New("Invalid subcommand")
	return err
}
