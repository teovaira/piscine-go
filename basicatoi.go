package piscine

func BasicAtoi(s string) int {
	n := 0
	for _, r := range s {
		d := int(r - '0')
		n = n*10 + d
	}
	return n
}
