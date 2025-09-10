package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			if newWord {
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				newWord = false
			} else {
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			newWord = true
		}
	}

	return string(runes)
}
