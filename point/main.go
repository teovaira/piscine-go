package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(nbr int) {
	if nbr == 0 {
		z01.PrintRune('0')
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	var digits []rune
	for nbr > 0 {
		d := nbr % 10
		digits = append([]rune{rune(d +'0')}, digits...)
		nbr/=10
	}
	for _,r := range digits {
		z01.PrintRune(r)
	}

}

func main() {
	points := point{}

	setPoint(&points)
	printString("x = ")
	printInt(points.x)
	printString(", y = ")
	printInt(points.y)

	
}
