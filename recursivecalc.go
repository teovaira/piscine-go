package piscine

func RecursiveCalc(num int) int {
	if num == 1 {
		return 1
	}

	if num > 1 {
		return num + RecursiveCalc(num-1)
	}

	return 0
}
