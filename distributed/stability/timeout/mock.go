// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package timeout

import mock "github.com/stretchr/testify/mock"

// MockActivator is an autogenerated mock type for the Activator type
type MockActivator struct {
	mock.Mock
}

// Activate provides a mock function with given fields: _a0
func (_m *MockActivator) Activate(_a0 int) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}