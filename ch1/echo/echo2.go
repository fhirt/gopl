// Echo2 prints its command-line arguments.
package echo

func rangeLoop(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}
