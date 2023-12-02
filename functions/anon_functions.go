package functions

import (
	"fmt"
)

func TestSquaresTwo() {
	x1, sq1 := squares()
	fmt.Println(sq1())
	fmt.Println(sq1())
	fmt.Println(sq1())
	*x1 = 0
	fmt.Println(sq1())
}

func testSquares() {
	_, squareInstance := squares()
	fmt.Println(squareInstance()) //1
	fmt.Println(squareInstance()) //4
	fmt.Println(squareInstance()) //9
}

func squares() (*int, func() int) {
	x := 0
	return &x, func() int {
		x++
		return x * x
	}
}
