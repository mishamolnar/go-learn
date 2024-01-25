# Questions composite data types

### 1. Ways to declare array? 
```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}   //[1 2 3]
	b := [...]int{1, 2, 3} //[1 2 3]
	var c [3]int           //[0 0 0]
	d := [...]int{3: 4}    //[0 0 0 4]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
```

### 2. Are arrays comparable?
If an array element type is comparable then whole array is comparable 

### 3. When array types are the same? 
When array element type and array length are same:
```go
package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := [3]int{1, 2, 3}
	fmt.Printf("Type of a is %T and of b is %T", a, b) //Type of a is []int and of b is [3]int
}
```

### 4. Name 3 main components of the slice?
Pointer - points to slices first element (not necessarily arrays first element)
Length - number of slice elements  
Capacity - number of buckets between pointer and end of underlying array

### 5. When slicing causes panic?
- When start < 0
- When end > cap(s) (but not when end > len(s))
example:

```go
package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} //length 8
	twoFour := arr[2:4]
	fmt.Println(twoFour) // [2 3]
	twoSeven := twoFour[0:6]
	fmt.Println(twoSeven) // [2 3 4 5 6 7] because 2 (start) + 6 (end) = 8 and len(arr) = 8
	twoEight := twoFour[0:7]
	fmt.Println(twoEight) //panic: runtime error: slice bounds out of range [:7] with capacity 6
}
```

### 6 How to modify array in a function? 
- Pass a pointer
- Create a slice of a whole array
Both functions work as expected, but slice is not bounded to array size
```go
package main

import "fmt"

func Entrypoint() {
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
```

### 7 Are slices comparable? 
No, the `==` operator compares only the reference of the slice, not the slice itself

### 8 How we can delete an element from the slice? 
Using copy function:
`copy(slice[i:], slice[i+1:])`
this deletes an element in the position i

### 9 When slice reference update is required?
When capacity or length of slice changes. Examples: append, delete

### 10 Is this valid to take an address of the map element? 
Example 
`
a := []int{1, 2, 3}
p := &a[1]
`

Answer: Yes, it will work as expected

### 11 How map keys are compared?
with `==` operator, that is why Map keys should be comparable by this operator 

`testMap := make(map[[]int]string)` - compile error invalid map key type []int

### 12. Why pointer to map element is an invalid operation? `&ages["bob"]`
Because rehashing might happen and the pointer to array element will be invalid

### 13. Explain long declaration of each aggregate type? (like `var arr [3]int`)
array: `arr[0] = 1` - ok, because size was allocated 
slice: `slice[0] = 1` - panic, because reference type
map: `ma[["Bob"] = 1` - panic, because reference type

### 14. How to check if value is present in the map?
`age, ok = ages["Bob"]` - ok is true only if map had "Bob" record

### 15. Is struct returned by function addressable?  
No, only if it is pointer to a struct
https://utcc.utoronto.ca/~cks/space/blog/programming/GoAddressableValues

### 16. Which fields of a struct are exported?
Only fields that start with Upper case letter

### 17. Will this code compile?
```go
package main

type Point struct {
	x, y int
}

func getPoint() Point {
	return Point{1, 2}
}

func toZero() {
	getPoint().x = 0
}
```

No, because getPoint().x is not addressable 

### 18. How to create Struct and Pointer to it?
`pp := &Point{1, 2}`

### 19. Can structs be compared?
Yes and it will give correct result if all the fields are comparable. Fields are compared in order

### 20. Explain anonymous fields in structs?
These are the fields that have no name when creating a struct and will be flattered. That means that object 
can refer to the nested object fields without mentioning intermediate field name. Anonymous fields have name 
that can be mentioned in object creation and the name will be same as anonymous object type. 
(Exporting also depends on exporting of type of anonymous filed)

Example:
```go
package composite

import (
	"fmt"
)

type Wheel struct {
	Circle
	Strokes int
}

type Circle struct {
	Point
	Radius int
}

type Point struct {
	x, y int
}

func printWheel() {
	wheel := Wheel{Circle{Point{1, 2}, 1}, 2}
	wheel.x = 100
	wheel.y = 100
	wheel.Radius = 1
	fmt.Println(wheel)
}
```

### 21. Can anonymous field be a pointer? 
Yes, it can be pointer or value

### 22. What will pring this code snippet? 
```go
package main

import "fmt"

func main() {
        var a [5]int
        b := a
        b[2] = 7
        fmt.Println(a, b)
}
```

It prints [0 0 0 0 0] [0 0 7 0 0], because array is a value type 
and line `b := a` declares a new variable b and copies all the contents of variable a

### 23. What will print code below? 
```go
package main

import "fmt"

func changeLength(sl []int) {
	sl[0] = -1
	sl = sl[1:]
	fmt.Printf("slice in changeLength %v \n", sl)
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	changeLength(sl)
	fmt.Printf("slice in main %v \n", sl)
}
```

this will print: 
```
slice in changeLength [2 3 4 5] 
slice in main [-1 2 3 4 5] 
```

Outcome: changeLength function has changed the first element of the slice,
but not the length. The reason why is that slice is the struct data type:
```
type slice struct {
        ptr   unsafe.Pointer
        len   int
        cap   int
}
```
And when we pass it to the function then it gets copied with new struct but same values. 
Because ptr is new and value of it is the same - we can modify slice elements.
But because len is not a pointer but just an int value - we modify it only for variable sl in function
`changeLength`. 


### 24. Can map contain a nil key?
Yes, and it will output correct result, but for this map key type should have nil value:
```go
package main

import (
	"fmt"
)

func main() {
	var m map[*string]int = make(map[*string]int)
	m[nil] = 1
	fmt.Println(m[nil]) //prints 1
}
```

### What will be if we assign to the nil map? To the nil pointer? Append to the slice?
Assign to the nil map causes panic
Assign to the nil slice will cause error because of index out of range
Append to the nil slice will behave as expected, will allocate memory and place the value on 0 index
