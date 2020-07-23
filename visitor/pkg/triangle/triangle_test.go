package triangle

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

const (
	testSquare                         = "test Square"
	testCheckStraightIsoscelesTriangle = "test CheckStraightIsoscelesTriangle"
	testTriangleCheckCorrectSuccess    = "test triangle CheckCorrect success"
	testTriangleCheckCorrectNegative   = "test triangle CheckCorrect negative"
	testTriangleCheckCorrectWrong      = "test triangle CheckCorrect wrong"

	errA                      = -1.0
	sideA                     = 3.0
	sideSmallA                = sideA / 2
	sideB                     = 4.0
	sideC                     = 5.0
	expectSquare              = 8.0
	StraightIsoscelesTriangle = true
	errNegative               = v1.NegativeTriangle
	errWrong                  = v1.WrongTriangle
)

var (
	sideStraightIsosceles = math.Sqrt(2) * 4.0
)

func TestTriangle_Square(t *testing.T) {
	t.Run(testSquare, func(t *testing.T) {
		triangle := NewTriangle(sideB, sideB, sideStraightIsosceles)
		square := triangle.Square()
		assert.Equal(t, expectSquare, square)
	})
}

func TestTriangle_CheckStraightIsoscelesTriangle(t *testing.T) {
	t.Run(testCheckStraightIsoscelesTriangle, func(t *testing.T) {
		triangle := NewTriangle(sideB, sideB, sideStraightIsosceles)
		b := triangle.CheckStraightIsoscelesTriangle()
		assert.EqualValues(t, StraightIsoscelesTriangle, b)
	})
}

func TestSquare_CheckCorrectSuccess(t *testing.T) {
	t.Run(testTriangleCheckCorrectSuccess, func(t *testing.T) {
		triangle := NewTriangle(sideA, sideB, sideC)
		err := triangle.CheckCorrect()
		assert.NoError(t, err)
	})
}

func TestTriangle_CheckCorrectNegative(t *testing.T) {
	t.Run(testTriangleCheckCorrectNegative, func(t *testing.T) {
		triangle := NewTriangle(errA, sideB, sideC)
		err := triangle.CheckCorrect()
		assert.EqualError(t, err, errNegative, err)
	})
}

func TestTriangle_CheckCorrectWrong(t *testing.T) {
	t.Run(testTriangleCheckCorrectWrong, func(t *testing.T) {
		triangle := NewTriangle(sideSmallA, sideSmallA, sideA)
		err := triangle.CheckCorrect()
		assert.EqualError(t, err, errWrong, err)
	})
}
