package piscine

func BasicAtoi2(s string) int {
	n := 0

	for _, r := range s {
		d := int(r - '0')
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + d
	}
	return n
}
