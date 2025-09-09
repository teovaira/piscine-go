package piscine

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'A' || r >= 'a') && (r <= 'Z' || r <= 'z') && r != ' ' {
			count++
		}
	}

	return count
}
