package piscine

import "unicode/utf8"

func StrLen(s string) int {
	// count := 0
	// for range s {
	// 	count++
	// }

	// return count

	// rs := []rune(s)
	// return len(rs)

	return utf8.RuneCountInString(s)
}
