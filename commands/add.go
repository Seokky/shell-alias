package commands

import (
	"bufio"
	"fmt"
	"os"
	"shell-alias/explorer"
	"shell-alias/shared"
	"strings"
)

// TODO comment
func Add(shellFilePath, name, command string) {
	if isAliasExists(shellFilePath, name) {
		fmt.Println("")
		fmt.Println("  [ERROR] alias already exists!")
		fmt.Println("")
		os.Exit(1)
	}

	file, err := os.OpenFile(shellFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("open file error:", err)
		os.Exit(1)
	}

	defer func(){
		if err := file.Close(); err != nil {
			fmt.Println("close file error:", err)
			os.Exit(1)
		}
	}()

	line := shared.BuildAliasLine(name, command)

	fmt.Println("")
	fmt.Println("  Pasting into", shellFilePath)
	fmt.Println("")

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("\n" + line + "\n")

	if err != nil {
		fmt.Println("write to file error:", err)
		os.Exit(1)
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("flush buffer error:", err)
		os.Exit(1)
	}

	fmt.Printf("  Done! Run \"source %s\" to update env immediately!", shellFilePath)
	fmt.Println("")
	fmt.Println("")
}

// TODO comment
func isAliasExists(shellFilePath, name string) bool {
	lines := explorer.ReadShellFile(shellFilePath)

	number := 1
	for _, line := range lines {
		if strings.HasPrefix(line, "alias") {
			parsed := shared.ParseAliasFromLine(line)

			if parsed.Name == name {
				return true
			}

			number++
		}
	}

	return false
}