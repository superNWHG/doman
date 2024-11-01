package packages

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

type packageManager struct {
	Command     string
	searchFlag  string
	installFlag string
}

func search(p *packageManager, query string) ([]string, error) {
	cmd := exec.Command(p.Command, p.searchFlag, query)

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
	cmd := exec.Command(p.Command, p.installFlag, program)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Search(os string, query string) (string, error) {
	switch os {
	case "arch":
		packages, err := search(&packageManager{Command: "pacman", searchFlag: "-Ssq"}, query)
		if err != nil {
			return "", err
		}

		installPkg, err := fuzzyfinder.Find(packages, func(i int) string {
			return packages[i]
		})
		if err != nil {
			return "", err
		}
		return packages[installPkg], nil
	default:
		return "", errors.New("Unsupported OS")
	}
}

func Install(os string, pkg string) error {
	switch os {
	case "arch":
		if err := install(&packageManager{Command: "pacman", installFlag: "-S"}, pkg); err != nil {
			return err
		}
	default:
		return errors.New("Unsupported OS")
	}

	return nil
}
