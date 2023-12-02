package composite

import "fmt"

func MapEntryPoint() {
}

func testMapPointer() {
	ages := make(map[string]int)
	fmt.Println(ages)
	//pointer := &ages["bob"] //invalid operation: cannot take address of ages["bob"] (map index expression of type int)
}

func testSlicesAsAKey() {
	//testMap := make(map[[]int]string) //compile error invalid map key type []int
}
