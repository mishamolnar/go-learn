package reflection

import (
	"fmt"
	"reflect"
)

func Playground() {
	v := reflect.ValueOf(9.0)
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)
}
