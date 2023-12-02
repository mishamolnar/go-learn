package pointers

import "fmt"

func EntryPoint() {
	x := 32
	y := &x
	fmt.Printf("Reference is %v, of type %T \n", y, y)
	fmt.Printf("Dereferenced pointer is %v, of type %T \n", *y, *y)
	increment(y)
	fmt.Printf("Pointer original reference is %v, of type %T \n", x, x)
	p := Person{"Misha", 23, false}
	fmt.Println(p)
	fmt.Println("______PASS BY VALUE_____")
	birthday(p)
	fmt.Println(p)
	fmt.Println("______PASS BY POINTER VALUE_____")
	mutableBirthday(&p)
	fmt.Println(p)
}

func birthday(person Person) {
	person.age++
}

func mutableBirthday(person *Person) {
	//next statement is same as (*person).age++
	person.age++
}

func increment(num *int) {
	*num++
}

type Person struct {
	name    string
	age     int
	married bool
}
