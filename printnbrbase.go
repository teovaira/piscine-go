package piscine

import "github.com/01-edu/z01"

// helper: validate base
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseRunes := []rune(base)
	baseLen := int64(len(baseRunes))
	n := int64(nbr)

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// recursive printing
	var printRec func(int64)
	printRec = func(num int64) {
		if num >= baseLen {
			printRec(num / baseLen)
		}
		z01.PrintRune(baseRunes[num%baseLen])
	}

	printRec(n)
}
