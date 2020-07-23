package square

import (
	"errors"
	"math"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

type visitor interface {
	VisitSquare(s Square) (msg string, err error)
}

// Square ...
type Square interface {
	GetSide() (side float64)
	Diagonal() (d float64)
	CircumscribedRadius() (r float64)
	CheckCorrect() (err error)
	Accept(v visitor) (msg string, err error)
}

type square struct {
	side float64
}

// GetSide return square's side
func (s *square) GetSide() (side float64) {
	return s.side
}

// Diagonal return square's diagonal
func (s *square) Diagonal() (d float64) {
	d = math.Sqrt(2) * s.side
	return
}

// CircumscribedRadius return radius of circle that circumscribe square
func (s *square) CircumscribedRadius() (r float64) {
	r = s.Diagonal() / 2
	return
}

// CheckCorrect ...
func (s *square) CheckCorrect() (err error) {
	if s.side == 0 {
		err = errors.New(v1.ZeroSideSquare)
	} else if s.side < 0 {
		err = errors.New(v1.NegativeSideSquare)
	}
	return
}

// Accept links square with visitor
func (s *square) Accept(v visitor) (msg string, err error) {
	msg, err = v.VisitSquare(s)
	return
}

// NewSquare ...
func NewSquare(side float64) Square {
	return &square{
		side: side,
	}
}
