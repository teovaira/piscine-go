package piscine

func NRune(s string, n int) rune {
	if n >= 0 && n < len([]byte(s)) {
		runes := []rune(s)
		return runes[n-1]
	}

	return 0
}
