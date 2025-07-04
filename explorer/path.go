package explorer

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func ShellFilePath() string {
	homepath, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	entries, err := os.ReadDir(homepath)

	if err != nil {
		fmt.Println("read home direactory error:", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}

		if entry.Name() == ".zshrc" {
			return path.Join(homepath, ".zshrc")
		}

		if entry.Name() == ".bashrc" {
			return path.Join(homepath, ".bashrc")
		}
	}

	fmt.Println("")
	fmt.Println("[WARN] No shell .rc files found at", homepath)
	fmt.Println("")
	fmt.Println("We will create it for you.")
	fmt.Println("If you on linux, probably you need .bashrc")
	fmt.Println("If you on MacOS, therefore you need .zshrc")
	fmt.Println("")

	UserInput:

	fmt.Print("Please, specify file name like .bashrc. .zshrc demand on your shell environment: ")

	var newShellFileName string
	_, err = fmt.Scanln(&newShellFileName)

	if err != nil {
		fmt.Println("scan user input error", err)
		os.Exit(1)
	}

	if !isValidShellFileName(newShellFileName) {
		fmt.Println("")
		fmt.Println("File name should starts with dot and ends with 'rc' like .bashrc, .zshrc.")
		fmt.Println("Try one more time, please.")
		fmt.Println("")

		goto UserInput
	}

	pathJoined := path.Join(homepath, newShellFileName)

	_, err = os.Stat(pathJoined)
	if err == nil {
		fmt.Println("")
		fmt.Println("[WARN]", pathJoined, "already exists")
		fmt.Println("")

		goto UserInput
	}

	_, err = os.Create(pathJoined)

	if err != nil {
		fmt.Println("error on creating file at", pathJoined)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("Successfully created at", pathJoined, "Now you can fill it like below:")
	fmt.Println("")
	fmt.Println("shell-alias add --name=\"sayhello\" --command=\"echo hello from alias!\"")
	fmt.Println("")

	os.Exit(1)
	return ""
}

func isValidShellFileName(name string) bool {
	if strings.HasPrefix(name, ".") && strings.HasSuffix(name, "rc") {
		return true
	}

	return false
}