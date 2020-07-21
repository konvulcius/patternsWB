package visitor

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
	"github.com/konvulcius/patternsWB/visitor/pkg/square"
	"github.com/konvulcius/patternsWB/visitor/pkg/triangle"
)

// Visitor ...
type Visitor interface {
	VisitTriangle(t triangle.Triangle) (msg string, err error)
	VisitSquare(s square.Square) (msg string, err error)
}

type element struct {
	op string
}

// VisitTriangle work with triangle's sides and do some op
func (e *element) VisitTriangle(t triangle.Triangle) (msg string, err error) {
	err = t.CheckCorrect()
	if err != nil {
		return
	}
	sideA, sideB, sideC := t.GetSides()
	switch {
	case strings.EqualFold(e.op, v1.OpSqrt):
		{
			perimeter := sideA + sideB + sideC
			res := math.Sqrt(perimeter)
			msg = v1.SqrtSum + strconv.FormatFloat(res, 'f', 2, 64)
		}
	case strings.EqualFold(e.op, v1.OpPow):
		{
			res := sideA*sideA + sideB*sideB + sideC*sideC
			msg = v1.SumPow + strconv.FormatFloat(res, 'f', 2, 64)
		}
	default:
		{
			err = errors.New(v1.UnknownOp)
		}
	}
	return
}

// VisitSquare work with square's sides and do some op
func (e *element) VisitSquare(s square.Square) (msg string, err error) {
	err = s.CheckCorrect()
	if err != nil {
		return
	}
	sideA, _, _, _ := s.GetSides()
	switch {
	case strings.EqualFold(e.op, v1.OpSqrt):
		{
			// не думмаю что количество сторон в квадрате поменяется
			perimeter := sideA * 4
			res := math.Sqrt(perimeter)
			msg = v1.SqrtSum + strconv.FormatFloat(res, 'f', 2, 64)
		}
	case strings.EqualFold(e.op, v1.OpPow):
		{
			// аналогично
			res := sideA * sideA * 4
			msg = v1.SumPow + strconv.FormatFloat(res, 'f', 2, 64)
		}
	default:
		{
			err = errors.New(v1.UnknownOp)
		}
	}
	return
}

// NewVisitor ...
func NewVisitor(op string) Visitor {
	return &element{
		op: op,
	}
}
