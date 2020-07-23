package circle

import (
	"errors"
	"math"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

type visitor interface {
	VisitCircle(c Circle) (msg string, err error)
}

// Circle ...
type Circle interface {
	GetRadius() (radius float64)
	Length() (len float64)
	Square() (s float64)
	CheckCorrect() (err error)
	Accept(v visitor) (msg string, err error)
}

type circle struct {
	radius float64
}

// GetRadius ...
func (c *circle) GetRadius() (radius float64) {
	return c.radius
}

// Length ...
func (c *circle) Length() (len float64) {
	len = math.Pi * c.radius * 2
	return
}

// Square ...
func (c *circle) Square() (s float64) {
	s = math.Pi * c.radius * c.radius
	return
}

// CheckCorrect check zero or negative input
func (c *circle) CheckCorrect() (err error) {
	if c.radius == 0 {
		err = errors.New(v1.ZeroRadius)
	} else if c.radius < 0 {
		err = errors.New(v1.NegativeRadius)
	}
	return
}

// Accept links circle with visitor
func (c *circle) Accept(v visitor) (msg string, err error) {
	msg, err = v.VisitCircle(c)
	return
}

// NewCircle ...
func NewCircle(radius float64) Circle {
	return &circle{
		radius: radius,
	}
}
