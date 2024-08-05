// Code generated by MockGen. DO NOT EDIT.
// Source: emoji.go
//
// Generated by this command:
//
//	mockgen -source emoji.go -destination emoji_mock_test.go -package emojidl
//

// Package emojidl is a generated GoMock package.
package emojidl

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// Mockemojidumper is a mock of emojidumper interface.
type Mockemojidumper struct {
	ctrl     *gomock.Controller
	recorder *MockemojidumperMockRecorder
}

// MockemojidumperMockRecorder is the mock recorder for Mockemojidumper.
type MockemojidumperMockRecorder struct {
	mock *Mockemojidumper
}

// NewMockemojidumper creates a new mock instance.
func NewMockemojidumper(ctrl *gomock.Controller) *Mockemojidumper {
	mock := &Mockemojidumper{ctrl: ctrl}
	mock.recorder = &MockemojidumperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockemojidumper) EXPECT() *MockemojidumperMockRecorder {
	return m.recorder
}

// DumpEmojis mocks base method.
func (m *Mockemojidumper) DumpEmojis(ctx context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpEmojis", ctx)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpEmojis indicates an expected call of DumpEmojis.
func (mr *MockemojidumperMockRecorder) DumpEmojis(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpEmojis", reflect.TypeOf((*Mockemojidumper)(nil).DumpEmojis), ctx)
}
