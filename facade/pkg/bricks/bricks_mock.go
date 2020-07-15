package bricks

import (
	"github.com/stretchr/testify/mock"
)

//MockBricks ...
type MockBricks struct {
	mock.Mock
}

//Get ...
func (b *MockBricks) Get() (cost float64, err error) {
	args := b.Called()
	if a, ok := args.Get(0).(float64); ok {
		return a, args.Error(1)
	}
	return cost, args.Error(1)
}
