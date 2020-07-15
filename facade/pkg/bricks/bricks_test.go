package bricks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/pkg/bricks/models"
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
	expectedErr = models.NoMoney
)

func Test_NewGetter(t *testing.T) {
	var bricks []Getter
	for _, cash := range givenAmounts {
		bricks = append(bricks, NewGetter(cash))
	}
	t.Run("Test bricks getter success", func(t *testing.T) {
		for i, brick := range bricks {
			brickCost, errBrick := brick.Get()
			assert.NoError(t, errBrick, "unexpected error:", errBrick)
			assert.Equal(t, expectedCost[i], brickCost)
		}
	})
	noMoneyForBricks := NewGetter(givenLow)
	t.Run("Test bricks getter fail", func(t *testing.T) {
		_, errBrick := noMoneyForBricks.Get()
		assert.EqualError(t, errBrick, models.NoMoney, errBrick)
	})
}
