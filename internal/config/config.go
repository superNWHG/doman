package config

import (
	"errors"
	"os"
	"path/filepath"
)

type (
	config struct {
		Add  Add  `toml:"new"`
		Sync Sync `toml:"sync"`
		Edit Edit `toml:"edit"`
	}

	Add struct {
		Existing bool `toml:"existing"`
	}

	Sync struct {
		Authentication bool   `toml:"authentication"`
		Message        string `toml:"message"`
		Push           bool   `toml:"push"`
	}

	Edit struct {
		Editor string `toml:"editor"`
	}
)

func setDefaults() *config {
	return &config{
		Add: Add{
			Existing: false,
		},
		Sync: Sync{
			Authentication: true,
			Message:        "New changes",
			Push:           false,
		},
		Edit: Edit{
			Editor: "",
		},
	}
}

func NewConfig(path string) error {
	path = filepath.Join(path, "config.toml")

	if _, err := os.Stat(path); err != nil {
		defaults := setDefaults()
		tomlData, err := encodeToml(defaults)
		if err != nil {
			return err
		}

		if err := os.WriteFile(path, tomlData, 0644); err != nil {
			return err
		}
	} else {
		err := errors.New("Config file already exists")
		return err
	}

	return nil
}
