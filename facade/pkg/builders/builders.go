package builders

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/builders/models"
)

//Getter take parameters from builders
type Getter interface {
	Get() (cost float64, err error)
}

type builders struct {
	builders string
	cash     float64
}

//Get get builders cost or return error
func (b *builders) Get() (cost float64, err error) {
	switch {
	case b.cash >= models.MasonsCost:
		{
			b.builders = models.Masons
			return models.MasonsCost, err
		}
	case b.cash >= models.HandymenCost:
		{
			b.builders = models.Handymen
			return models.HandymenCost, err
		}
	}
	return cost, errors.New(models.NoMoney)
}

//NewGetter choose new builders depending on the amount of cash
func NewGetter(amount float64) (g Getter) {
	return &builders{
		cash: amount,
	}
}
