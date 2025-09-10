package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	total := 0
	for _, s := range strs {
		total += len(s)
	}
	total += (len(strs) - 1) * len(sep)
	buf := make([]byte, total)

	pos := 0
	for i, s := range strs {

		for j := 0; j < len(s); j++ {
			buf[pos] = s[j]
			pos++
		}

		if i < len(strs)-1 {
			for j := 0; j < len(sep); j++ {
				buf[pos] = sep[j]
				pos++
			}
		}
	}

	return string(buf)
}
