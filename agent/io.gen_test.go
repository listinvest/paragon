// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: Writer,WriteCloser)

// Package agent_test is a generated GoMock package.
package agent_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWriter is a mock of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
}

// MockWriteCloser is a mock of WriteCloser interface
type MockWriteCloser struct {
	ctrl     *gomock.Controller
	recorder *MockWriteCloserMockRecorder
}

// MockWriteCloserMockRecorder is the mock recorder for MockWriteCloser
type MockWriteCloserMockRecorder struct {
	mock *MockWriteCloser
}

// NewMockWriteCloser creates a new mock instance
func NewMockWriteCloser(ctrl *gomock.Controller) *MockWriteCloser {
	mock := &MockWriteCloser{ctrl: ctrl}
	mock.recorder = &MockWriteCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriteCloser) EXPECT() *MockWriteCloserMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockWriteCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWriteCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWriteCloser)(nil).Close))
}

// Write mocks base method
func (m *MockWriteCloser) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockWriteCloserMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriteCloser)(nil).Write), arg0)
}
