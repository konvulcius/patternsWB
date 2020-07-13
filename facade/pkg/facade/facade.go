package facade

import (
	"math/rand"
	"strings"
)

type Specifier interface {
	Specify() string
}

type shirt struct {
	shirtType	*shirtType
	size		*size
	color		*color
}

type shirtType struct {
}

type size struct {
}

type color struct {
}

func (s *shirt) Specify() string {
	result := []string{
		"type: " + s.shirtType.takeType(),
		"size: " + s.size.takeSize(),
		"color: " + s.color.takeColor(),
	}
	return strings.Join(result, "\n")
}

func (s *shirtType) takeType() string {
	var types []string = []string{
		"slim",
		"turbo slim",
		"regular",
	}
	return types[rand.Intn(len(types))]
}

func (s *size) takeSize() string {
	var sizes []string = []string{
		"S",
		"M",
		"L",
		"XL",
		"XXL",
	}
	return sizes[rand.Intn(len(sizes))]
}

func (c *color) takeColor() string {
	var colores []string = []string{
		"red",
		"green",
		"blue",
		"white",
		"black",
	}
	return colores[rand.Intn(len(colores))]
}

func NewSpecifier() Specifier {
	shirtNew := &shirt{
		&shirtType{},
		&size{},
		&color{},
	}
	return Specifier(shirtNew)
}