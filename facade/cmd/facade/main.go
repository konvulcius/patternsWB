package main

import (
	"fmt"
	"github.com/konvulcius/patternsWB/facade/pkg/facade"
)

func main() {
	shirt := facade.NewSpecifier()
	fmt.Println(shirt.Specify())
}