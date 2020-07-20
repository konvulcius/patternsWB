package bricks

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/api/v1"
)

// BrickCostGetter parameters from bricks
type BrickCostGetter interface {
	BrickCostGet() (cost float64, err error)
}

type bricks struct {
	cash float64
}

// BrickCostGet get bricks or return error
func (b *bricks) BrickCostGet() (cost float64, err error) {
	switch {
	case b.cash >= v1.CeramicCost:
		{
			cost = v1.CeramicCost
		}
	case b.cash >= v1.SilicateCost:
		{
			cost = v1.SilicateCost
		}
	case b.cash >= v1.GasBlockCost:
		{
			cost = v1.GasBlockCost
		}
	default:
		{
			err = errors.New(v1.NoMoneyForBricks)
		}
	}
	return
}

// NewBrickCostGetter choose bricks depending on the amount of cash
func NewBrickCostGetter(amount float64) BrickCostGetter {
	return &bricks{
		cash: amount,
	}
}
