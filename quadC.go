package piscine

// import "fmt"

// func QuadC(x, y int) {
// 	if x <= 0 || y <= 0 {
// 		return
// 	}

// 	for i := 1; i <= y; i++ {

// 		for j := 1; j <= x; j++ {

// 			var s string

// 			isTop := i == 1
// 			isBottom := i == y
// 			isLeft := j == 1
// 			isRight := j == x

// 			switch {

// 			case isTop && isLeft:
// 				s = "A"
// 			case isTop && isRight:
// 				s = "A"
// 			case isTop:
// 				s = "B"

// 			case isBottom && isLeft:
// 				s = "C"
// 			case isBottom && isRight:
// 				s = "C"
// 			case isBottom:
// 				s = "B"

// 			case isLeft || isRight:
// 				s = "B"

// 			default:
// 				s = ""
// 			}
// 			fmt.Print(s)
// 		}

// 		fmt.Println()

// 	}
// }
