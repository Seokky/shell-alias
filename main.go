package main

import (
	"os"
	"shell-alias/internal/commands"
	"shell-alias/internal/explorer"
	"shell-alias/internal/router"
)

// Command to work with: "help" | "list" | "add" | "delete"
var command string

// Provided parameters values
// Allowed params: --name, --command
// See more in shell-alias/internal/router
var parameters []string

// Value from --path parameter
var forcedPath string

// .shellrc file path to work with
var shellFilePath string

/*
	Parse current command with parameters and assign it to variables listed above;
	calculate and initialize .shellrc file path to shellFilePath
*/
func init() {
	info := router.Command()
	command, parameters, forcedPath = info.Cmd, info.Params, info.ForcedPath
	shellFilePath = explorer.ShellFilePath(forcedPath)
}

// Call relevant command handler
func main() {
	switch command {
	case "help":
		commands.Help()
	case "list":
		commands.List(shellFilePath)
	case "add":
		commands.Add(shellFilePath, parameters[0], parameters[1])
	case "delete":
		commands.Delete(shellFilePath, parameters[0])
	default:
		os.Exit(1)
	}
}
