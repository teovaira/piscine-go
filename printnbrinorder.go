package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	strnb := string(n)
	strlen := len([]rune(strnb))

	if strlen == 1 {
		z01.PrintRune(rune(n + '0'))
	}

	table := make([]int, strlen)

	for i := 0; i < strlen; i++ {
		for j := 0; j < strlen-1; j++ {
			if table[j] == table[strlen-1] {
				break
			} else if table[j] > table[j+1] {
				table[j+1] = table[j]
			}
		}
	}
}
