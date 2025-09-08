package piscine

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 && i*i == nb {
			return false
		}
	}

	return true
}
