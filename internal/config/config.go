package config

import (
	"errors"
	"os"
	"path/filepath"
)

func NewConfig(path string, configStruct interface{}) error {
	path = filepath.Join(path, "config.toml")

	if _, err := os.Stat(path); err != nil {
		tomlData, err := encodeToml(configStruct)
		if err != nil {
			return err
		}

		if err := os.WriteFile(path, tomlData, 0644); err != nil {
			return err
		}
	} else {
		return errors.New("Config file already exists")
	}

	return nil
}

func ReadConfig(path string, configStruct any) (any, error) {
	path = filepath.Join(path, "config.toml")

	if _, err := os.Stat(path); err != nil {
		return nil, errors.New("Config file does not exist")
	}

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	userConfig, err := decodeTomlAny(fileContent, configStruct)
	if err != nil {
		return nil, err
	}

	return userConfig, nil
}
