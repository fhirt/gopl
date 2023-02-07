package main

import (
	"fmt"
	"gopl/ch3/basename"
)

func main() {

	files := []string{"a/b/c/d.go", "text.txt", "tst", "a/dir/this.too.exe"}

	for _, file := range files {
		fmt.Println(basename.Basename1(file))
		fmt.Println(basename.Basename2(file))
	}
}
