package builders

import (
	"github.com/stretchr/testify/mock"
)

//MockBuilders ...
type MockBuilders struct {
	mock.Mock
}

//BuilderCostGet ...
func (b *MockBuilders) BuilderCostGet() (cost float64, err error) {
	args := b.Called()
	return args.Get(0).(float64), args.Error(1)
}
