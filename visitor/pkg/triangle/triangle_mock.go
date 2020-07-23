package triangle

import (
	"github.com/stretchr/testify/mock"
)

// MockTriangle ...
type MockTriangle struct {
	mock.Mock
}

// GetSides ...
func (t *MockTriangle) GetSides() (a, b, c float64) {
	args := t.Called()
	return args.Get(0).(float64), args.Get(1).(float64), args.Get(2).(float64)
}

// Square ...
func (t *MockTriangle) Square() (len float64) {
	args := t.Called()
	return args.Get(0).(float64)
}

// CheckStraightIsoscelesTriangle ...
func (t *MockTriangle) CheckStraightIsoscelesTriangle() (b bool) {
	args := t.Called()
	return args.Bool(0)
}

// CheckCorrect ...
func (t *MockTriangle) CheckCorrect() (err error) {
	args := t.Called()
	return args.Error(0)
}

// Accept ...
func (t *MockTriangle) Accept(v visitor) (msg string, err error) {
	args := t.Called()
	return args.String(0), args.Error(1)
}
