package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
	}

	printAbs(n)
}

func printAbs(n int) {
	if n <= -10 || n >= 10 {
		printAbs(n / 10)
	}

	d := n % 10
	if d < 0 {
		d = -d
	}

	z01.PrintRune(rune('0' + d))
}
