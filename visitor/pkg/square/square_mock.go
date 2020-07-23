package square

import (
	"github.com/stretchr/testify/mock"
)

// MockSquare ...
type MockSquare struct {
	mock.Mock
}

// GetSide ...
func (s *MockSquare) GetSide() (len float64) {
	args := s.Called()
	return args.Get(0).(float64)
}

// Diagonal ...
func (s *MockSquare) Diagonal() (len float64) {
	args := s.Called()
	return args.Get(0).(float64)
}

// CircumscribedRadius ...
func (s *MockSquare) CircumscribedRadius() (len float64) {
	args := s.Called()
	return args.Get(0).(float64)
}

// CheckCorrect ...
func (s *MockSquare) CheckCorrect() (err error) {
	args := s.Called()
	return args.Error(0)
}

// Accept ...
func (s *MockSquare) Accept(v visitor) (msg string, err error) {
	args := s.Called()
	return args.String(0), args.Error(1)
}
