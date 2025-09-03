package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for r := 0; r <= 7; r++ {
		for j := 1; j <= 8; j++ {
			for k := 2; k <= 9; k++ {
				if r < j && j < k {
					z01.PrintRune(rune(r + '0'))
					z01.PrintRune(rune(j + '0'))
					z01.PrintRune(rune(k + '0'))

					if !(r == 7 && j == 8 && k == 9) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}
