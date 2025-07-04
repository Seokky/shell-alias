package commands

import (
	"fmt"
	"shell-alias/explorer"
	"strings"
)

func List(shellFilePath string) {
	lines := explorer.ReadShellFile(shellFilePath)

	fmt.Println("")
	fmt.Println("List of your aliases:")
	fmt.Println("")

	number := 1
	for _, line := range lines {
		if strings.HasPrefix(line, "alias") {
			fmt.Printf("%d. ", number)
			fmt.Print(printLine(line))
			number++
		}
	}

	fmt.Println("")
}

func printLine(line string) string {
	return line[6:]
}