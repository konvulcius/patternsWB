package builders

import (
	"github.com/stretchr/testify/mock"
)

//MockBuilders ...
type MockBuilders struct {
	mock.Mock
}

//Get ...
func (b *MockBuilders) Get() (cost float64, err error) {
	args := b.Called()
	if a, ok := args.Get(0).(float64); ok {
		return a, args.Error(1)
	}
	return cost, args.Error(1)
}
