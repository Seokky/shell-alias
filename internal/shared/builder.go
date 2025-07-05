package shared

import (
	"fmt"
	"strings"
)

// Construct alias line in .shellrc file format
func AliasToLine(name, command string) string {
	return fmt.Sprintf("alias %s='%s'", name, command)
}

// Parse alias from unformatted .shellrc file line
func LineToAlias(line string) ParsedAlias {
	line = line[6:]
	parts := strings.Split(line, "=")

	return ParsedAlias{
		Name: parts[0],
		Line: line,
	}
}

type ParsedAlias struct {
	Name string
	Line string
}
