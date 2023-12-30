package intefaces

import "fmt"

func TestTypeSwitch() {
	idiotFormatting(1)
}

func idiotFormatting(in interface{}) {
	switch x := in.(type) {
	case nil:
		fmt.Println("nil nil nil")
	case bool:
		fmt.Println(x, " is a value of boolean x")
	case int:
		b := max(x, 12)
		fmt.Println(b, " is an int")
	case string:
		fmt.Printf("This was a string %s", x)
	}
}
