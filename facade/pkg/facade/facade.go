package facade

import (
	"strconv"

	"github.com/konvulcius/patternsWB/facade/pkg/bricks"
	"github.com/konvulcius/patternsWB/facade/pkg/builders"
	"github.com/konvulcius/patternsWB/facade/pkg/facade/models"
)

//BrigadierWorker our facade
type BrigadierWorker interface {
	BrigadierWork() (string, error)
}

type brigadier struct {
	amount   float64
	bricks   bricks.Getter
	builders builders.Getter
	err      error
}

//BrigadierWork brigadier buy materials and hire workers
func (b *brigadier) BrigadierWork() (s string, err error) {
	b.builders = builders.NewGetter(b.amount)
	buildersCost, errBuilders := b.builders.Get()
	if errBuilders != nil {
		b.err = errBuilders
		return s, b.err
	}
	b.amount -= buildersCost
	b.bricks = bricks.NewGetter(b.amount)
	bricksCost, errBricks := b.bricks.Get()
	if errBricks != nil {
		b.err = errBricks
		return s, b.err
	}
	b.amount -= bricksCost
	s = models.Prefix + strconv.FormatFloat(bricksCost+buildersCost, 'f', 0, 64) + models.Suffix
	return s, err
}

//NewBrigadierWorker hire a brigadier
func NewBrigadierWorker(amount float64) (b BrigadierWorker) {
	return newBrigadier(amount)
}

func newBrigadier(amount float64) (b *brigadier) {
	return &brigadier{
		amount: amount,
	}
}
