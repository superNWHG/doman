package gitcredentials

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func GetGitCredentials() (error, string, string, string) {
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
