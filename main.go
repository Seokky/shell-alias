package main

import (
	"os"
	"path/filepath"
	"shell-alias/commands"
	"shell-alias/explorer"
	"shell-alias/router"
)

/*
	TODO create command to delete user alias and update shell
	TODO create help command
	TODO what if creating alias already exists (user can write Y to stdin)
	TODO know how to install this package on remote machine without installing go and cloning from github
	TODO add to ShellFilePath hint about flag to force shell file path (please create or specify)
	TODO comment all functions and package itself
	TODO placeholder no aliases yet
	TODO resolve ~ in --path
*/

var command string
var parameters []string
var forcedPath string
var shellFilePath string

func init() {
	data := router.Command()
	command, parameters, forcedPath = data.Cmd, data.Params, data.ForcedPath
	shellFilePath, _ = filepath.Localize(forcedPath)

	if len(forcedPath) == 0 {
		shellFilePath = explorer.ShellFilePath()
	}
}

func main() {
	switch command {
	case "list":
		commands.List(shellFilePath)
	case "add":
		commands.Add(shellFilePath, parameters[0], parameters[1])
	default:
		os.Exit(1)
	}
}
