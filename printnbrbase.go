package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseRunes := []rune(base)
	b := len(baseRunes)

	if nbr == 0 {
		z01.PrintRune(baseRunes[0])
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
	}

	n := nbr
	if n > 0 {
		n = -n
	}

	printNegInBase(n, b, baseRunes)
}

func printNegInBase(n int, b int, baseRunes []rune) {
	q := n / b

	r := -(n % b)

	if q != 0 {
		printNegInBase(q, b, baseRunes)
	}
	z01.PrintRune(baseRunes[r])
}

func isValidBase(base string) bool {
	br := []rune(base)

	if len(br) < 2 {
		return false
	}

	for i := 0; i < len(br); i++ {
		if br[i] == '+' || br[i] == '-' {
			return false
		}
		for j := i + 1; j < len(br); j++ {
			if br[i] == br[j] {
				return false
			}
		}
	}
	return true
}
