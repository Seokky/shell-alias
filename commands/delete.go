package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"shell-alias/shared"
	"strings"
)

// TODO comment
func Delete(shellFilePath, name string) {
	curFile, err := os.Open(shellFilePath)
	if err != nil {
		fmt.Println("open file error:", err)
		os.Exit(1)
	}
	defer func(){
		if err := curFile.Close(); err != nil {
			fmt.Println("close file error:", err)
			os.Exit(1)
		}
	}()

	tempNewShellFilePath := shellFilePath + "copy"

	_, err = os.Create(tempNewShellFilePath)
	if err != nil {
		fmt.Println("creating file error:", err)
		os.Exit(1)
	}

	newFile, err := os.OpenFile(tempNewShellFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("open new file error:", err)
		os.Exit(1)
	}
	defer func(){
		if err := newFile.Close(); err != nil {
			fmt.Println("close file error:", err)
			os.Exit(1)
		}
	}()


	reader := bufio.NewReader(curFile)
	writer := bufio.NewWriter(newFile)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("read file error:", err)
				os.Exit(1)
			}
			break
		}


		if strings.HasPrefix(line, "alias") {
			parsed := shared.ParseAliasFromLine(line)

			if parsed.Name == name {
				continue
			}
		}

		_, err = writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("write to file error:", err)
			os.Exit(1)
		}
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("flush buffer error:", err)
		os.Exit(1)
	}

	os.Rename(tempNewShellFilePath, shellFilePath)

	fmt.Println("")
	fmt.Printf("Alas \"%s\" successfully deleted from %s\n", name, shellFilePath)
	fmt.Println("")
}