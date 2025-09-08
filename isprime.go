package piscine

func IsPrime(nb int) bool {
	if nb <= 0 {
		return false
	}

	for i := 2; i < nb; i++ {
		if nb%i == 0 && i*i == nb {
			return false
		}
	}

	return true
}
