package composite

import "fmt"

func Entrypoint() {
	testPointer()
}

func testPointer() {
	a := []int{1, 2, 3}
	p := &a[1]
	fmt.Println(*p)
	var b [3]int
	fmt.Println(b[0])
}

func testEqual() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println(a, b)
	//fmt.Println(a == b) compile error : invalid operation: a == b (slice can only be compared to nil)
}

func reverseTest() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverseSlice(arr[:])
	fmt.Println(arr) //[5 4 3 2 1]
	reverseArr(&arr)
	fmt.Println(arr) //[1 2 3 4 5]
}

func reverseSlice(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseArr(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func outOfBound() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} //length 8
	twoFour := arr[2:4]
	fmt.Println(twoFour) // [2 3]
	twoSeven := twoFour[0:6]
	fmt.Println(twoSeven) // [2 3 4 5 6 7] because 2 (start) + 6 (end) = 8 and len(arr) = 8
	twoEight := twoFour[0:7]
	fmt.Println(twoEight) //panic: runtime error: slice bounds out of range [:7] with capacity 6
}
