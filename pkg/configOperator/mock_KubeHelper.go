// Code generated by mockery v1.0.0. DO NOT EDIT.
package configOperator

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/client-go/pkg/api/v1"
import watch "k8s.io/apimachinery/pkg/watch"

// MockKubeHelper is an autogenerated mock type for the KubeHelper type
type MockKubeHelper struct {
	mock.Mock
}

// createClientConfigSecret provides a mock function with given fields: clientId, serviceInstanceName, serviceInstanceId, pushAppId
func (_m *MockKubeHelper) createClientConfigSecret(clientId string, serviceInstanceName string, serviceInstanceId string, pushAppId string) *v1.Secret {
	ret := _m.Called(clientId, serviceInstanceName, serviceInstanceId, pushAppId)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(string, string, string, string) *v1.Secret); ok {
		r0 = rf(clientId, serviceInstanceName, serviceInstanceId, pushAppId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	return r0
}

// deleteSecret provides a mock function with given fields: name
func (_m *MockKubeHelper) deleteSecret(name string) {
	_m.Called(name)
}

// deleteServiceBinding provides a mock function with given fields: bindingName
func (_m *MockKubeHelper) deleteServiceBinding(bindingName string) error {
	ret := _m.Called(bindingName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bindingName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// findMobileClientConfig provides a mock function with given fields: clientId
func (_m *MockKubeHelper) findMobileClientConfig(clientId string) *v1.Secret {
	ret := _m.Called(clientId)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(string) *v1.Secret); ok {
		r0 = rf(clientId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	return r0
}

// getServiceBindingNameByID provides a mock function with given fields: bindingId
func (_m *MockKubeHelper) getServiceBindingNameByID(bindingId string) (string, error) {
	ret := _m.Called(bindingId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(bindingId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bindingId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// listSecrets provides a mock function with given fields: selector
func (_m *MockKubeHelper) listSecrets(selector string) (*v1.SecretList, error) {
	ret := _m.Called(selector)

	var r0 *v1.SecretList
	if rf, ok := ret.Get(0).(func(string) *v1.SecretList); ok {
		r0 = rf(selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.SecretList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// startSecretWatch provides a mock function with given fields:
func (_m *MockKubeHelper) startSecretWatch() (watch.Interface, error) {
	ret := _m.Called()

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func() watch.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// updateSecret provides a mock function with given fields: secret
func (_m *MockKubeHelper) updateSecret(secret *v1.Secret) (*v1.Secret, error) {
	ret := _m.Called(secret)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(*v1.Secret) *v1.Secret); ok {
		r0 = rf(secret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Secret) error); ok {
		r1 = rf(secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}