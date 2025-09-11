package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	upper := false
	start := 0
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		start = 1
	}

	printed := false

	for i := start; i < len(args); i++ {
		n, ok := atoiStrict(args[i])
		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
			printed = true
			continue
		}

		var r rune
		if upper {
			r = rune('A' + (n - 1))
		} else {
			r = rune('a' + (n - 1))
		}
		z01.PrintRune(r)
		printed = true
	}

	if printed {
		z01.PrintRune('\n')
	}
}

func atoiStrict(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	val := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		val = val*10 + int(r-'0')
	}
	return val, true
}
