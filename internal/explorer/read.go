package explorer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Get all alines from .shellrc file
func ReadShellFile(shellFilePath string) []string {
	file, err := os.Open(shellFilePath)

	if err != nil {
		fmt.Println("open file error:", err)
		os.Exit(1)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("close file error:", err)
			os.Exit(1)
		}
	}()

	reader := bufio.NewReader(file)
	lines := []string{}

	for {
		line, err := reader.ReadString('\n')

		if (err != nil) && (err != io.EOF) {
			fmt.Println("read file error:", err)
			os.Exit(1)
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			lines = append(lines, line)
		}

		if err == io.EOF {
			break
		}
	}

	return lines
}
