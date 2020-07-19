package builders

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/api/v1"
)

//BuilderCostGetter take parameters from builders
type BuilderCostGetter interface {
	BuilderCostGet() (cost float64, err error)
}

type builders struct {
	builders string
	cash     float64
}

//BuilderCostGet get builders cost or return error
func (b *builders) BuilderCostGet() (cost float64, err error) {
	switch {
	case b.cash >= v1.MasonsCost:
		{
			b.builders = v1.Masons
			return v1.MasonsCost, err
		}
	case b.cash >= v1.HandymenCost:
		{
			b.builders = v1.Handymen
			return v1.HandymenCost, err
		}
	}
	err = errors.New(v1.NoMoneyForBuilders)
	return
}

//NewBuilderCostGet choose new builders depending on the amount of cash
func NewBuilderCostGet(amount float64) BuilderCostGetter {
	return &builders{
		cash: amount,
	}
}
