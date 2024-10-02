package main

import (
	"log"

	"github.com/superNWHG/doman/internal/flags"
)

func main() {
	if err := flags.SetSubcommands(); err != nil {
		log.Fatal(err)
	}
}
