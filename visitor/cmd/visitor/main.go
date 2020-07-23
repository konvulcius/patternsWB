package main

import (
	"fmt"
	"math"

	"github.com/konvulcius/patternsWB/visitor/pkg/circle"
	"github.com/konvulcius/patternsWB/visitor/pkg/square"
	"github.com/konvulcius/patternsWB/visitor/pkg/triangle"
	"github.com/konvulcius/patternsWB/visitor/pkg/visitor"
)

const (
	prefixTriangle = "triangle: "
	prefixSquare   = "square: "
	prefixCircle   = "circle: "
	triangleSide   = 4.0
	squareSide     = 4.0
	circleRadius   = 6.0
)

var (
	triangleSideC = math.Sqrt(2) * triangleSide
)

func main() {
	triangle := triangle.NewTriangle(triangleSide, triangleSide, triangleSideC)
	square := square.NewSquare(squareSide)
	circle := circle.NewCircle(circleRadius)
	visitor := visitor.NewVisitor()
	var (
		msg string
		err error
	)

	// visit triangle
	msg, err = triangle.Accept(visitor)
	if err != nil {
		msg = err.Error()
	}
	fmt.Print(prefixTriangle, msg)

	// visit square
	msg, err = square.Accept(visitor)
	if err != nil {
		msg = err.Error()
	}
	fmt.Print(prefixSquare, msg)

	// visit circle
	msg, err = circle.Accept(visitor)
	if err != nil {
		msg = err.Error()
	}
	fmt.Print(prefixCircle, msg)
}
