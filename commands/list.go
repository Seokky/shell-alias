package commands

import (
	"fmt"
	"shell-alias/explorer"
)

func List() {
	lines := explorer.ReadShellFile()

	for _, line := range lines {
		fmt.Println(line)
	}
}