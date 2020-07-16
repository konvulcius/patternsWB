package bricks

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/bricks/models"
)

//Getter parameters from bricks
type Getter interface {
	Get() (cost float64, err error)
}

type bricks struct {
	brickType string
	cash      float64
}

//Get get bricks or return error
func (b *bricks) Get() (cost float64, err error) {
	switch {
	case b.cash >= models.CeramicCost:
		{
			b.brickType = models.Ceramic
			return models.CeramicCost, err
		}
	case b.cash >= models.SilicateCost:
		{
			b.brickType = models.Silicate
			return models.SilicateCost, err
		}
	case b.cash >= models.GasBlockCost:
		{
			b.brickType = models.GasBlock
			return models.GasBlockCost, err
		}
	}
	return cost, errors.New(models.NoMoney)
}

//NewGetter choose bricks depending on the amount of cash
func NewGetter(amount float64) (g Getter) {
	return &bricks{
		cash: amount,
	}
}
