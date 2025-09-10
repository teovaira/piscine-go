package piscine

func BasicJoin(elems []string) string {
	if len(elems) == 0 {
		return ""
	}

	total := 0
	for _, s := range elems {
		total += len(s)
	}

	buf := make([]byte, total)

	pos := 0
	for _, s := range elems {
		for i := 0; i < len(s); i++ {
			buf[pos] = s[i]
			pos++
		}
	}

	return string(buf)
}
