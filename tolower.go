package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for index, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			runes[index] += 32
		}
	}
	return string(runes)
}
