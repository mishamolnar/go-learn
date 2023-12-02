package methods

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
	//cp := ColoredPoint{Point: Point{1, 1}, Color: color.RGBA{R: 1, G: 1, B: 1, A: 1}}
	scaleFunc := (*Point).Scale
	fmt.Printf("Type: %T \n", scaleFunc)             //Type: func(*methods.Point, int)
	fmt.Printf("Type: %T \n", (*ColoredPoint).Scale) //Type: func(*methods.ColoredPoint, int)
	scaleFunc(&Point{1, 2}, 10)
}

func (p *Point) Distance(q *Point) float64 { //(p Point) - receiver
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) Scale(factor int) {
	p.X *= float64(factor)
	p.Y *= float64(factor)
}
