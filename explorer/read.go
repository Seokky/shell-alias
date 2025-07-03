package explorer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func ReadShellFile() []string {
	homepath, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(path.Join(homepath, ".zshrc"))
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