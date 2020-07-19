package bricks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/pkg/api/v1"
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
		v1.CeramicCost,
		v1.SilicateCost,
		v1.GasBlockCost,
	}
	expectedErr = v1.NoMoneyForBricks
)

func TestBricks_BrickCostGetSuccess(t *testing.T) {
	var bricks []BrickCostGetter
	for _, cash := range givenAmounts {
		bricks = append(bricks, NewBrickCostGetter(cash))
	}
	t.Run(testBricksGetterSuccess, func(t *testing.T) {
		for i, brick := range bricks {
			brickCost, errBrick := brick.BrickCostGet()
			assert.NoError(t, errBrick, unexpectedError, errBrick)
			assert.Equal(t, expectedCost[i], brickCost)
		}
	})
}

func TestBricks_BrickCostGetFail(t *testing.T) {
	noMoneyForBricks := NewBrickCostGetter(givenLow)
	t.Run(testBricksGetterFail, func(t *testing.T) {
		_, errBrick := noMoneyForBricks.BrickCostGet()
		assert.EqualError(t, errBrick, expectedErr, errBrick)
	})
}
