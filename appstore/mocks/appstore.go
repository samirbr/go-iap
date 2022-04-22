// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/awa/go-iap/appstore (interfaces: IAPClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	appstore "github.com/samirbr/go-iap/appstore"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAPClient is a mock of IAPClient interface
type MockIAPClient struct {
	ctrl     *gomock.Controller
	recorder *MockIAPClientMockRecorder
}

// MockIAPClientMockRecorder is the mock recorder for MockIAPClient
type MockIAPClientMockRecorder struct {
	mock *MockIAPClient
}

// NewMockIAPClient creates a new mock instance
func NewMockIAPClient(ctrl *gomock.Controller) *MockIAPClient {
	mock := &MockIAPClient{ctrl: ctrl}
	mock.recorder = &MockIAPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAPClient) EXPECT() *MockIAPClientMockRecorder {
	return m.recorder
}

// Verify mocks base method
func (m *MockIAPClient) Verify(arg0 context.Context, arg1 appstore.IAPRequest, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockIAPClientMockRecorder) Verify(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockIAPClient)(nil).Verify), arg0, arg1, arg2)
}
