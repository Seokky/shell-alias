package explorer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadShellFile(shellFilePath string) []string {
	file, err := os.Open(shellFilePath)

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

	reader := bufio.NewReader(file)
	lines := make([]string, 0, 10)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				fmt.Println("read file error:", err)
				os.Exit(1)
			}
			break
		}

		lines = append(lines, line)
	}

	return lines
}