package flow

import (
	"fmt"
	"math/rand"
)

func EntryPoint() {
	booleansTest()
}

func booleansTest() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)

}

func forRange() {
	arr := [5]int{10, 11, 12, 13, 589}
	for i, v := range arr {
		fmt.Printf("In index %d is number with value %d \n", i, v)
	}

	people := map[string]int{
		"Maria": 46,
		"Misha": 47,
	}

	for key, value := range people {
		fmt.Printf("%s is %d \n", key, value)
	}

	if val, ok := people["Misha"]; ok {
		fmt.Printf("Map contains person with age %d", val)
	}

	if val, ok := people["Misha"]; ok && val%2 == 1 {
		fmt.Printf("Misha is persent in the map and his age is odd")
	}
}

func statementStatementFlow() {
	values := make(map[string]int)
	values["number"] = int(rand.Int31n(200))
	if z, ok := values["number"]; ok && z > 100 {
		fmt.Printf("z is greater the 100, because it's value is %d \n", z)
	} else {
		fmt.Printf("z is less the 100, because it's value is %d \n", z)
	}
	z := 99
	fmt.Printf("z was initialized out of scope of conditional operator, with value %d and type %T", z, z)

}

func switchCase() {
	x := rand.Intn(300)
	fmt.Printf("Value of x is %d \t", x)
	switch {
	case x < 100:
		fmt.Println("Less then 100")
	case x < 200:
		fmt.Println("Between 100 and 200")
	case x < 300:
		fmt.Println("Between 200 and 300")
	case x >= 300:
		fmt.Println("More then 300")
	}
}

func switchCaseTwo() {
	for i := 0; i < 100; i++ {
		x, y := rand.Intn(10), rand.Intn(10)
		fmt.Printf("Value of x is %d and y is %d \n", x, y)
		switch {
		case x < 4 && y < 4:
			fmt.Println("Both less then 4")
		case x > 6 && y > 6:
			fmt.Println("Both greater then 6")
		case x >= 4 && x <= 6:
			fmt.Println("X is between 4 and 6")
		case y != 5:
			fmt.Println("Y is not 5")
		default:
			fmt.Println("No of this cases where met")
		}
	}
}

func forLoop() {
	i := 10
	for {
		fmt.Println(i)
		i++
		if i%10 == 0 {
			continue
		}
		switch {
		case i%2 == 0:
			fmt.Println("Even")
		default:
			fmt.Println("odd")
		}
		if i == 101 {
			break
		}
	}
}
