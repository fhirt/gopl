package echo

import (
	"strings"
)

func stringsJoin(args []string) string {
	return strings.Join(args[1:], " ")
}
