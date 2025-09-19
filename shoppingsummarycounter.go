package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	i := 0
	n := len(str)

	for i < n {
		for i < n && str[i] == ' ' {
			i++
		}
		if i >= n {
			break
		}
		j := i
		for j < n && str[j] != ' ' {
			j++
		}
		word := str[i:j]
		counts[word]++
		i = j
	}

	return counts
}
