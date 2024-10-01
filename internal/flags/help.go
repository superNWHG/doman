package flags

import (
	"flag"
	"fmt"
)

func getHelp(
	new flag.FlagSet,
	init flag.FlagSet,
	add flag.FlagSet,
	read flag.FlagSet,
	sync flag.FlagSet,
	link flag.FlagSet,
) {
	fmt.Println("doman is a tool to manage your dotfiles")
	fmt.Println("\nUsage:")
	fmt.Println("doman [subcommand] [flag(s)]")
	fmt.Println("\nSubcommands:")
	fmt.Println("\nnew - Create a new dotfiles repository")
	new.PrintDefaults()
	fmt.Println("\ninit - Init in an existing repository")
	init.PrintDefaults()
	fmt.Println("\nadd - Add a new configuration")
	add.PrintDefaults()
	fmt.Println("\nread - See what configurations are being tracked")
	read.PrintDefaults()
	fmt.Println("\nsync - Sync configurations with the remote repository")
	sync.PrintDefaults()
	fmt.Println("\nlink - Create a symlink for all files that do not have a symlink yet")
	link.PrintDefaults()
}
