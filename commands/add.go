package commands

import (
	"bufio"
	"fmt"
	"os"
	"shell-alias/shared"
)

/*
	1. TODO check if exists already
	2. update via source shellFilePath
	3. BUG shell-alias list not out new alias
*/

func Add(shellFilePath, name, command string) {
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
	fmt.Println("Pasting into", shellFilePath)
	fmt.Println("")

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(line + "\n")

	if err != nil {
		fmt.Println("write to file error:", err)
		os.Exit(1)
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("flush buffer error:", err)
		os.Exit(1)
	}

	fmt.Printf("Done! Run \"source %s\" to update env immediately!", shellFilePath)
	fmt.Println("")
	fmt.Println("")
}