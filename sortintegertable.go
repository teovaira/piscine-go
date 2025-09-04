package piscine

func SortIntegerTable(table []int) {
	n := len(table) // how many boxes we have
	if n < 2 {      // 0 or 1 item is already sorted
		return
	}

	for i := 0; i < n-1; i++ { // do several passes
		swapped := false             // did we swap anything this pass?
		for j := 0; j < n-1-i; j++ { // walk neighbors: [j] and [j+1]
			if table[j] > table[j+1] {
				// swap the two neighbors
				table[j], table[j+1] = table[j+1], table[j]
				swapped = true
			}
		}
		if !swapped { // nothing moved → already sorted
			break
		}
	}
}
