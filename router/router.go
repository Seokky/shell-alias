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

// TODO comment
func Command() Info {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("")
		fmt.Println("  Allowed commands:")
		fmt.Println("")
		fmt.Println("  1. shell-alias list")
		fmt.Println("  2. shell-alias add --name='' --command=''")
		fmt.Println("  3. shell-alias delete --name=''")
		fmt.Println("")
		fmt.Println("  [LOOK] By default program will scan your home directory to find shellrc file to work with")
		fmt.Println("  [LOOK] If not found, it will be created automatically with name you have entered")
		fmt.Println("  [LOOK] Each command can accept --path='' argument to force path of your shellrc file")
		fmt.Println("")
		os.Exit(1)
	}


	if args[0] == "list" {
		list := flag.NewFlagSet("list", flag.ExitOnError)
		path := list.String("path", "", "[opt] Forced path of shellrc file")

		list.Parse(flag.Args()[1:])

		return Info{
			Cmd: "list",
			Params: []string{},
			ForcedPath: *path,
		}
	}

	if args[0] == "add" {
		add := flag.NewFlagSet("add", flag.ExitOnError)
		name := add.String("name", "", "[req] Name of alias you want to add")
		cmd := add.String("command", "", "[req] Command of alias you want to add")
		path := add.String("path", "", "[opt] Forced path of shellrc file")

		add.Parse(flag.Args()[1:])

		return Info{
			Cmd: "add",
			Params: []string{*name, *cmd},
			ForcedPath: *path,
		}
	}

	if args[0] == "delete" {
		add := flag.NewFlagSet("delete", flag.ExitOnError)
		name := add.String("name", "", "[req] Name of alias you want to delete")
		path := add.String("path", "", "[opt] Forced path of shellrc file")

		add.Parse(flag.Args()[1:])

		return Info{
			Cmd: "delete",
			Params: []string{*name},
			ForcedPath: *path,
		}
	}

	panic("provide some command")
}