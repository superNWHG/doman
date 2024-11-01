package data

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/superNWHG/doman/pkg/symlink"
)

func NewDataFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	dataFilePath := filepath.Join(path, "dotfiles.json")
	if _, err := os.Stat(dataFilePath); err == nil {
		return errors.New("Data file already exists")
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

	keys, values, err := jsonToMap(obj)
	if err != nil {
		return nil, nil, nil, err
	}

	combined := make(map[string]interface{})
	for i := range keys {
		combined[keys[i]] = values[i]
	}

	return keys, values, combined, err
}

func NewData(path string, newDataKeys []string, newDataValues []string, format bool) error {
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

	data, err = formatJson(data)
	if err != nil {
		return err
	}

	if err = os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

func EditData(path string, name string, editor string, format bool) error {
	tmpFile := "/tmp/dotfiles.json"

	if editor == "" {
		if editorEnv := os.Getenv("EDITOR"); editorEnv != "" {
			editor = editorEnv
		} else {
			return errors.New("No editor found")
		}
	}

	path = filepath.Join(path, "dotfiles.json")
	_, _, data, err := ReadDataFile(path)
	if err != nil {
		return err
	}

	if data[name] == nil {
		return errors.New("Name not found")
	}

	nameData := map[string]interface{}{name: data[name]}
	jsonData, err := encodeJson(nameData)
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmpFile, jsonData, 0644); err != nil {
		return err
	}

	cmd := exec.Command(editor, tmpFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	newJsonData, err := os.ReadFile(tmpFile)
	if err != nil {
		return err
	}

	obj, err := decodeJson(newJsonData)
	if err != nil {
		return err
	}

	keys, values, err := jsonToMap(obj)
	if err != nil {
		return err
	}

	data[keys[0]] = values[0]

	delete(data, name)

	newJson, err := encodeJson(data)
	if err != nil {
		return err
	}

	if format {
		newJson, err = formatJson(newJson)
		if err != nil {
			return err
		}
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	if err := os.WriteFile(path, newJson, 0644); err != nil {
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
