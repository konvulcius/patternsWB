package bricks

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/api/v1"
)

//BrickCostGetter parameters from bricks
type BrickCostGetter interface {
	BrickCostGet() (cost float64, err error)
}

type bricks struct {
	brickType string
	cash      float64
}

//BrickCostGet get bricks or return error
func (b *bricks) BrickCostGet() (cost float64, err error) {
	switch {
	case b.cash >= v1.CeramicCost:
		{
			b.brickType = v1.Ceramic
			return v1.CeramicCost, err
		}
	case b.cash >= v1.SilicateCost:
		{
			b.brickType = v1.Silicate
			return v1.SilicateCost, err
		}
	case b.cash >= v1.GasBlockCost:
		{
			b.brickType = v1.GasBlock
			return v1.GasBlockCost, err
		}
	}
	return cost, errors.New(v1.NoMoneyForBricks)
}

//NewBrickCostGetter choose bricks depending on the amount of cash
func NewBrickCostGetter(amount float64) BrickCostGetter {
	return &bricks{
		cash: amount,
	}
}
