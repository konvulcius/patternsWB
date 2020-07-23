package circle

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
)

const (
	testCircleLength              = "Test circle length"
	testCircleSquare              = "Test circle square"
	testCircleCheckCorrectSuccess = "Test circle CheckCorrect success"
	testCircleCheckCorrectFail    = "Test circle CheckCorrect fail"

	circleRadiusRight = 10.0
	circleSquare      = math.Pi * circleRadiusRight * circleRadiusRight
	circleLength      = math.Pi * circleRadiusRight * 2
)

var (
	errInput = []float64{
		0.0,
		-10.0,
	}
	errCircle = []string{
		v1.ZeroRadius,
		v1.NegativeRadius,
	}
)

func TestCircle_Length(t *testing.T) {
	t.Run(testCircleLength, func(t *testing.T) {
		circle := NewCircle(circleRadiusRight)
		length := circle.Length()
		assert.Equal(t, circleLength, length)
	})
}

func TestCircle_Square(t *testing.T) {
	t.Run(testCircleSquare, func(t *testing.T) {
		circle := NewCircle(circleRadiusRight)
		square := circle.Square()
		assert.Equal(t, circleSquare, square)
	})
}

func TestCircle_CheckCorrectSuccess(t *testing.T) {
	t.Run(testCircleCheckCorrectSuccess, func(t *testing.T) {
		circle := NewCircle(circleRadiusRight)
		err := circle.CheckCorrect()
		assert.NoError(t, err)
	})
}

func TestCircle_CheckCorrectFail(t *testing.T) {
	t.Run(testCircleCheckCorrectFail, func(t *testing.T) {
		for i, radius := range errInput {
			circle := NewCircle(radius)
			err := circle.CheckCorrect()
			assert.EqualError(t, err, errCircle[i], err)
		}
	})
}
