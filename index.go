package piscine

func Index(s string, toFind string) int {
	for index := range s {
		runes := []rune(s)
		toFindRunes := []rune(toFind)

		if string(runes[index]) == toFind {
			return index
		} else if len([]rune(toFind)) > 0 && toFindRunes[0] == runes[index] {
			return index
		}
	}
	return -1
}
