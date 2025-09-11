package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	n := len(args)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if greaterASCII(args[j], args[j+1]) {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, s := range args {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func greaterASCII(a, b string) bool {
	la, lb := len(a), len(b)
	min := la
	if lb < min {
		min = lb
	}
	for i := 0; i < min; i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}

	return la > lb
}
