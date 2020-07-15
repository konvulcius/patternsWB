package bricks

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/bricks/models"
)

//Getter parameters from bricks
type Getter interface {
	Get() (float64, error)
}

type bricks struct {
	brickType models.BrickType
	cost      models.Cost
}

//Get get bricks or return error
func (b *bricks) Get() (cost float64, err error) {
	if b.brickType != "" {
		return float64(b.cost), err
	}
	return cost, errors.New(models.NoMoney)
}

func (b *bricks) choose(amount models.Cost) {
	switch {
	case amount >= models.CeramicCost:
		{
			b.brickType = models.Ceramic
			b.cost = models.CeramicCost
		}
	case amount >= models.SilicateCost:
		{
			b.brickType = models.Silicate
			b.cost = models.SilicateCost
		}
	case amount >= models.GasBlockCost:
		{
			b.brickType = models.GasBlock
			b.cost = models.GasBlockCost
		}
	}
}

//NewGetter choose bricks depending on the amount of cash
func NewGetter(amount float64) (g Getter) {
	bricks := newBricks()
	bricks.choose(models.Cost(amount))
	return bricks
}

func newBricks() (b *bricks) {
	return &bricks{}
}
