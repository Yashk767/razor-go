// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// TimeInterface is an autogenerated mock type for the TimeInterface type
type TimeInterface struct {
	mock.Mock
}

// Sleep provides a mock function with given fields: duration
func (_m *TimeInterface) Sleep(duration time.Duration) {
	_m.Called(duration)
}

type mockConstructorTestingTNewTimeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimeInterface creates a new instance of TimeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimeInterface(t mockConstructorTestingTNewTimeInterface) *TimeInterface {
	mock := &TimeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
