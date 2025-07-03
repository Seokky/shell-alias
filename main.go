package main

import (
	"fmt"
	"os"
	"shell-alias/commands"
	"shell-alias/router"
)

/*
	TODO create rc file if not exists
	TODO dynamic shell rc file name
	TODO pick only strings that starts with alias
	TODO beaty list format
	TODO write new alias to file
	TODO save file with new alias
	TODO update shell to alias became work (source ~/.zshrc)
	TODO create command to add alias and update shell
	TODO create command to view list of all user aliases
	TODO create command to delete user alias and update shell
	TODO create help command
	TODO make github repo
	TODO english comments
	TODO know how to install this package on remote machine without installing go and cloning from github
	TODO what if creating alias already exists (user can write Y to stdin)
*/

var command string

func init() {
	command = router.Command()
}

func main() {
	switch command {
	case "list":
		commands.List()
	default:
		fmt.Println(fmt.Errorf("unknown command %s", command))
		os.Exit(1)
	}

	// readerWriter := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))

	// for {
	// 	line, err := readerWriter.ReadString('\n')

	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println("read file error:", err)
	// 		}
	// 		break
	// 	}

	// 	fmt.Println("line:", line)
	// }
}
