// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	http "net/http"

	device "github.com/rna-vt/devicecommander/src/device"

	mock "github.com/stretchr/testify/mock"

	model "github.com/rna-vt/devicecommander/src/graph/model"
)

// IDeviceClient is an autogenerated mock type for the IDeviceClient type
type IDeviceClient struct {
	mock.Mock
}

// EvaluateHealthCheckResponse provides a mock function with given fields: resp, d
func (_m *IDeviceClient) EvaluateHealthCheckResponse(resp *http.Response, d device.Device) bool {
	ret := _m.Called(resp, d)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*http.Response, device.Device) bool); ok {
		r0 = rf(resp, d)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Health provides a mock function with given fields: _a0
func (_m *IDeviceClient) Health(_a0 device.Device) (*http.Response, error) {
	ret := _m.Called(_a0)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(device.Device) *http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(device.Device) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Info provides a mock function with given fields: _a0
func (_m *IDeviceClient) Info(_a0 device.Device) (model.NewDevice, error) {
	ret := _m.Called(_a0)

	var r0 model.NewDevice
	if rf, ok := ret.Get(0).(func(device.Device) model.NewDevice); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.NewDevice)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(device.Device) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}