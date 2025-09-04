package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if table[j] > table[j+1] {

				table[j], table[j+1] = table[j+1], table[j]
				swapped = true
			}
		}
		if !swapped {
		}
	}
}
