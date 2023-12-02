# Methods

### 1. What is the method?
It is a function that is attached to a specific type by specifying it in before the function name

### 2. What is method receiver and selector?
```go
func (p Point) Distance(q Point) float64 { //(p Point) - receiver
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func MainMethods() {
	a := Point{1, 2}
	b := Point{4, 6}
	fmt.Println(a.Distance(b)) //a.Distance - selector
}
```

### 3. Will this method print same addresses?
```go
func (p Point) Distance(q Point) float64 { //(p Point) - receiver
	fmt.Printf("Receiver address is %p \n", &p)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func MainMethods() {
	a := Point{1, 2}
	fmt.Printf("A address is %p \n", &a)
	a.Distance(Point{5, 6}) //a.Distance - selector
}
```

No, to use same address (for optimization or changing receiver) - pointer should be used

### 4. Name 3 possibilities when calling go method?
- Argument and receiver have the same type, like here `T` - `Point{1, 2}.Distance(q)`
- Argument is `T` and receiver is `*T` - then compiler takes the address `p.ScaleBy(2)` - implicit `(&p).ScaleBy(2)`
- Argyment is a type `*T` and receiver is `T` -> then implicit copy `pptr.Distance(q)` - implicit `(*pptr).Distance(q)` and pass by value

### 5. Can receiver be a `nil` value?
Yes, and it will work fine if this is handled in method
Example - map that returns for each key `nil` if the map is `nil`. 

### 6. Are methods promoted to the parent type?
Yes, by pointer and non-pointer objects:
```go
type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func MainMethods() {
	cp := ColoredPoint{Point: Point{1, 1}, Color: color.RGBA{R: 1, G: 1, B: 1, A: 1}}
	cp.Distance(&Point{1, 1}) //Color point has a method of Point
}
```
it is effectively same as `p.Point.Distance()`

### 7. Which method is picked when there are 2 Promoted with same name?
1 which is closer to the root named type

### 8. What is method value in Go? 
```go
package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func MainMethods() {
	cp := ColoredPoint{Point: Point{1, 1}, Color: color.RGBA{R: 1, G: 1, B: 1, A: 1}}
	distanceToCP := cp.Distance             //distance to CP - binded to a specific receiver - cp
	fmt.Printf("Type: %T \n", distanceToCP) //Type: func(*methods.Point) float64
	fmt.Println(distanceToCP(&Point{0, 0}))
}

func (p *Point) Distance(q *Point) float64 { //(p Point) - receiver
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

we can assign a method to a variable, and it will be bound to specific receiver. 

### 9. What is method expression in Go? 
```go
func MainMethods() {
	//cp := ColoredPoint{Point: Point{1, 1}, Color: color.RGBA{R: 1, G: 1, B: 1, A: 1}}
	scaleFunc := (*Point).Scale
	fmt.Printf("Type: %T \n", scaleFunc)             //Type: func(*methods.Point, int)
	fmt.Printf("Type: %T \n", (*ColoredPoint).Scale) //Type: func(*methods.ColoredPoint, int)
	scaleFunc(&Point{1, 2}, 10)
}

func (p *Point) Scale(factor int) {
    p.X *= float64(factor)
    p.Y *= float64(factor)
}
```
When we define `method expression` - new function is created with a receiver as a first argument and other arguments of a method then


### 10. 3 Reasons to use encapsulation in golang?
- User don't need to modify internal variables, so API is easier.
- It's a way to declare sufficient API and leave a room for implementation change.
- Prevent client from changing internal values of a type and maintain invariant of type easier. 