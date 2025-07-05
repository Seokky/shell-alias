package commands

import (
	"fmt"
	"os"
)

// Print help info
func Help() {
	fmt.Println("")
	fmt.Println("  Allowed commands:")
	fmt.Println("")
	fmt.Println("  0. shell-alias help")
	fmt.Println("  1. shell-alias list")
	fmt.Println("  2. shell-alias add --name='' --command=''")
	fmt.Println("  3. shell-alias delete --name=''")
	fmt.Println("")
	os.Exit(1)
}
