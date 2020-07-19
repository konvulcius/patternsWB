package bricks

import (
	"github.com/stretchr/testify/mock"
)

//MockBricks ...
type MockBricks struct {
	mock.Mock
}

//BrickCostGet ...
func (b *MockBricks) BrickCostGet() (cost float64, err error) {
	args := b.Called()
	return args.Get(0).(float64), args.Error(1)
}
