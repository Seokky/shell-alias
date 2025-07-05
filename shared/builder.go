package shared

import (
	"fmt"
	"strings"
)

// TODO comment
func BuildAliasLine(name, command string) string {
	return fmt.Sprintf("alias %s='%s'", name, command)
}

type ParsedAlias struct {
	Name string
	Line string
}

// TODO comment
func ParseAliasFromLine(line string) ParsedAlias {
	line = line[6:]
	parts := strings.Split(line, "=")

	return ParsedAlias{
		Name: parts[0],
		Line: line,
	}
}