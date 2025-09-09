package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for index, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			runes[index] -= 32
		}
	}
	return string(runes)
}
