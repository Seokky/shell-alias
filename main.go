package main

import (
	"os"
	"path/filepath"
	"shell-alias/commands"
	"shell-alias/explorer"
	"shell-alias/router"
)

/*
	TODO print some text after delete alias
	TODO create help command
	TODO add to ShellFilePath hint about flag to force shell file path (please create or specify)
	TODO comment all functions and package itself
	TODO resolve ~ in --path
	TODO check for race
	TODO refactor dir structure for guides
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
	shellFilePath, _ = filepath.Localize(forcedPath)

	if len(forcedPath) == 0 {
		shellFilePath = explorer.ShellFilePath()
	}
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
