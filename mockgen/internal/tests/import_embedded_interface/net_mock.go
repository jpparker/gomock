// Code generated by MockGen. DO NOT EDIT.
// Source: net.go

// Package bugreport is a generated GoMock package.
package bugreport

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/jpparker/gomock/gomock"
)

// MockNet is a mock of Net interface.
type MockNet struct {
	ctrl     *gomock.Controller
	recorder *MockNetMockRecorder
}

// MockNetMockRecorder is the mock recorder for MockNet.
type MockNetMockRecorder struct {
	mock *MockNet
}

// NewMockNet creates a new mock instance.
func NewMockNet(ctrl *gomock.Controller) *MockNet {
	mock := &MockNet{ctrl: ctrl}
	mock.recorder = &MockNetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNet) EXPECT() *MockNetMockRecorder {
	return m.recorder
}

// Header mocks base method.
func (m *MockNet) Header() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockNetMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockNet)(nil).Header))
}

// Write mocks base method.
func (m *MockNet) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockNetMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockNet)(nil).Write), arg0)
}

// WriteHeader mocks base method.
func (m *MockNet) WriteHeader(statusCode int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", statusCode)
}

// WriteHeader indicates an expected call of WriteHeader.
func (mr *MockNetMockRecorder) WriteHeader(statusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockNet)(nil).WriteHeader), statusCode)
}
