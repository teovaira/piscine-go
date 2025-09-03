package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for r := 0; r <= 98; r++ {
		for j := 0; j <= 99; j++ {
			if r != j {
				z01.PrintRune(rune('0' + r/10))
				z01.PrintRune(rune(r%10 + '0'))
				z01.PrintRune(' ')
				z01.PrintRune(rune(j/10 + '0'))
				z01.PrintRune(rune(j%10 + '0'))

				if !(r == 98 && j == 99) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
