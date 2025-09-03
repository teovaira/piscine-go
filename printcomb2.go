package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for r := '0'; r <= '9'; r++ {
		for j := '0'; j <= '9'; j++ {
			if r != j {
				z01.PrintRune(r)
				z01.PrintRune(j)

				if !(r == '9' && j == '8') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
