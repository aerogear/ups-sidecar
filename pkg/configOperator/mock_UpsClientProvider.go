// Code generated by mockery v1.0.0. DO NOT EDIT.

package configOperator

import mock "github.com/stretchr/testify/mock"

// MockUpsClientProvider is an autogenerated mock type for the UpsClientProvider type
type MockUpsClientProvider struct {
	mock.Mock
}

// getPushClient provides a mock function with given fields:
func (_m *MockUpsClientProvider) getPushClient() UpsClient {
	ret := _m.Called()

	var r0 UpsClient
	if rf, ok := ret.Get(0).(func() UpsClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(UpsClient)
		}
	}

	return r0
}
