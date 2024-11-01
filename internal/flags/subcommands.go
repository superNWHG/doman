package flags

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/superNWHG/doman/internal/config"
	"github.com/superNWHG/doman/internal/data"
)

type (
	ConfigOptions struct {
		NewRepo `toml:"NewRepo"`
		Add     `toml:"Add"`
		Sync    `toml:"Sync"`
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

	Add struct {
		AddName     string `default:"" toml:"addName"`
		AddEntry    string `default:"" toml:"addEntry"`
		AddPath     string `default:"./" toml:"addPath"`
		AddExisting bool   `default:"false" toml:"addExisting"`
		AddFormat   bool   `default:"true" toml:"addFormat"`
	}

	Sync struct {
		SyncMessage string   `default:"New changes" toml:"syncMessage"`
		SyncPath    string   `default:"./" toml:"syncPath"`
		SyncFiles   []string `default:"[]string{}" toml:"syncFiles"`
		SyncAuth    bool     `default:"false" toml:"syncAuth"`
		SyncGitAuth bool     `default:"false" toml:"syncGitAuth"`
		SyncPush    bool     `default:"false" toml:"syncPush"`
	}

	Edit struct {
		EditName   string `default:"" toml:"editName"`
		EditEditor string `default:"" toml:"editEditor"`
		EditPath   string `default:"./" toml:"editPath"`
		EditFormat bool   `default:"true" toml:"editFormat"`
	}

	Config struct {
		ConfigNew  bool `default:"false" toml:"configNew"`
		ConfigRead bool `default:"false" toml:"configRead"`
	}

	Install struct {
		InstallOs    string   `default:"" toml:"os"`
		InstallPath  string   `default:"./" toml:"installPath"`
		InstallNames []string `default:"[]string{}" toml:"installNames"`
	}
)

func getUserConfig(path string) (*ConfigOptions, error) {
	userDefaults, err := config.ReadConfig(path, &ConfigOptions{})
	if err != nil {
		if err.Error() == "Config file does not exist" {
			return &ConfigOptions{}, nil
		}
		return nil, err
	}

	return userDefaults.(*ConfigOptions), nil
}

func SetSubcommands() error {
	var path string
	for i, v := range os.Args {
		if strings.Contains(v, "path") {
			if strings.Contains(v, "=") {
				path = strings.Split(v, "=")[1]
				os.Args = append(os.Args[:i], os.Args[i+1:]...)
			} else {
				path = os.Args[i+1]
				os.Args = append(os.Args[:i], os.Args[i+2:]...)
			}
			break
		}
		path = "./"
	}

	defaults, err := getUserConfig(path)
	if err != nil {
		return err
	}

	newRepoCmd := pflag.NewFlagSet("new", pflag.ExitOnError)
	newRepoUrl := newRepoCmd.String("url", defaults.NewRepoUrl, "URL to the repository")
	newRepoClone := newRepoCmd.Bool("clone", defaults.NewRepoClone, "Clone a repository instead of initializing a new one")
	newRepoDataFile := newRepoCmd.Bool("datafile", defaults.NewRepoDataFile, "Don't create a data file to keep track of your dotfiles")

	initCmd := pflag.NewFlagSet("init", pflag.ExitOnError)

	addCmd := pflag.NewFlagSet("add", pflag.ExitOnError)
	addName := addCmd.String("name", defaults.AddName, "Name of the dotfile entry")
	addEntry := addCmd.String("entry", defaults.AddEntry, "Path to the new dotfile entry")
	addExisting := addCmd.Bool("existing", defaults.AddExisting, "Add an existing file in your dotfiles directory")
	addFormat := addCmd.Bool("format", defaults.AddFormat, "Automatically format dotfiles.json")

	readCmd := pflag.NewFlagSet("read", pflag.ExitOnError)

	syncCmd := pflag.NewFlagSet("sync", pflag.ExitOnError)
	syncMessage := syncCmd.String("message", defaults.SyncMessage, "Custom commit message")
	syncFiles := syncCmd.StringSlice("files", defaults.SyncFiles, "Files you want to sync (leave empty to sync all)")
	syncAuth := syncCmd.Bool("authentication", defaults.SyncAuth, "Ask for username and password")
	syncGitAuth := syncCmd.Bool("gitauthentication", defaults.SyncGitAuth, "Use local git credentials")
	syncPush := syncCmd.Bool("push", defaults.SyncPush, "Automatically push to the remote repository")

	linkCmd := pflag.NewFlagSet("link", pflag.ExitOnError)

	editCmd := pflag.NewFlagSet("edit", pflag.ExitOnError)
	editName := editCmd.String("name", defaults.EditName, "Name of the dotfile entry to edit")
	editEditor := editCmd.String("editor", defaults.EditEditor, "Editor you want to use (leave empty to use $EDITOR)")
	editFormat := editCmd.Bool("format", defaults.EditFormat, "Automatically format dotfiles.json")

	configCmd := pflag.NewFlagSet("config", pflag.ExitOnError)
	configNew := configCmd.Bool("new", defaults.ConfigNew, "Create a new configuration file")
	configRead := configCmd.Bool("read", defaults.ConfigRead, "Read the configuration file")

	installCmd := pflag.NewFlagSet("install", pflag.ExitOnError)
	installOs := installCmd.String("os", defaults.InstallOs, "Specify the os")
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
			if err := newCloneRepo(path, *newRepoUrl); err != nil {
				return err
			}
		} else {
			if err := newInitRepo(path, *newRepoUrl); err != nil {
				return err
			}
		}

		if *newRepoDataFile {
			if err := data.NewDataFile(path); err != nil {
				return err
			}
		}
		return nil

	case "init":
		if err := data.NewDataFile(path); err != nil {
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
			if err := addExistingData(path, *addName, *addEntry, *addFormat); err != nil {
				return err
			}
		} else {
			if err := addData(path, *addName, *addEntry, *addFormat); err != nil {
				return err
			}
		}

		return nil

	case "read":
		if err := readCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := readData(path); err != nil {
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

		if err := data.Sync(path, *syncMessage, *syncPush, *syncAuth, *syncGitAuth, *syncFiles); err != nil {
			return err
		}

		return nil

	case "link":
		if err := linkCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := data.LinkData(path); err != nil {
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

		if err := data.EditData(path, *editName, *editEditor, *editFormat); err != nil {
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
			if err := config.NewConfig(path, ConfigOptions{}); err != nil {
				return err
			}
		}

		if *configRead {
			if err := readConfig(path); err != nil {
				return err
			}
		}

		return nil

	case "install":
		if err := installCmd.Parse(os.Args[2:]); err != nil {
			return err
		}

		if err := install(path, *installNames, *installOs); err != nil {
			return err
		}

		return nil
	}

	getHelp(*newRepoCmd, *initCmd, *addCmd, *readCmd, *syncCmd, *linkCmd, *editCmd, *configCmd, *installCmd)
	err = errors.New("Invalid subcommand")
	return err
}
