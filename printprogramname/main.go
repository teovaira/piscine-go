package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]

	start := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' || path[i] == '\\' {
			start = i + 1
		}
	}

	name := path[start:]
	for _, r := range name {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
