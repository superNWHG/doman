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
		err := errors.New("Config file already exists")
		return err
	}

	return nil
}

func ReadConfig(path string, configStruct interface{}) (interface{}, error) {
	path = filepath.Join(path, "config.toml")

	if _, err := os.Stat(path); err != nil {
		err := errors.New("Config file does not exist")
		return nil, err
	}

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	userConfig, err := decodeToml(fileContent, &configStruct)
	if err != nil {
		return nil, err
	}

	return userConfig, nil
}
