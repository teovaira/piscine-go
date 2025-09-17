package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int { return a - b } 

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 4, 1} 

	fmt.Println(piscine.IsSorted(f, a1)) 
	fmt.Println(piscine.IsSorted(f, a2)) 
	fmt.Println(piscine.IsSorted(f, a3)) 
}

