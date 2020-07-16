package main

import (
	"fmt"
	"os"

	"github.com/konvulcius/patternsWB/facade/pkg/bricks"
	"github.com/konvulcius/patternsWB/facade/pkg/builders"
	"github.com/konvulcius/patternsWB/facade/pkg/facade"
)

func main() {
	bricks := bricks.NewGetter(2000)
	builders := builders.NewGetter(12200)
	brigadier := facade.NewBrigadierWorker(bricks, builders)
	spending, err := brigadier.BrigadierWork()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(spending)
}
