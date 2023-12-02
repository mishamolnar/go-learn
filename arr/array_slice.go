package arr

import "fmt"

func EntryPoint() {
	sliceLiteralCapacity()
}

func sliceLiteralCapacity() {
	sl := []int{1, 2, 3, 4}
	fmt.Printf("Slice %v has length %v and capacity %v \n", sl, len(sl), cap(sl))
	sl = append(sl, 12)
	fmt.Printf("Slice %v has length %v and capacity %v \n", sl, len(sl), cap(sl))
}

func sliceArrayComparison() {
	aNums := [...]int{0, 1, 2, 3}
	sNums := aNums[:]
	fmt.Printf("Type of slice is %T, and array is %T \n", sNums, aNums)
	sNums = append(sNums, 32)
	fmt.Println(sNums)
	fmt.Println(aNums)
	for index, value := range sNums {
		fmt.Printf("Index is %v with value %v \n", index, value)
	}
	fmt.Println("DELETE from 2 from the slice ")
	sNums = append(sNums[0:2], sNums[3:]...)
	fmt.Println(sNums)
}

func arraysTest() {
	var test5 = [5]int{1, 2, 3, 4, 5}
	test4 := [...]int{1, 2, 3, 4}
	//test5 = test4 does not work
	fmt.Print(test4)
	fmt.Println(test5)
	fmt.Println("Default array")
	var testEmpty [10]int
	testEmpty[4] = 10
	fmt.Println(testEmpty)
	fmt.Println("Default string array")
	var testString [10]string
	testString[5] = "test"
	fmt.Println(testString)
}
