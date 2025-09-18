package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func DescendComb() {
	for r := 99; r >= 0; r-- {
		for j := r - 1; j >= 0; j-- {

			z01.PrintRune(rune(r/10 + '0'))
			z01.PrintRune(rune(r%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10 + '0'))
			z01.PrintRune(rune(j%10 + '0'))

			if !(r == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
	}
	z01.PrintRune('\n')
	os.Exit(1)
}
