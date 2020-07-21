package visitor

type rambler interface {
	ClickToRambler() (s string)
}

type google interface {
	ClickToGoogle() (s string)
}

type wildBerries interface {
	ClickToWildBerries() (s string)
}

type Visitor interface {
	VisitRambler(r rambler) (s string)
	VisitGoogle(g google) (s string)
	VisitWildBerries(w wildBerries) (s string)
}

type human struct {
}

func (h *human) VisitRambler(r rambler) string {
	return r.ClickToRambler()
}

func (h *human) VisitGoogle(g google) string {
	return g.ClickToGoogle()
}

func (h *human) VisitWildBerries(w wildBerries) string {
	return w.ClickToWildBerries()
}

func NewVisitor() Visitor {
	return &human{}
}
