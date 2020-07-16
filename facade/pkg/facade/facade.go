package facade

import (
	"strconv"

	"github.com/konvulcius/patternsWB/facade/pkg/facade/models"
)

type getter interface {
	Get() (cost float64, err error)
}

//BrigadierWorker our facade
type BrigadierWorker interface {
	BrigadierWork() (string, error)
}

type brigadier struct {
	bricks   getter
	builders getter
}

//BrigadierWork brigadier buy materials and hire workers
func (b *brigadier) BrigadierWork() (s string, err error) {
	buildersCost, errBuilders := b.builders.Get()
	if errBuilders != nil {
		return s, errBuilders
	}
	bricksCost, errBricks := b.bricks.Get()
	if errBricks != nil {
		return s, errBricks
	}
	s = models.Prefix + strconv.FormatFloat(bricksCost+buildersCost, 'f', 0, 64) + models.Suffix
	return s, err
}

//NewBrigadierWorker hire a brigadier
func NewBrigadierWorker(bricks getter, builders getter) (b BrigadierWorker) {
	return &brigadier{
		bricks:   bricks,
		builders: builders,
	}
}
