// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func CountDuplicates() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		counts[line]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}