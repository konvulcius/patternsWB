package square

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

const (
	testSquareDiagonal            = "Test square Diagonal"
	testSquareCircumscribedRadius = "Test square CircumscribedRadius"
	testSquareCheckCorrectSuccess = "Test square CheckCorrect success"
	testSquareCheckCorrectFail    = "Test square CheckCorrect fail"

	sideRight = 4.0
)

var (
	squareDiagonal      = math.Sqrt(2) * sideRight
	circumscribedRadius = squareDiagonal / 2

	errInput = []float64{
		0.0,
		-8.0,
	}
	errSquare = []string{
		v1.ZeroSideSquare,
		v1.NegativeSideSquare,
	}
)

func TestSquare_Diagonal(t *testing.T) {
	t.Run(testSquareDiagonal, func(t *testing.T) {
		square := NewSquare(sideRight)
		diag := square.Diagonal()
		assert.Equal(t, squareDiagonal, diag)
	})
}

func TestSquare_CircumscribedRadius(t *testing.T) {
	t.Run(testSquareCircumscribedRadius, func(t *testing.T) {
		square := NewSquare(sideRight)
		rad := square.CircumscribedRadius()
		assert.Equal(t, circumscribedRadius, rad)
	})
}

func TestSquare_CheckCorrectSuccess(t *testing.T) {
	t.Run(testSquareCheckCorrectSuccess, func(t *testing.T) {
		square := NewSquare(sideRight)
		err := square.CheckCorrect()
		assert.NoError(t, err)
	})
}

func TestSquare_CheckCorrectFail(t *testing.T) {
	t.Run(testSquareCheckCorrectFail, func(t *testing.T) {
		for i, side := range errInput {
			square := NewSquare(side)
			err := square.CheckCorrect()
			assert.EqualError(t, err, errSquare[i], err)
		}
	})
}
