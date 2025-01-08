package gitcredentials

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"syscall"

	"golang.org/x/term"
)

func AskGitCredentials() (name string, mail string, pass string, err error) {
	fmt.Print("Name: ")
	if _, err = fmt.Scan(&name); err != nil {
		return
	}

	fmt.Print("Email: ")
	if _, err = fmt.Scan(&mail); err != nil {
		return
	}

	fmt.Print("Password:")
	bytePass, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return
	}
	pass = string(bytePass)

	return
}

func GetGitCredentials(gitUrl string) (name string, mail string, password string, error error) {
	credentialFile := filepath.Join(os.Getenv("HOME"), ".git-credentials")
	gitconfigFile := filepath.Join(os.Getenv("HOME"), ".gitconfig")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		credentialContent, err := os.ReadFile(credentialFile)
		if err != nil {
			error = err
			return
		}

		rg := regexp.MustCompile("https?://(.+):(.+)@" + gitUrl)
		if credSlice := rg.FindAllStringSubmatch(string(credentialContent), -1)[0]; len(credSlice) < 3 {
			error = errors.New("Could not find username and password in credential file")
			return
		} else {
			name = credSlice[1]
			password = credSlice[2]
		}
	}()

	go func() {
		defer wg.Done()
		gitconfigContent, err := os.ReadFile(gitconfigFile)
		if err != nil {
			error = err
			return
		}

		rg := regexp.MustCompile("email ?= ?(.+)")
		if mailSlice := rg.FindAllStringSubmatch(string(gitconfigContent), -1)[0]; len(mailSlice) < 2 {
			error = errors.New("Could not find email in gitconfig file")
			return
		} else {
			mail = mailSlice[1]
		}
	}()

	wg.Wait()
	if error != nil {
		return
	}

	return
}
