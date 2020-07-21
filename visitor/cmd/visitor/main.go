package main

import (
	"fmt"

	"github.com/konvulcius/patternsWB/visitor/pkg/square"
	"github.com/konvulcius/patternsWB/visitor/pkg/triangle"
	"github.com/konvulcius/patternsWB/visitor/pkg/visitor"
)

const (
	prefixTriangle = "triangle: "
	prefixSquare   = "square: "
	triangleSideA  = 3.0
	triangleSideB  = 4.0
	triangleSideC  = 5.0
	squareSide     = 11.0
	op             = "sqrt"
)

func main() {
	triangle := triangle.NewTriangle(triangleSideA, triangleSideB, triangleSideC)
	square := square.NewSquare(squareSide, squareSide, squareSide, triangleSideC)
	visitor := visitor.NewVisitor(op)

	var (
		msg string
		err error
	)
	// visit triangle
	msg, err = triangle.Accept(visitor)
	if err != nil {
		msg = err.Error()
	}
	fmt.Println(prefixTriangle, msg)
	//visit square
	msg, err = square.Accept(visitor)
	if err != nil {
		msg = err.Error()
	}
	fmt.Println(prefixSquare, msg)
}
