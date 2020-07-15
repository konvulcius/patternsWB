package builders

import (
	"errors"

	"github.com/konvulcius/patternsWB/facade/pkg/builders/models"
)

//Getter take parameters from builders
type Getter interface {
	Get() (float64, error)
}

type builders struct {
	builders models.Builders
	cost     models.Cost
}

//Get get builders cost or return error
func (b *builders) Get() (cost float64, err error) {
	if b.builders != "" {
		return float64(b.cost), err
	}
	return cost, errors.New(models.NoMoney)
}

func (b *builders) choose(amount models.Cost) {
	switch {
	case amount >= models.MasonsCost:
		{
			b.builders = models.Masons
			b.cost = models.MasonsCost
		}
	case amount >= models.HandymenCost:
		{
			b.builders = models.Handymen
			b.cost = models.HandymenCost
		}
	}
}

//NewGetter choose new builders depending on the amount of cash
func NewGetter(amount float64) (g Getter) {
	builders := newBuilders()
	builders.choose(models.Cost(amount))
	return builders
}

func newBuilders() (b *builders) {
	return &builders{}
}
