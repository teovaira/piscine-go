package piscine

func FindNextPrime(nb int) int {
	for i := 3; i*i <= nb; i += 2 {
		if nb <= 1 && nb%2 == 0 && nb%i == 0 {
			return 0
		} else if nb == 2 {
			return 3
		} else {
			for j := nb; j >= nb; j++ {
				if j%2 != 0 {
					return j
				}
			}
		}
	}

	return j
}
