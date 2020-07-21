package google

import (
	"github.com/konvulcius/patternsWB/visitor/pkg/api/v1"
	"github.com/konvulcius/patternsWB/visitor/pkg/visitor"
)

type Google interface {
	ClickToGoogle() (s string)
	Accept(v visitor.Visitor) (s string)
}

type google struct {
}

func (g *google) ClickToGoogle() (s string) {
	s = v1.GoogleBar
	return
}

func (g *google) Accept(v visitor.Visitor) (s string) {
	return v.VisitGoogle(g)
}

func NewGoogle() Google {
	return &google{}
}
