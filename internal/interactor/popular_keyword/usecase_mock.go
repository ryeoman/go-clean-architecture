// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase.go

// Package popularkeyword is a generated GoMock package.
package popularkeyword

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecaseItf is a mock of UsecaseItf interface.
type MockUsecaseItf struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseItfMockRecorder
}

// MockUsecaseItfMockRecorder is the mock recorder for MockUsecaseItf.
type MockUsecaseItfMockRecorder struct {
	mock *MockUsecaseItf
}

// NewMockUsecaseItf creates a new mock instance.
func NewMockUsecaseItf(ctrl *gomock.Controller) *MockUsecaseItf {
	mock := &MockUsecaseItf{ctrl: ctrl}
	mock.recorder = &MockUsecaseItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseItf) EXPECT() *MockUsecaseItfMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUsecaseItf) Get(keyword string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", keyword)
}

// Get indicates an expected call of Get.
func (mr *MockUsecaseItfMockRecorder) Get(keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsecaseItf)(nil).Get), keyword)
}
