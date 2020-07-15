package facade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	masonsCost        = 10000.0
	handymenCost      = 5000.0
	ceramicCost       = 1800.0
	silicateCost      = 1500.0
	gasBlockCost      = 1100.0
	noMoneyToBricks   = 5100.0
	noMoneyToBuilders = 4900.0
	spendingMoney     = "we spent %v to build a wall"
	noMoneyBricks     = "no money for bricks"
	noMoneyBuilders   = "not enough money for builders"
)

var (
	givenAmounts = []float64{
		masonsCost + ceramicCost,
		masonsCost + silicateCost,
		masonsCost + gasBlockCost,
		handymenCost + ceramicCost,
		handymenCost + silicateCost,
		handymenCost + gasBlockCost,
	}
	givenLow = []float64{
		noMoneyToBricks,
		noMoneyToBuilders,
	}
	expectedCost = []string{
		fmt.Sprintf(spendingMoney, givenAmounts[0]),
		fmt.Sprintf(spendingMoney, givenAmounts[1]),
		fmt.Sprintf(spendingMoney, givenAmounts[2]),
		fmt.Sprintf(spendingMoney, givenAmounts[3]),
		fmt.Sprintf(spendingMoney, givenAmounts[4]),
		fmt.Sprintf(spendingMoney, givenAmounts[5]),
	}
	expectedErr = []string{
		noMoneyBricks,
		noMoneyBuilders,
	}
)

func Test_BrigadierWork(t *testing.T) {
	var brigadiers []BrigadierWorker
	for _, cash := range givenAmounts {
		brigadiers = append(brigadiers, NewBrigadierWorker(cash))
	}
	t.Run("brigadier spent all money SUCCESS", func(t *testing.T) {
		for i, brigadier := range brigadiers {
			lastCry, errCry := brigadier.BrigadierWork()
			assert.NoError(t, errCry, "unexpected error:", errCry)
			assert.Equal(t, expectedCost[i], lastCry)
		}
	})
	var notEnoughMoney []BrigadierWorker
	for _, cash := range givenLow {
		notEnoughMoney = append(notEnoughMoney, NewBrigadierWorker(cash))
	}
	t.Run("brigadier spent all money FAIL", func(t *testing.T) {
		for i, brigadier := range notEnoughMoney {
			_, errCry := brigadier.BrigadierWork()
			assert.EqualError(t, errCry, expectedErr[i], errCry)
		}
	})
}
