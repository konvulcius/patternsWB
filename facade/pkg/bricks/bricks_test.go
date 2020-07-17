package bricks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/api/v1/models"
)

const (
	testBricksGetterSuccess = "Test bricks getter success"
	testBricksGetterFail    = "Test bricks getter fail"
	unexpectedError         = "unexpected error:"
)

var (
	givenAmounts = []float64{
		2000.0,
		1700.0,
		1200.0,
	}
	givenLow     = 100.0
	expectedCost = []float64{
		models.CeramicCost,
		models.SilicateCost,
		models.GasBlockCost,
	}
	expectedErr = models.NoMoneyForBricks
)

func Test_NewGetter(t *testing.T) {
	var bricks []Getter
	for _, cash := range givenAmounts {
		bricks = append(bricks, NewGetter(cash))
	}
	t.Run(testBricksGetterSuccess, func(t *testing.T) {
		for i, brick := range bricks {
			brickCost, errBrick := brick.Get()
			assert.NoError(t, errBrick, unexpectedError, errBrick)
			assert.Equal(t, expectedCost[i], brickCost)
		}
	})
	noMoneyForBricks := NewGetter(givenLow)
	t.Run(testBricksGetterFail, func(t *testing.T) {
		_, errBrick := noMoneyForBricks.Get()
		assert.EqualError(t, errBrick, expectedErr, errBrick)
	})
}
