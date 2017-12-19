// Code generated by MockGen. DO NOT EDIT.
// Source: /vagrant/golang/src/jobkeeper/core/common.go

// Package core is a generated GoMock package.
package core

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIExecutor is a mock of IExecutor interface
type MockIExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockIExecutorMockRecorder
}

// MockIExecutorMockRecorder is the mock recorder for MockIExecutor
type MockIExecutorMockRecorder struct {
	mock *MockIExecutor
}

// NewMockIExecutor creates a new mock instance
func NewMockIExecutor(ctrl *gomock.Controller) *MockIExecutor {
	mock := &MockIExecutor{ctrl: ctrl}
	mock.recorder = &MockIExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIExecutor) EXPECT() *MockIExecutorMockRecorder {
	return m.recorder
}

// Listen mocks base method
func (m *MockIExecutor) Listen(arg0 <-chan *CTask, arg1 chan<- *CResult) {
	m.ctrl.Call(m, "Listen", arg0, arg1)
}

// Listen indicates an expected call of Listen
func (mr *MockIExecutorMockRecorder) Listen(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockIExecutor)(nil).Listen), arg0, arg1)
}

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
func (m *MockIElection) Campaign(arg0 string) error {
	ret := m.ctrl.Call(m, "Campaign", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Campaign indicates an expected call of Campaign
func (mr *MockIElectionMockRecorder) Campaign(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Campaign", reflect.TypeOf((*MockIElection)(nil).Campaign), arg0)
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

// MockIStore is a mock of IStore interface
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
}

// MockIStoreMockRecorder is the mock recorder for MockIStore
type MockIStoreMockRecorder struct {
	mock *MockIStore
}

// NewMockIStore creates a new mock instance
func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

// Put mocks base method
func (m *MockIStore) Put(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockIStoreMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIStore)(nil).Put), arg0, arg1)
}

// Get mocks base method
func (m *MockIStore) Get(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIStore)(nil).Get), arg0)
}

// GetByPrefix mocks base method
func (m *MockIStore) GetByPrefix(arg0 string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "GetByPrefix", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPrefix indicates an expected call of GetByPrefix
func (mr *MockIStoreMockRecorder) GetByPrefix(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPrefix", reflect.TypeOf((*MockIStore)(nil).GetByPrefix), arg0)
}

// Delete mocks base method
func (m *MockIStore) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIStoreMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIStore)(nil).Delete), arg0)
}

// DeleteByPrefix mocks base method
func (m *MockIStore) DeleteByPrefix(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteByPrefix", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByPrefix indicates an expected call of DeleteByPrefix
func (mr *MockIStoreMockRecorder) DeleteByPrefix(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByPrefix", reflect.TypeOf((*MockIStore)(nil).DeleteByPrefix), arg0)
}

// Watch mocks base method
func (m *MockIStore) Watch(arg0 string) ([]CEvent, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].([]CEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockIStoreMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockIStore)(nil).Watch), arg0)
}

// WatchByPrefix mocks base method
func (m *MockIStore) WatchByPrefix(arg0 string) ([]CEvent, error) {
	ret := m.ctrl.Call(m, "WatchByPrefix", arg0)
	ret0, _ := ret[0].([]CEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchByPrefix indicates an expected call of WatchByPrefix
func (mr *MockIStoreMockRecorder) WatchByPrefix(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchByPrefix", reflect.TypeOf((*MockIStore)(nil).WatchByPrefix), arg0)
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

// MockIScheduler is a mock of IScheduler interface
type MockIScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockISchedulerMockRecorder
}

// MockISchedulerMockRecorder is the mock recorder for MockIScheduler
type MockISchedulerMockRecorder struct {
	mock *MockIScheduler
}

// NewMockIScheduler creates a new mock instance
func NewMockIScheduler(ctrl *gomock.Controller) *MockIScheduler {
	mock := &MockIScheduler{ctrl: ctrl}
	mock.recorder = &MockISchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIScheduler) EXPECT() *MockISchedulerMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockIScheduler) Dispatch(arg0 map[string]string, arg1 <-chan *CResult, arg2 chan<- *CTask) {
	m.ctrl.Call(m, "Dispatch", arg0, arg1, arg2)
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockISchedulerMockRecorder) Dispatch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockIScheduler)(nil).Dispatch), arg0, arg1, arg2)
}