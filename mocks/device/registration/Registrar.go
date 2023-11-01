// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	device "github.com/rna-vt/devicecommander/pkg/device"
	mock "github.com/stretchr/testify/mock"

	scanner "github.com/rna-vt/devicecommander/pkg/scanner"
)

// Registrar is an autogenerated mock type for the Registrar type
type Registrar struct {
	mock.Mock
}

// HandleProspects provides a mock function with given fields: _a0
func (_m *Registrar) HandleProspects(_a0 []scanner.FoundDevice) ([]device.Device, error) {
	ret := _m.Called(_a0)

	var r0 []device.Device
	var r1 error
	if rf, ok := ret.Get(0).(func([]scanner.FoundDevice) ([]device.Device, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]scanner.FoundDevice) []device.Device); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]device.Device)
		}
	}

	if rf, ok := ret.Get(1).(func([]scanner.FoundDevice) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRegistrar creates a new instance of Registrar. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegistrar(t interface {
	mock.TestingT
	Cleanup(func())
}) *Registrar {
	mock := &Registrar{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
