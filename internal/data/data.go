package data

import (
	"encoding/json"
	"errors"
	"os"
)

func NewDataFile(path string) error {
	path = checkForSlash(path)
	if _, err := os.Stat(path); err != nil {
		return err
	}
	dataFilePath := path + "/dotfiles.json"
	if _, err := os.Stat(dataFilePath); err == nil {
		err := errors.New("Data file already exists")
		return err
	}

	if _, err := os.Create(dataFilePath); err != nil {
		return err
	}

	return nil
}

func ReadDataFile(path string) (error, []string, []string, map[string]interface{}) {
	path = checkForSlash(path)

	file, err := os.ReadFile(path)
	if err != nil {
		return err, nil, nil, nil
	}

	err, obj := decodeJson(file)
	if err != nil {
		return err, nil, nil, nil
	}

	keys := []string{}
	values := []string{}

	for key, rawMsg := range obj {
		var value string
		if err := json.Unmarshal(*rawMsg, &value); err != nil {
			return err, nil, nil, nil
		}

		keys = append(keys, key)
		values = append(values, value)
	}

	combined := make(map[string]interface{})
	for i := range keys {
		combined[keys[i]] = values[i]
	}

	return nil, keys, values, combined
}

func checkForSlash(slashString string) string {
	if slashString[len(slashString)-1:] == "/" {
		slashString = slashString[:len(slashString)-1]
	}

	return slashString
}
