// Echo1 prints its command-line arguments.
package ch1

func forLoop(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}
