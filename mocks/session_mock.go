// Code generated by MockGen. DO NOT EDIT.
// Source: ./utils/session/session.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/patrickchagastavares/conta-corrent/model"
)

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockSession) Generate(ctx context.Context, account *model.Account) (string, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, account)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Generate indicates an expected call of Generate.
func (mr *MockSessionMockRecorder) Generate(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockSession)(nil).Generate), ctx, account)
}

// LoadSession mocks base method.
func (m *MockSession) LoadSession(ctx context.Context, tokenString string) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSession", ctx, tokenString)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadSession indicates an expected call of LoadSession.
func (mr *MockSessionMockRecorder) LoadSession(ctx, tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSession", reflect.TypeOf((*MockSession)(nil).LoadSession), ctx, tokenString)
}