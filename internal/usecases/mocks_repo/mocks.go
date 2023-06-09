// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "siteavliable/internal/models"

	gomock "github.com/golang/mock/gomock"
)

// MockIRedisRepoStats is a mock of IRedisRepoStats interface.
type MockIRedisRepoStats struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisRepoStatsMockRecorder
}

// MockIRedisRepoStatsMockRecorder is the mock recorder for MockIRedisRepoStats.
type MockIRedisRepoStatsMockRecorder struct {
	mock *MockIRedisRepoStats
}

// NewMockIRedisRepoStats creates a new mock instance.
func NewMockIRedisRepoStats(ctrl *gomock.Controller) *MockIRedisRepoStats {
	mock := &MockIRedisRepoStats{ctrl: ctrl}
	mock.recorder = &MockIRedisRepoStatsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisRepoStats) EXPECT() *MockIRedisRepoStatsMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIRedisRepoStats) Get(arg0 context.Context, arg1 []string) ([]models.CounterStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([]models.CounterStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIRedisRepoStatsMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRedisRepoStats)(nil).Get), arg0, arg1)
}

// Save mocks base method.
func (m *MockIRedisRepoStats) Save(arg0 context.Context, arg1 []models.CounterStats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIRedisRepoStatsMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRedisRepoStats)(nil).Save), arg0, arg1)
}

// MockIRedisRepoClients is a mock of IRedisRepoClients interface.
type MockIRedisRepoClients struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisRepoClientsMockRecorder
}

// MockIRedisRepoClientsMockRecorder is the mock recorder for MockIRedisRepoClients.
type MockIRedisRepoClientsMockRecorder struct {
	mock *MockIRedisRepoClients
}

// NewMockIRedisRepoClients creates a new mock instance.
func NewMockIRedisRepoClients(ctrl *gomock.Controller) *MockIRedisRepoClients {
	mock := &MockIRedisRepoClients{ctrl: ctrl}
	mock.recorder = &MockIRedisRepoClientsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisRepoClients) EXPECT() *MockIRedisRepoClientsMockRecorder {
	return m.recorder
}

// GetByURL mocks base method.
func (m *MockIRedisRepoClients) GetByURL(ctx context.Context, siteName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByURL", ctx, siteName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByURL indicates an expected call of GetByURL.
func (mr *MockIRedisRepoClientsMockRecorder) GetByURL(ctx, siteName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByURL", reflect.TypeOf((*MockIRedisRepoClients)(nil).GetByURL), ctx, siteName)
}

// GetWithMax mocks base method.
func (m *MockIRedisRepoClients) GetWithMax(ctx context.Context) (string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithMax", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWithMax indicates an expected call of GetWithMax.
func (mr *MockIRedisRepoClientsMockRecorder) GetWithMax(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithMax", reflect.TypeOf((*MockIRedisRepoClients)(nil).GetWithMax), ctx)
}

// GetWithMin mocks base method.
func (m *MockIRedisRepoClients) GetWithMin(ctx context.Context) (string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithMin", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWithMin indicates an expected call of GetWithMin.
func (mr *MockIRedisRepoClientsMockRecorder) GetWithMin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithMin", reflect.TypeOf((*MockIRedisRepoClients)(nil).GetWithMin), ctx)
}

// MockIRedisRepoUpdate is a mock of IRedisRepoUpdate interface.
type MockIRedisRepoUpdate struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisRepoUpdateMockRecorder
}

// MockIRedisRepoUpdateMockRecorder is the mock recorder for MockIRedisRepoUpdate.
type MockIRedisRepoUpdateMockRecorder struct {
	mock *MockIRedisRepoUpdate
}

// NewMockIRedisRepoUpdate creates a new mock instance.
func NewMockIRedisRepoUpdate(ctrl *gomock.Controller) *MockIRedisRepoUpdate {
	mock := &MockIRedisRepoUpdate{ctrl: ctrl}
	mock.recorder = &MockIRedisRepoUpdateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisRepoUpdate) EXPECT() *MockIRedisRepoUpdateMockRecorder {
	return m.recorder
}

// SetByURL mocks base method.
func (m *MockIRedisRepoUpdate) SetByURL(ctx context.Context, siteName string, accessTime int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetByURL", ctx, siteName, accessTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetByURL indicates an expected call of SetByURL.
func (mr *MockIRedisRepoUpdateMockRecorder) SetByURL(ctx, siteName, accessTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByURL", reflect.TypeOf((*MockIRedisRepoUpdate)(nil).SetByURL), ctx, siteName, accessTime)
}
