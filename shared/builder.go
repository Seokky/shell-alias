package shared

import "fmt"

func BuildAliasLine(name, command string) string {
	return fmt.Sprintf("alias %s='%s'", name, command)
}