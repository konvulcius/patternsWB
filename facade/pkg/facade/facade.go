package facade

import (
	"strconv"

	"github.com/konvulcius/patternsWB/facade/pkg/api/v1"
)

type brickGetter interface {
	BrickCostGet() (cost float64, err error)
}

type builderGetter interface {
	BuilderCostGet() (cost float64, err error)
}

// BrigadierWorker our facade
type BrigadierWorker interface {
	BrigadierWork() (s string, err error)
}

type brigadier struct {
	bricks   brickGetter
	builders builderGetter
}

// BrigadierWork brigadier buy materials and hire workers
func (b *brigadier) BrigadierWork() (s string, err error) {
	var buildersCost, bricksCost float64
	buildersCost, err = b.builders.BuilderCostGet()
	if err != nil {
		return
	}
	bricksCost, err = b.bricks.BrickCostGet()
	if err != nil {
		return
	}
	s = v1.Prefix + strconv.FormatFloat(bricksCost+buildersCost, 'f', 0, 64) + v1.Suffix
	return
}

// NewBrigadierWorker hire a brigadier
func NewBrigadierWorker(bricks brickGetter, builders builderGetter) BrigadierWorker {
	return &brigadier{
		bricks:   bricks,
		builders: builders,
	}
}
