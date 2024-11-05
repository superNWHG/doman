package flags

import (
	"fmt"

	"github.com/superNWHG/doman/internal/git"
)

func getStatus(path string) error {

	status, err := git.Status(path)
	if err != nil {
		return err
	}
	fmt.Println(status)

	return nil
}
