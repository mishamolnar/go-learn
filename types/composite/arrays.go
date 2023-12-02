package composite

import "fmt"

func EntryPoint() {
	types()
}

func types() {
	a := []int{1, 2}
	b := [3]int{1, 2, 3}
	fmt.Printf("Type of a is %T and of b is %T", a, b) //Type of a is []int and of b is [3]int
}

func declaration() {
	a := [3]int{1, 2, 3}   //[1 2 3]
	b := [...]int{1, 2, 3} //[1 2 3]
	var c [3]int           //[0 0 0]
	d := [...]int{3: 4}    //[0 0 0 4]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
