package timeout

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// NewMockActivator - simple factory for init MockActivator entity
func NewMockActivator() *MockActivator {
	mockActivator := &MockActivator{}

	mockActivator.On("Activate", 5).
		Run(func(args mock.Arguments) {
			time.Sleep(time.Second)
		}).
		Return(5, nil).Once()

	mockActivator.On("Activate", 5).
		Run(func(args mock.Arguments) {
			time.Sleep(time.Second)
		}).
		Return(2, nil).Twice()

	mockActivator.On("Activate", 5).
		Run(func(args mock.Arguments) {
			time.Sleep(time.Second)
		}).
		Return(1, nil)

	return mockActivator
}
