package router

import (
	"flag"
	"fmt"
	"os"
)

type Info struct {
	Cmd string
	Params []string
	ForcedPath string
}

func Command() Info {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("provide some command")
		os.Exit(1)
	}


	if args[0] == "list" {
		list := flag.NewFlagSet("list", flag.ExitOnError)
		path := list.String("path", "", "")

		list.Parse(flag.Args()[1:])

		return Info{
			Cmd: "list",
			Params: []string{},
			ForcedPath: *path,
		}
	}

	if args[0] == "add" {
		add := flag.NewFlagSet("add", flag.ExitOnError)
		name := add.String("name", "", "")
		cmd := add.String("command", "", "")
		path := add.String("path", "", "")

		add.Parse(flag.Args()[1:])

		return Info{
			Cmd: "add",
			Params: []string{*name, *cmd},
			ForcedPath: *path,
		}
	}

	panic("provide some command")
}