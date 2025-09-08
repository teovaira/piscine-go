package piscine

// import "fmt"

// func QuadB(x, y int) {

// 	if x <= 0 || y <= 0 {
// 		return
// 	}

// 	for i := 0; i < x; i++ {
// 		if i == 0 {
// 			fmt.Print("/")
// 		} else if i == x-1 {
// 			fmt.Print("\\")
// 		} else {
// 			fmt.Print("*")
// 		}
// 	}

// 	fmt.Println()

// 	for j := 0; j < y-2; j++ {
// 		for i := 0; i < x; i++ {
// 			if i == 0 || i == x-1 {
// 				fmt.Print("*")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}

// 		fmt.Println()
// 	}

// 	if y >= 2 {
// 		for i := 0; i < x; i++ {
// 			if i == 0 {
// 				fmt.Print("\\")
// 			} else if i == x-1 {
// 				fmt.Print("/")
// 			} else {
// 				fmt.Print("*")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
