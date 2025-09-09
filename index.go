package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	target := []rune(toFind)

	if len(target) == 0 {
		return 0
	}
	if len(target) > len(runes) {
		return -1
	}

	for i := 0; i <= len(runes)-len(target); i++ {
		match := true
		for j := 0; j < len(target); j++ {
			if runes[i+j] != target[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
