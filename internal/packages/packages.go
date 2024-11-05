package packages

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

type packageManager struct {
	command          string
	searchFlag       string
	installFlag      string
	abortErrorString string
}

func search(p *packageManager, query string) ([]string, error) {
	cmd := exec.Command(p.command, p.searchFlag, query)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var packages []string
	packages = append(packages, lines...)
	return packages, nil
}

func install(p *packageManager, program string) error {
	cmd := exec.Command(p.command, p.installFlag, program)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil && err.Error() != p.abortErrorString {
		return err
	}

	return nil
}

func Search(os string, query string) (string, error) {
	var packages []string
	var err error
	switch os {
	case "arch":
		packages, err = search(&packageManager{command: "pacman", searchFlag: "-Ssq"}, query)
		if err != nil {
			return "", err
		}
	case "debian":
		packages, err = search(&packageManager{command: "apt-cache", searchFlag: "search"}, query)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("Unsupported OS")
	}

	installPkg, err := fuzzyfinder.Find(packages, func(i int) string {
		return packages[i]
	})
	if err != nil && err.Error() != "abort" {
		return "", err
	}
	return packages[installPkg], nil
}

func Install(os string, pkg string) error {
	switch os {
	case "arch":
		if err := install(&packageManager{command: "pacman", installFlag: "-S", abortErrorString: "exit status 1"}, pkg); err != nil {
			return err
		}
	case "debian":
		if err := install(&packageManager{command: "apt", installFlag: "install", abortErrorString: "Abort."}, pkg); err != nil {
			return err
		}
	default:
		return errors.New("Unsupported OS")
	}

	return nil
}
