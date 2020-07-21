package rambler

import (
	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
	"github.com/konvulcius/patternsWB/visitor/pkg/visitor"
)

type Rambler interface {
	ClickToRambler() (s string)
	Accept(v visitor.Visitor) (s string)
}

type rambler struct {
}

func (r *rambler) ClickToRambler() (s string) {
	s = v1.RamblerBar
	return
}

func (r *rambler) Accept(v visitor.Visitor) (s string) {
	return v.VisitRambler(r)
}

func NewRambler() Rambler {
	return &rambler{}
}
