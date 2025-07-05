package main

import (
	"os"
	"shell-alias/commands"
	"shell-alias/explorer"
	"shell-alias/router"
)

/*
	TODO create help command
	TODO resolve ~ in --path
	TODO refactor dir structure for guides
	TODO comment all functions and package itself
	TODO check for race
	TODO readme and adding to PATH
*/

// TODO comment
var command string
// TODO comment
var parameters []string
// TODO comment
var forcedPath string
// TODO comment
var shellFilePath string

// TODO comment
func init() {
	data := router.Command()
	command, parameters, forcedPath = data.Cmd, data.Params, data.ForcedPath
	shellFilePath = explorer.ShellFilePath(forcedPath)
}

// TODO comment
func main() {
	switch command {
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
