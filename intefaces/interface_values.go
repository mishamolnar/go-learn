package intefaces

import (
	"bytes"
	"fmt"
	"io"
)

func InterVal() {
	testNull()
}
func testNull() {
	var buff *bytes.Buffer //buffer satisfies Writer
	f(buff)
}

func f(w io.Writer) {
	if w != nil {
		w.Write([]byte("test"))
	} else {
		fmt.Println("writer is NIL")
	}
}

func testComparison() {
	sl := []int{1, 2, 3}
	var x interface{} = sl
	//fmt.Println(sl == sl) //compile error
	fmt.Printf("Slices compared with result %v \n", x == x) //runtime panic
}
