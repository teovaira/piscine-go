package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	length := 0
	for _, arg := range args {
		length += len(arg) + 1
	}

	result := make([]byte, 0, length)

	for i, arg := range args {
		result = append(result, arg...)
		if i < len(args)-1 {
			result = append(result, '\n')
		}
	}

	return string(result)
}
