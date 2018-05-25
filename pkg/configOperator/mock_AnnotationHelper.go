// Code generated by mockery v1.0.0. DO NOT EDIT.
package configOperator

import mock "github.com/stretchr/testify/mock"

// MockAnnotationHelper is an autogenerated mock type for the AnnotationHelper type
type MockAnnotationHelper struct {
	mock.Mock
}

// addAnnotationToMobileClient provides a mock function with given fields: clientId, appType, variantUrl, serviceInstanceName
func (_m *MockAnnotationHelper) addAnnotationToMobileClient(clientId string, appType string, variantUrl string, serviceInstanceName string) {
	_m.Called(clientId, appType, variantUrl, serviceInstanceName)
}

// removeAnnotationFromMobileClient provides a mock function with given fields: clientId, appType, serviceInstanceName
func (_m *MockAnnotationHelper) removeAnnotationFromMobileClient(clientId string, appType string, serviceInstanceName string) {
	_m.Called(clientId, appType, serviceInstanceName)
}
