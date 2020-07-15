package main

import (
	"fmt"
	"os"

	"github.com/konvulcius/patternsWB/facade/pkg/facade"
)

func main() {
	fmt.Println("Enter your amount of cash")
	var amount float64
	_, errScan := fmt.Scan(&amount)
	if errScan != nil {
		fmt.Println("you should enter a digit")
		return
	}
	brigadier := facade.NewBrigadierWorker(amount)
	spending, err := brigadier.BrigadierWork()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(spending)
}
