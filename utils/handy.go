package utils

import "fmt"

func init() {
	fmt.Println("Handy intialized")
}

const (
	nine int     = 9
	Ten  float32 = 10.1
	lie  bool    = false
)

var mutable int = 111
var ExportedMutable = 1111

func Foo() {
	fmt.Println("bar")
}
