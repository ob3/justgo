// Code generated by MockGen. DO NOT EDIT.
// Source: app_interface.go

// Package justgo is a generated GoMock package.
package justgo

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppInterface is a mock of AppInterface interface
type MockAppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAppInterfaceMockRecorder
}

// MockAppInterfaceMockRecorder is the mock recorder for MockAppInterface
type MockAppInterfaceMockRecorder struct {
	mock *MockAppInterface
}

// NewMockAppInterface creates a new mock instance
func NewMockAppInterface(ctrl *gomock.Controller) *MockAppInterface {
	mock := &MockAppInterface{ctrl: ctrl}
	mock.recorder = &MockAppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppInterface) EXPECT() *MockAppInterfaceMockRecorder {
	return m.recorder
}

// Serve mocks base method
func (m *MockAppInterface) Serve() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Serve")
}

// Serve indicates an expected call of Serve
func (mr *MockAppInterfaceMockRecorder) Serve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockAppInterface)(nil).Serve))
}

// ShutDown mocks base method
func (m *MockAppInterface) ShutDown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShutDown")
}

// ShutDown indicates an expected call of ShutDown
func (mr *MockAppInterfaceMockRecorder) ShutDown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutDown", reflect.TypeOf((*MockAppInterface)(nil).ShutDown))
}
