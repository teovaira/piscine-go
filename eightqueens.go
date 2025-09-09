package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	
	var cols [8]int
	var usedRow [8]bool       
	var usedDiagMain [15]bool 
	var usedDiagAnti [15]bool 

	backtrack(0, &cols, &usedRow, &usedDiagMain, &usedDiagAnti)
}

func backtrack(c int, cols *[8]int, usedRow *[8]bool, usedDiagMain *[15]bool, usedDiagAnti *[15]bool) {
	
	if c == 8 {
		printSolution(*cols)
		return
	}

	
	for r := 1; r <= 8; r++ {
		ri := r - 1          
		dm := (r - 1) - c + 7 
		da := (r - 1) + c     
		if !usedRow[ri] && !usedDiagMain[dm] && !usedDiagAnti[da] {
			
			cols[c] = r
			usedRow[ri] = true
			usedDiagMain[dm] = true
			usedDiagAnti[da] = true

			
			backtrack(c+1, cols, usedRow, usedDiagMain, usedDiagAnti)

			
			usedRow[ri] = false
			usedDiagMain[dm] = false
			usedDiagAnti[da] = false
		}
	}
}

func printSolution(cols [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune('0' + cols[i])) 
	}
	z01.PrintRune('\n')
}