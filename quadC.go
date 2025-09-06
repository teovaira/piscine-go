package piscine

import "fmt"

func QuadC(x,y int) {

	if x <= 0 || y <= 0 {
		return
	}

	for i := 1; i <= y; i++ {

		for j:= 1; j <= x; j++ {

			var ch rune

			isTop := i == 1
			isBottom := i == y
			isLeft := j == 1
			isRight := j == x

			switch {

			case isTop && isLeft:
				ch = 'A'
			case isTop && isRight:
				ch = 'A'
			case isTop:
				ch = 'B'

			case isBottom && isLeft:
				ch = 'C'
			case isBottom && isRight:
				ch = 'C'
			case isBottom:
				ch = 'B'

			case isLeft || isRight:
				ch = 'B'

			default:
				ch = ' '
			} 
			fmt.Print(string(ch))
		}

		fmt.Println()

	}

	


}