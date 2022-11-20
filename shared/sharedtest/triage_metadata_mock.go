// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: TriageMetadata)

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
)

// MockTriageMetadata is a mock of TriageMetadata interface.
type MockTriageMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockTriageMetadataMockRecorder
}

// MockTriageMetadataMockRecorder is the mock recorder for MockTriageMetadata.
type MockTriageMetadataMockRecorder struct {
	mock *MockTriageMetadata
}

// NewMockTriageMetadata creates a new mock instance.
func NewMockTriageMetadata(ctrl *gomock.Controller) *MockTriageMetadata {
	mock := &MockTriageMetadata{ctrl: ctrl}
	mock.recorder = &MockTriageMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTriageMetadata) EXPECT() *MockTriageMetadataMockRecorder {
	return m.recorder
}

// Triage mocks base method.
func (m *MockTriageMetadata) Triage(arg0 shared.MetadataResults) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Triage", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Triage indicates an expected call of Triage.
func (mr *MockTriageMetadataMockRecorder) Triage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Triage", reflect.TypeOf((*MockTriageMetadata)(nil).Triage), arg0)
}
