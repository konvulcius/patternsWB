package main

import (
	"fmt"
	"os"

	"github.com/konvulcius/patternsWB/facade/pkg/bricks"
	"github.com/konvulcius/patternsWB/facade/pkg/builders"
	"github.com/konvulcius/patternsWB/facade/pkg/facade"
)

const (
	cashForMaterials = 2000
	cashForWorkers   = 12200
)

func main() {
	bricks := bricks.NewGetter(cashForMaterials)
	builders := builders.NewGetter(cashForWorkers)
	brigadier := facade.NewBrigadierWorker(bricks, builders)
	spending, err := brigadier.BrigadierWork()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(spending)
}
