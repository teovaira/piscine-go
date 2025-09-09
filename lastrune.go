package piscine

func LastRune(s string) rune {
	length := len([]byte(s))
	runes := []rune(s)
	return runes[length-1]
}
