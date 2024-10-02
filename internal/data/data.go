package data

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/superNWHG/doman/pkg/symlink"
)

func NewDataFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	dataFilePath := filepath.Join(path, "dotfiles.json")
	if _, err := os.Stat(dataFilePath); err == nil {
		err := errors.New("Data file already exists")
		return err
	}

	if _, err := os.Create(dataFilePath); err != nil {
		return err
	}

	if err := os.WriteFile(dataFilePath, []byte("{}"), 0644); err != nil {
		return err
	}

	return nil
}

func ReadDataFile(path string) ([]string, []string, map[string]interface{}, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}

	obj, err := decodeJson(file)
	if err != nil {
		return nil, nil, nil, err
	}

	keys := []string{}
	values := []string{}

	for key, rawMsg := range obj {
		var value string
		if err := json.Unmarshal(*rawMsg, &value); err != nil {
			return nil, nil, nil, err
		}

		keys = append(keys, key)
		values = append(values, value)
	}

	combined := make(map[string]interface{})
	for i := range keys {
		combined[keys[i]] = values[i]
	}

	return keys, values, combined, err
}

func NewData(path string, newDataKeys []string, newDataValues []string) error {
	_, _, oldData, err := ReadDataFile(path)
	if err != nil {
		return err
	}

	for i := range newDataKeys {
		oldData[newDataKeys[i]] = newDataValues[i]
	}

	data, err := encodeJson(oldData)
	if err != nil {
		return err
	}

	if err = os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

func LinkData(path string) error {
	oldPaths, newPaths, _, err := ReadDataFile(path)
	if err != nil {
		return err
	}

	newPathsToLink := []string{}
	oldPathsToLink := []string{}
	for i := range newPaths {
		info, err := os.Lstat(newPaths[i])
		if err != nil {
			return err
		}
		if info.Mode() == os.ModeSymlink {
			newPathsToLink = append(newPathsToLink, newPaths[i])
			oldPathsToLink = append(oldPathsToLink, oldPaths[i])
		}
	}

	if err := symlink.NewLink(oldPathsToLink, newPathsToLink, "deleteNew"); err != nil {
		return err
	}

	return nil
}
