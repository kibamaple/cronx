// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kibamaple/cronx/tool (interfaces: IElection,ITimer)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIElection is a mock of IElection interface
type MockIElection struct {
	ctrl     *gomock.Controller
	recorder *MockIElectionMockRecorder
}

// MockIElectionMockRecorder is the mock recorder for MockIElection
type MockIElectionMockRecorder struct {
	mock *MockIElection
}

// NewMockIElection creates a new mock instance
func NewMockIElection(ctrl *gomock.Controller) *MockIElection {
	mock := &MockIElection{ctrl: ctrl}
	mock.recorder = &MockIElectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIElection) EXPECT() *MockIElectionMockRecorder {
	return m.recorder
}

// Campaign mocks base method
func (m *MockIElection) Campaign() error {
	ret := m.ctrl.Call(m, "Campaign")
	ret0, _ := ret[0].(error)
	return ret0
}

// Campaign indicates an expected call of Campaign
func (mr *MockIElectionMockRecorder) Campaign() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Campaign", reflect.TypeOf((*MockIElection)(nil).Campaign))
}

// Destroy mocks base method
func (m *MockIElection) Destroy() error {
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockIElectionMockRecorder) Destroy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockIElection)(nil).Destroy))
}

// Resign mocks base method
func (m *MockIElection) Resign() error {
	ret := m.ctrl.Call(m, "Resign")
	ret0, _ := ret[0].(error)
	return ret0
}

// Resign indicates an expected call of Resign
func (mr *MockIElectionMockRecorder) Resign() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resign", reflect.TypeOf((*MockIElection)(nil).Resign))
}

// MockITimer is a mock of ITimer interface
type MockITimer struct {
	ctrl     *gomock.Controller
	recorder *MockITimerMockRecorder
}

// MockITimerMockRecorder is the mock recorder for MockITimer
type MockITimerMockRecorder struct {
	mock *MockITimer
}

// NewMockITimer creates a new mock instance
func NewMockITimer(ctrl *gomock.Controller) *MockITimer {
	mock := &MockITimer{ctrl: ctrl}
	mock.recorder = &MockITimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITimer) EXPECT() *MockITimerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockITimer) Add(arg0, arg1 string, arg2 func(int64)) error {
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockITimerMockRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockITimer)(nil).Add), arg0, arg1, arg2)
}

// Remove mocks base method
func (m *MockITimer) Remove(arg0 string) {
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove
func (mr *MockITimerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockITimer)(nil).Remove), arg0)
}

// Start mocks base method
func (m *MockITimer) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockITimerMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockITimer)(nil).Start))
}

// Stop mocks base method
func (m *MockITimer) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockITimerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockITimer)(nil).Stop))
}
