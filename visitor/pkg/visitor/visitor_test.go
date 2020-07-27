package visitor

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/visitor/pkg/circle"
	"github.com/konvulcius/patternsWB/visitor/pkg/square"
	"github.com/konvulcius/patternsWB/visitor/pkg/triangle"
)

const (
	testVisitTriangle = "test VisitTriangle"
	testVisitSquare   = "test VisitSquare"
	testVisitCircle   = "test VisitCircle"

	checkCorrect                   = "CheckCorrect"
	triangleGetSides               = "GetSides"
	circumscribedRadius            = "CircumscribedRadius"
	getRadius                      = "GetRadius"
	methodLength                   = "Length"
	methodSquare                   = "Square"
	checkStraightIsoscelesTriangle = "CheckStraightIsoscelesTriangle"
	canCreateSquare                = "you can create a square by doubling the triangle! It side is 4.00\n"
	canCreateCircle                = "you can create a circle circumscribed square! It radius is 4.00\n"
	circleStatus                   = "Radius: 6.00\n\tLength: 37.70\n\tSquare: 113.10\n"
)

var (
	cathetus   = 4.0
	hypotenuse = math.Sqrt(2) * cathetus
)

func TestVisitor_VisitTriangle(t *testing.T) {
	t.Run(testVisitTriangle, func(t *testing.T) {
		triangle := new(triangle.MockTriangle)
		triangle.On(checkCorrect).
			Return(nil).
			Once()
		triangle.On(triangleGetSides).
			Return(cathetus, cathetus, hypotenuse).
			Once()
		triangle.On(checkStraightIsoscelesTriangle).
			Return(true).
			Once()
		visitor := NewVisitor()
		msg, err := visitor.VisitTriangle(triangle)
		assert.NoError(t, err)
		assert.Equal(t, canCreateSquare, msg)
	})
}

func TestElement_VisitSquare(t *testing.T) {
	t.Run(testVisitSquare, func(t *testing.T) {
		square := new(square.MockSquare)
		square.On(checkCorrect).
			Return(nil).
			Once()
		square.On(circumscribedRadius).
			Return(4.0).
			Once()
		visitor := NewVisitor()
		msg, err := visitor.VisitSquare(square)
		assert.NoError(t, err)
		assert.Equal(t, canCreateCircle, msg)
	})
}

func TestElement_VisitCircle(t *testing.T) {
	t.Run(testVisitCircle, func(t *testing.T) {
		circle := new(circle.MockCircle)
		circle.On(checkCorrect).
			Return(nil).
			Once()
		circle.On(getRadius).
			Return(6.0).
			Once()
		circle.On(methodLength).
			Return(37.70).
			Once()
		circle.On(methodSquare).
			Return(113.10).
			Once()
		visitor := NewVisitor()
		msg, err := visitor.VisitCircle(circle)
		assert.NoError(t, err)
		assert.Equal(t, circleStatus, msg)
	})
}
