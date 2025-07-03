package router

import (
	"flag"
	"fmt"
	"os"
)

func Command() string {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("provide some command")
		os.Exit(1)
	}

	if args[0] == "list" {
		return "list"
	}

	fmt.Println("provide some command")
	os.Exit(1)
	return ""
}