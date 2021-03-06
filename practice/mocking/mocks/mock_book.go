// Code generated by MockGen. DO NOT EDIT.
// Source: 1d1go/practice/mocking/book (interfaces: Book)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBook is a mock of Book interface
type MockBook struct {
	ctrl     *gomock.Controller
	recorder *MockBookMockRecorder
}

// MockBookMockRecorder is the mock recorder for MockBook
type MockBookMockRecorder struct {
	mock *MockBook
}

// NewMockBook creates a new mock instance
func NewMockBook(ctrl *gomock.Controller) *MockBook {
	mock := &MockBook{ctrl: ctrl}
	mock.recorder = &MockBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBook) EXPECT() *MockBookMockRecorder {
	return m.recorder
}

// Page mocks base method
func (m *MockBook) Page(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Page", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Page indicates an expected call of Page
func (mr *MockBookMockRecorder) Page(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Page", reflect.TypeOf((*MockBook)(nil).Page), arg0)
}

// Read mocks base method
func (m *MockBook) Read(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockBookMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockBook)(nil).Read), arg0)
}
