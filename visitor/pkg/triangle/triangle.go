package triangle

import (
	"errors"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

type visitor interface {
	VisitTriangle(t Triangle) (msg string, err error)
}

// Triangle ...
type Triangle interface {
	GetSides() (a, b, c float64)
	CheckCorrect() (err error)
	Accept(v visitor) (msg string, err error)
}

type triangle struct {
	a float64
	b float64
	c float64
}

// GetSides return triangle's sides
func (t *triangle) GetSides() (a, b, c float64) {
	a = t.a
	b = t.b
	c = t.c
	return
}

// CheckCorrect checks right input sides
func (t *triangle) CheckCorrect() (err error) {
	if t.a+t.b <= t.c || t.a+t.c <= t.b || t.b+t.c <= t.a {
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
