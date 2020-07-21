package main

import (
	"fmt"

	"github.com/konvulcius/patternsWB/visitor/pkg/google"
	"github.com/konvulcius/patternsWB/visitor/pkg/rambler"
	"github.com/konvulcius/patternsWB/visitor/pkg/visitor"
	"github.com/konvulcius/patternsWB/visitor/pkg/wildberries"
)

func main() {
	sites := visitor.NewSites()

	sites.Add(rambler.NewRambler())
	sites.Add(google.NewGoogle())
	sites.Add(wildberries.NewWildBerries())

	fmt.Println(sites.Accept(visitor.NewVisitor()))
}
