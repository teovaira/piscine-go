package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)

	current := ""
	inWord := false

	for _, r := range str {
		if r == ' ' {
			if inWord {
				counts[current]++
				current = ""
				inWord = false
			}
			continue
		}
		current += string(r)
		inWord = true
	}

	if inWord {
		counts[current]++
	}

	return counts
}
