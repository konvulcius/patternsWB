package visitor

type site interface {
	Accept(v Visitor) (s string)
}

type Sites interface {
	Add(s site)
	Accept(v Visitor) (s string)
}

type internet struct {
	sites []site
}

func (i *internet) Add(s site) {
	i.sites = append(i.sites, s)
}

func (i *internet) Accept(v Visitor) (s string) {
	var result string
	for _, s := range i.sites {
		result += s.Accept(v)
	}
	return result
}

func NewSites() Sites {
	return &internet{}
}
