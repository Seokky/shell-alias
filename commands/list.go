package commands

import (
	"fmt"
	"os"
	"shell-alias/explorer"
	"strings"
)

// TODO comment
func List(shellFilePath string) {
	lines := explorer.ReadShellFile(shellFilePath)

	if len(lines) == 0 {
		fmt.Println("")
		fmt.Println("  You have no aliases yet.")
		fmt.Println("")
		fmt.Println("  Run \"shell-alias add --name='name' --command='command'\" to add the new one")
		fmt.Println("")
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("  List of your aliases:")
	fmt.Println("")

	number := 1
	for _, line := range lines {
		if strings.HasPrefix(line, "alias") {
			fmt.Printf("    %d. ", number)
			fmt.Print(printLine(line))
			number++
		}
	}

	fmt.Println("")
}

// TODO comment
func printLine(line string) string {
	return line[6:]
}