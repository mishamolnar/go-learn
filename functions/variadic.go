package functions

import (
	"fmt"
	"strings"
)

func Variadic() {
	concat("a", "b", "c")
	//same as
	values := []string{"a", "b", "c"}
	concat(values...)
}

func concat(stringsArgs ...string) string {
	fmt.Printf("Type of string is %T", stringsArgs) //Type of string is []string
	return strings.Join(stringsArgs, " ")
}
