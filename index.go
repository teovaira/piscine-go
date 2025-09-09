package piscine

func Index(s string, toFind string) int {
	for index := range s {
		runes := []rune(s)
		toFindRunes := []rune(toFind)

		if len(toFindRunes) == 0 {
			return 0
		}

		if len(toFindRunes) > 0 && toFindRunes[0] == runes[index] {
			return index
		}
	}

	return -1
}
