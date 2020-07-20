package builders

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/api/v1"
)

// BuilderCostGetter take parameters from builders
type BuilderCostGetter interface {
	BuilderCostGet() (cost float64, err error)
}

type builders struct {
	cash float64
}

// BuilderCostGet get builders cost or return error
func (b *builders) BuilderCostGet() (cost float64, err error) {
	switch {
	case b.cash >= v1.MasonsCost:
		{
			cost = v1.MasonsCost
		}
	case b.cash >= v1.HandymenCost:
		{
			cost = v1.HandymenCost
		}
	default:
		{
			err = errors.New(v1.NoMoneyForBuilders)
		}
	}
	return
}

// NewBuilderCostGet choose new builders depending on the amount of cash
func NewBuilderCostGet(amount float64) BuilderCostGetter {
	return &builders{
		cash: amount,
	}
}
