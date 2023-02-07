package main

import (
	"strings"
)

func stringsJoin(args []string) string {
	return strings.Join(args[1:], " ")
}
