package functions

import (
	"fmt"
	"time"
)

func TestDefer() {
	fmt.Println(testPanic()) //prints 1
}

func testPanic() (a int) {
	defer func() {
		if p := recover(); p != nil {
			a = 1
		}
	}()
	panic("ABCD")
}

func deferChangeResult() (a int) {
	defer func() { fmt.Println("After panic") }()
	panic("Test")
	return 1
}

func testDeferWithSleep() {
	deferWithSleep()
	fmt.Println("After deferWithSleep")
}

func deferWithSleep() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Println("defer execution")
	}()
	fmt.Println("deferWithSleep")
}

func deferExecutionOrderq() {
	a := 1
	defer fmt.Printf("defer 1, value %v \n", a)
	a = 1
	defer fmt.Printf("defer 2, value %v \n", a)
	sideFunction(3)
	a = 4
	defer fmt.Printf("defer 4, value %v \n", a)
	panic("acd")
}

func sideFunction(a int) error {
	defer fmt.Printf("defer in method, value %v \n", a)
	err := fmt.Errorf("error in fucntion %v \n", "sideFunction")
	return err
}

//defer in method, value 3
//defer 4
//defer 2
//defer 1
