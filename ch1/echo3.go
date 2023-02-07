package ch1

import (
	"strings"
)

func stringsJoin(args []string) string {
	return strings.Join(args[1:], " ")
}
