package gitcredentials

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func AskGitCredentials() (error, string, string, string) {
	fmt.Print("Name: ")
	var name string
	if _, err := fmt.Scan(&name); err != nil {
		return err, "", "", ""
	}

	fmt.Print("Email: ")
	var mail string
	if _, err := fmt.Scan(&mail); err != nil {
		return err, "", "", ""
	}

	fmt.Print("Password:")
	pass, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err, "", "", ""
	}

	return nil, name, mail, string(pass)
}

func GetGitCredentials(url string) (string, string, string, error) {
	credentialFile := filepath.Join(os.Getenv("HOME"), ".git-credentials")

	credentialContent, err := os.ReadFile(credentialFile)
	if err != nil {
		return "", "", "", err
	}
	if !strings.Contains(string(credentialContent), url) {
		err := errors.New(url + " not in credential file")
		return "", "", "", err
	}

	credentials := strings.Split(string(credentialContent), "\n")

	var urlCredential string
	for i := range credentials {
		if strings.Contains(credentials[i], url) {
			urlCredential = string(credentials[i])
			break
		}
	}

	noUrl := strings.Replace(urlCredential, "https://", "", -1)
	noUrl = strings.Replace(noUrl, "@"+url, "", -1)

	var name, password string
	for i, v := range noUrl {
		if string(v) == ":" {
			name = noUrl[:i]
			password = noUrl[i+1:]
		}
	}

	gitconfigFile := filepath.Join(os.Getenv("HOME"), ".gitconfig")

	gitconfigContent, err := os.ReadFile(gitconfigFile)
	if err != nil {
		return "", "", "", err
	}
	if !strings.Contains(string(gitconfigContent), "email = ") {
		err := errors.New("email not in gitconfig file")
		return "", "", "", err
	}

	var mail string
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

	return name, mail, password, nil
}
