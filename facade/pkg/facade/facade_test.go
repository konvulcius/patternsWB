package facade

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konvulcius/patternsWB/facade/api/v1"
	"github.com/konvulcius/patternsWB/facade/pkg/bricks"
	"github.com/konvulcius/patternsWB/facade/pkg/builders"
)

const (
	masonsCost                    = 10000.0
	silicateCost                  = 1500.0
	spendingMoney                 = "you spent %v to build a wall"
	brigadierSpentAllMoneySuccess = "brigadier spent all money success"
	brigadierSpentAllMoneyFail    = "brigadier spent all money fail"
	methodBrickGet                = "BrickCostGet"
	methodBuilderGet              = "BuilderCostGet"
	unexpectedError               = "unexpected error:"
)

var (
	errBricks   = errors.New(v1.NoMoneyForBricks)
	errBuilders = errors.New(v1.NoMoneyForBuilders)
)

func TestBrigadier_BrigadierWorkSuccess(t *testing.T) {
	bricks := new(bricks.MockBricks)
	bricks.On(methodBrickGet).Return(silicateCost, nil).Once()
	builders := new(builders.MockBuilders)
	builders.On(methodBuilderGet).Return(masonsCost, nil).Once()
	brigadier := NewBrigadierWorker(bricks, builders)
	t.Run(brigadierSpentAllMoneySuccess, func(t *testing.T) {
		lastCry, errCry := brigadier.BrigadierWork()
		assert.NoError(t, errCry, unexpectedError, errCry)
		assert.Equal(t, fmt.Sprintf(spendingMoney, masonsCost+silicateCost), lastCry)
	})
}

func TestBrigadier_BrigadierWorkFail(t *testing.T) {
	bricks := new(bricks.MockBricks)
	bricks.On(methodBrickGet).Return("", errBricks).Once()
	builders := new(builders.MockBuilders)
	builders.On(methodBuilderGet).Return("", errBuilders).Once()
	brigadier := NewBrigadierWorker(bricks, builders)
	t.Run(brigadierSpentAllMoneyFail, func(t *testing.T) {
		_, errCry := brigadier.BrigadierWork()
		assert.EqualError(t, errCry, v1.NoMoneyForBuilders, errCry)
	})
}
