package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	words := []string{}
	current := ""

	for _, r := range str {
		if r == ' ' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}

	if current != "" {
		words = append(words, current)
	}

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	return counts
}
