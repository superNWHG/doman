package data

import (
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

func checkForSlash(slashString string) string {
	if slashString[len(slashString)-1:] == "/" {
		slashString = slashString[:len(slashString)-1]
	}

	return slashString
}
