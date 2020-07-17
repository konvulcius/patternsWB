package facade

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/api/v1/models"
	"github.com/konvulcius/patternsWB/facade/pkg/bricks"
	"github.com/konvulcius/patternsWB/facade/pkg/builders"
)

const (
	masonsCost                    = 10000.0
	silicateCost                  = 1500.0
	spendingMoney                 = "we spent %v to build a wall"
	brigadierSpentAllMoneySuccess = "brigadier spent all money success"
	brigadierSpentAllMoneyFail    = "brigadier spent all money fail"
	methodGet                     = "Get"
	unexpectedError               = "unexpected error:"
)

var (
	errBricks   = errors.New(models.NoMoneyForBricks)
	errBuilders = errors.New(models.NoMoneyForBuilders)
)

func TestBrigadier_BrigadierWork(t *testing.T) {
	bricks := new(bricks.MockBricks)
	bricks.On(methodGet).Return(silicateCost, nil).Once()
	builders := new(builders.MockBuilders)
	builders.On(methodGet).Return(masonsCost, nil).Once()
	brigadier := NewBrigadierWorker(bricks, builders)
	t.Run(brigadierSpentAllMoneySuccess, func(t *testing.T) {
		lastCry, errCry := brigadier.BrigadierWork()
		assert.NoError(t, errCry, unexpectedError, errCry)
		assert.Equal(t, fmt.Sprintf(spendingMoney, masonsCost+silicateCost), lastCry)
	})
	bricks.On(methodGet).Return("", errBricks).Once()
	builders.On(methodGet).Return("", errBuilders).Once()
	t.Run(brigadierSpentAllMoneyFail, func(t *testing.T) {
		_, errCry := brigadier.BrigadierWork()
		assert.EqualError(t, errCry, models.NoMoneyForBuilders, errCry)
	})
}
