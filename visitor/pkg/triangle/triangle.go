package triangle

import (
	"errors"
	"math"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

type visitor interface {
	VisitTriangle(t Triangle) (msg string, err error)
}

// Triangle ...
type Triangle interface {
	GetSides() (a, b, c float64)
	Square() (sqr float64)
	CheckStraightIsoscelesTriangle() (right bool)
	CheckCorrect() (err error)
	Accept(v visitor) (msg string, err error)
}

type triangle struct {
	a float64
	b float64
	c float64
}

// GetSides ...
func (t *triangle) GetSides() (a, b, c float64) {
	return t.a, t.b, t.c
}

// TriangleSquare ...
func (t *triangle) Square() (sqr float64) {
	// formula
	p := (t.a + t.b + t.c) / 2
	sqr = math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	sqr = math.Round(sqr)
	return
}

// CheckStraightIsoscelesTriangle ...
func (t *triangle) CheckStraightIsoscelesTriangle() (b bool) {
	aSquare := t.a * t.a
	bSquare := t.b * t.b
	cSquare := t.c * t.c
	aSquare = math.Round(aSquare)
	bSquare = math.Round(bSquare)
	cSquare = math.Round(cSquare)
	if (aSquare+bSquare == cSquare && t.a == t.b) ||
		(aSquare+cSquare == bSquare && t.a == t.c) ||
		(bSquare+cSquare == aSquare && t.b == t.c) {
		b = true
	}
	return
}

// CheckCorrect checks right input
func (t *triangle) CheckCorrect() (err error) {
	if t.a <= 0.0 || t.b <= 0.0 || t.c <= 0.0 {
		err = errors.New(v1.NegativeTriangle)
	} else if t.a+t.b <= t.c || t.a+t.c <= t.b || t.b+t.c <= t.a {
		err = errors.New(v1.WrongTriangle)
	}
	return
}

// Accept links triangle with visitor
func (t *triangle) Accept(v visitor) (msg string, err error) {
	msg, err = v.VisitTriangle(t)
	return
}

// NewTriangle ...
func NewTriangle(a, b, c float64) Triangle {
	return &triangle{
		a: a,
		b: b,
		c: c,
	}
}
