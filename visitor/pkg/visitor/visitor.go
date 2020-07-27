package visitor

import (
	"errors"
	"fmt"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
	"github.com/konvulcius/patternsWB/visitor/pkg/circle"
	"github.com/konvulcius/patternsWB/visitor/pkg/square"
	"github.com/konvulcius/patternsWB/visitor/pkg/triangle"
)

// Visitor ...
type Visitor interface {
	VisitTriangle(t triangle.Triangle) (msg string, err error)
	VisitSquare(s square.Square) (msg string, err error)
	VisitCircle(c circle.Circle) (msg string, err error)
}

type element struct {
}

// VisitTriangle check triangle, and notify can build square or not
func (e *element) VisitTriangle(t triangle.Triangle) (msg string, err error) {
	err = t.CheckCorrect()
	if err != nil {
		return
	}
	if !t.CheckStraightIsoscelesTriangle() {
		err = errors.New(v1.CantCreateSquare)
		return
	}

	// choose square's side
	a, b, c := t.GetSides()
	var squareSide float64
	switch {
	case a <= b:
		fallthrough
	case a <= c:
		squareSide = a
	default:
		squareSide = b
	}
	msg = fmt.Sprintf(v1.CanCreateSquare, squareSide)
	return
}

// VisitSquare notify radius of circumscribed circle
func (e *element) VisitSquare(s square.Square) (msg string, err error) {
	err = s.CheckCorrect()
	if err != nil {
		return
	}
	radius := s.CircumscribedRadius()
	msg = fmt.Sprintf(v1.CanCreateCircle, radius)
	return
}

// VisitCircle notify simple parameters of circle
func (e *element) VisitCircle(c circle.Circle) (msg string, err error) {
	err = c.CheckCorrect()
	if err != nil {
		return
	}
	radius := c.GetRadius()
	length := c.Length()
	square := c.Square()
	msg = fmt.Sprintf(v1.CircleStatus, radius, length, square)
	return
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &element{}
}
