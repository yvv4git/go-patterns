package breaker

import (
	"github.com/stretchr/testify/mock"
	"log"
	"time"
)

// NewMockActivator - simple factory for init MockActivator entity
func NewMockActivator() *MockActivator {
	mockActivator := &MockActivator{}

	mockActivator.On("Activate", mock.Anything).
		Run(func(args mock.Arguments) {
			log.Println("wait in mock-1")
			time.Sleep(time.Second)
		}).
		Return(
			Result{}, ErrTimeoutAction,
		).Times(5)

	mockActivator.On("Activate", mock.Anything).
		Run(func(args mock.Arguments) {
			log.Println("wait in mock-2")
			time.Sleep(time.Second)
		}).
		Return(
			Result{X: 5}, nil,
		)

	return mockActivator
}
