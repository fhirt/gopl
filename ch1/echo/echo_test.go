package echo

import "testing"

func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forLoop(args)
	}
}

func BenchmarkRangeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeLoop(args)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoin(args)
	}
}

var args = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"}
