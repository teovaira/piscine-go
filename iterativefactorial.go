package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}

	num := 1
	for i := 1; i <= nb; i++ {
		num *= i
	}

	return num
}
