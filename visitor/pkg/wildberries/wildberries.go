package wildberries

import (
	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
	"github.com/konvulcius/patternsWB/visitor/pkg/visitor"
)

type WildBerries interface {
	ClickToWildBerries() (s string)
	Accept(v visitor.Visitor) (s string)
}

type wildBerries struct {
}

func (w *wildBerries) ClickToWildBerries() (s string) {
	s = v1.WildBerriesBar
	return
}

func (w *wildBerries) Accept(v visitor.Visitor) (s string) {
	return v.VisitWildBerries(w)
}

func NewWildBerries() WildBerries {
	return &wildBerries{}
}
