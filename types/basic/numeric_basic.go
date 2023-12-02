package basic

import (
	"fmt"
)

/*
priority
- `* / % << >> & &^`
- `+ - | ^`
- `== != < <= > >=`
- `&& ||`
*/

const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func bitWise() {
	fmt.Println("________BITWISE_________")
	fmt.Println(2 >> 1)  //right shift on 1 bit -> 1
	fmt.Println(4 << 2)  //left shift on 2 bits -> 16
	fmt.Println(15 &^ 7) // and not -> 8
	fmt.Println(15 & 7)  // and  -> 7
	fmt.Println(15 | 7)  // or  -> 15
	fmt.Println(15 ^ 7)  // xor  -> 8
	fmt.Println(^255)    //every but is inverted. -> -2
	fmt.Printf("255 in binary %b, -2 in binary %b \n", 255, -256)
	fmt.Println("________END_BITWISE_________")
}

func priority() {
	fmt.Println("________PRIORITY_________")
	fmt.Println(3*2 ^ 2)   //same as (3 * 2) ^ 2 -> 4
	fmt.Println(2*3 | 1)   // same as (2 * 3) | 1 -> 7
	fmt.Println(15&^1 | 1) // same as (15 &^ 1) | 2 -> 15
	fmt.Println("________END_PRIORITY_________")
}

func specialCases() {
	fmt.Println("________SPECIAL_________")
	var z float64
	fmt.Println(z)
	fmt.Println(-z)
	fmt.Println(1 / z)
	fmt.Println(-1 / z)
	fmt.Println(z / z) //NAN
	fmt.Println("_______END_SPECIAL_________")
}
