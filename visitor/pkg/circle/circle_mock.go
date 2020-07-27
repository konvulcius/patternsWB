package circle

import (
	"github.com/stretchr/testify/mock"
)

// MockCircle ...
type MockCircle struct {
	mock.Mock
}

// GetRadius ...
func (c *MockCircle) GetRadius() (len float64) {
	args := c.Called()
	return args.Get(0).(float64)
}

// Length ...
func (c *MockCircle) Length() (len float64) {
	args := c.Called()
	return args.Get(0).(float64)
}

// Square ...
func (c *MockCircle) Square() (len float64) {
	args := c.Called()
	return args.Get(0).(float64)
}

// CheckCorrect ...
func (c *MockCircle) CheckCorrect() (err error) {
	args := c.Called()
	return args.Error(0)
}

// Accept ...
func (c *MockCircle) Accept(v visitor) (msg string, err error) {
	args := c.Called()
	return args.String(0), args.Error(1)
}
