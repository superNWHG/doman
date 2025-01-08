package gitcredentials

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		if !strings.Contains(string(credentialContent), gitUrl) {
			error = errors.New(gitUrl + " not in credential file")
			return
		}

		credentials := strings.Split(string(credentialContent), "\n")

		var urlCredential string
		for i := range credentials {
			if strings.Contains(credentials[i], gitUrl) {
				urlCredential = string(credentials[i])
				break
			}
		}

		noUrl := strings.Replace(urlCredential, "https://", "", -1)
		noUrl = strings.Replace(noUrl, "@"+gitUrl, "", -1)

		for i, v := range noUrl {
			if string(v) == ":" {
				name = noUrl[:i]
				password = noUrl[i+1:]
			}
		}
	}()

	go func() {
		defer wg.Done()
		gitconfigContent, err := os.ReadFile(gitconfigFile)
		if err != nil {
			error = err
			return
		}
		if !strings.Contains(string(gitconfigContent), "email = ") {
			error = errors.New("email not in gitconfig file")
			return
		}

		gitConfigLines := strings.Split(string(gitconfigContent), "\n")
		for i := range gitConfigLines {
			if strings.Contains(gitConfigLines[i], "email = ") {
				for j, v := range gitConfigLines[i] {
					if string(v) == "=" {
						mail = gitConfigLines[i][j+2:]
					}
				}
			}
		}
	}()

	wg.Wait()
	if error != nil {
		return "", "", "", error
	}

	return name, mail, password, nil
}
