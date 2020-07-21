package square

import (
	"errors"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

type visitor interface {
	VisitSquare(s Square) (msg string, err error)
}

// Square ...
type Square interface {
	GetSides() (a, b, c, d float64)
	CheckCorrect() (err error)
	Accept(v visitor) (msg string, err error)
}

type square struct {
	a float64
	b float64
	c float64
	d float64
}

// GetSides return square's sides
func (s *square) GetSides() (a, b, c, d float64) {
	a = s.a
	b = s.b
	c = s.c
	d = s.d
	return
}

// CheckCorrect checks right input sides
func (s *square) CheckCorrect() (err error) {
	if s.a == 0.0 || s.a != s.b || s.b != s.c || s.c != s.d {
		err = errors.New(v1.WrongSquare)
	}
	return
}

// Accept links square with visitor
func (s *square) Accept(v visitor) (msg string, err error) {
	msg, err = v.VisitSquare(s)
	return
}

// NewSquare ...
func NewSquare(a, b, c, d float64) Square {
	return &square{
		a: a,
		b: b,
		c: c,
		d: d,
	}
}
