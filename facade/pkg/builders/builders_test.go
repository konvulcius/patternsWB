package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/api/v1"
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
		v1.MasonsCost,
		v1.HandymenCost,
	}
	expectedErr = v1.NoMoneyForBuilders
)

func TestBuilders_BuilderCostGetSuccess(t *testing.T) {
	var builders []BuilderCostGetter
	for _, cash := range givenAmounts {
		builders = append(builders, NewBuilderCostGet(cash))
	}
	t.Run(testBuildersGetterSuccess, func(t *testing.T) {
		for i, builder := range builders {
			builderCost, errBuilder := builder.BuilderCostGet()
			assert.NoError(t, errBuilder, unexpectedError, errBuilder)
			assert.Equal(t, expectedCost[i], builderCost)
		}
	})
}

func TestBuilders_BuilderCostGetFail(t *testing.T) {
	noMoneyForBuilder := NewBuilderCostGet(givenLow)
	t.Run(testBuildersGetterFail, func(t *testing.T) {
		_, errBuilder := noMoneyForBuilder.BuilderCostGet()
		assert.EqualError(t, errBuilder, expectedErr, errBuilder)
	})
}
