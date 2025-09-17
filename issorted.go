package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n <= 1 {
		return true
	}

	nonDecreasing := true
	nonIncreasing := true

	for i := 0; i < n-1; i++ {
		c := f(a[i], a[i+1])
		if c > 0 {
			nonDecreasing = false
		}
		if c < 0 {
			nonIncreasing = false
		}
	}

	return nonDecreasing || nonIncreasing
}
