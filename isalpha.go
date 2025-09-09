package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if letter < '0' || (letter > '9' && letter < 'A') || (letter > 'Z' && letter < 'a') || letter > 'z' {
			return false
		}
	}

	return true
}
