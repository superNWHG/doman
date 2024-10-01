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
	fmt.Println("\nnew")
	new.PrintDefaults()
	fmt.Println("\ninit")
	init.PrintDefaults()
	fmt.Println("\nadd")
	add.PrintDefaults()
	fmt.Println("\nread")
	read.PrintDefaults()
	fmt.Println("\nsync")
	sync.PrintDefaults()
	fmt.Println("\nlink")
	link.PrintDefaults()
}
