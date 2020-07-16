package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/pkg/builders/models"
)

const (
	testBuildersGetterSuccess = "Test builders getter success"
	testBuildersGetterFail    = "Test builders getter fail"
	unexpectedError           = "unexpected error:"
)

var (
	givenAmounts = []float64{
		25000.0,
		5500.0,
	}
	givenLow     = 4900.0
	expectedCost = []float64{
		models.MasonsCost,
		models.HandymenCost,
	}
	expectedErr = models.NoMoney
)

func Test_NewGetter(t *testing.T) {
	var builders []Getter
	for _, cash := range givenAmounts {
		builders = append(builders, NewGetter(cash))
	}
	t.Run(testBuildersGetterSuccess, func(t *testing.T) {
		for i, builder := range builders {
			builderCost, errBuilder := builder.Get()
			assert.NoError(t, errBuilder, unexpectedError, errBuilder)
			assert.Equal(t, expectedCost[i], builderCost)
		}
	})
	noMoneyForBuilder := NewGetter(givenLow)
	t.Run(testBuildersGetterFail, func(t *testing.T) {
		_, errBuilder := noMoneyForBuilder.Get()
		assert.EqualError(t, errBuilder, expectedErr, errBuilder)
	})
}
