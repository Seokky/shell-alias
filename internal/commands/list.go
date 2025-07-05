package commands

import (
	"fmt"
	"os"
	"shell-alias/internal/explorer"
	"strings"
)

// Print aliases user have in .shellrc file
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
			printLine(line, number)
			number++
		}
	}

	fmt.Println("")
}

// Print alias line with index
func printLine(line string, number int) {
	fmt.Printf("    %d. ", number)
	fmt.Print(line[6:])
}
