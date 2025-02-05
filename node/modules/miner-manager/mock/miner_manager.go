// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/venus-miner/node/modules/miner-manager (interfaces: MinerManageAPI)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	address "github.com/filecoin-project/go-address"
	types "github.com/filecoin-project/venus-miner/types"
	gomock "github.com/golang/mock/gomock"
)

// MockMinerManageAPI is a mock of MinerManageAPI interface.
type MockMinerManageAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMinerManageAPIMockRecorder
}

// MockMinerManageAPIMockRecorder is the mock recorder for MockMinerManageAPI.
type MockMinerManageAPIMockRecorder struct {
	mock *MockMinerManageAPI
}

// NewMockMinerManageAPI creates a new mock instance.
func NewMockMinerManageAPI(ctrl *gomock.Controller) *MockMinerManageAPI {
	mock := &MockMinerManageAPI{ctrl: ctrl}
	mock.recorder = &MockMinerManageAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinerManageAPI) EXPECT() *MockMinerManageAPIMockRecorder {
	return m.recorder
}

// CloseMining mocks base method.
func (m *MockMinerManageAPI) CloseMining(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseMining", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseMining indicates an expected call of CloseMining.
func (mr *MockMinerManageAPIMockRecorder) CloseMining(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseMining", reflect.TypeOf((*MockMinerManageAPI)(nil).CloseMining), arg0, arg1)
}

// Get mocks base method.
func (m *MockMinerManageAPI) Get(arg0 context.Context, arg1 address.Address) (*types.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*types.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMinerManageAPIMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMinerManageAPI)(nil).Get), arg0, arg1)
}

// Has mocks base method.
func (m *MockMinerManageAPI) Has(arg0 context.Context, arg1 address.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockMinerManageAPIMockRecorder) Has(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMinerManageAPI)(nil).Has), arg0, arg1)
}

// IsOpenMining mocks base method.
func (m *MockMinerManageAPI) IsOpenMining(arg0 context.Context, arg1 address.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOpenMining", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOpenMining indicates an expected call of IsOpenMining.
func (mr *MockMinerManageAPIMockRecorder) IsOpenMining(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOpenMining", reflect.TypeOf((*MockMinerManageAPI)(nil).IsOpenMining), arg0, arg1)
}

// List mocks base method.
func (m *MockMinerManageAPI) List(arg0 context.Context) (map[address.Address]*types.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(map[address.Address]*types.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMinerManageAPIMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMinerManageAPI)(nil).List), arg0)
}

// OpenMining mocks base method.
func (m *MockMinerManageAPI) OpenMining(arg0 context.Context, arg1 address.Address) (*types.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenMining", arg0, arg1)
	ret0, _ := ret[0].(*types.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenMining indicates an expected call of OpenMining.
func (mr *MockMinerManageAPIMockRecorder) OpenMining(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenMining", reflect.TypeOf((*MockMinerManageAPI)(nil).OpenMining), arg0, arg1)
}

// Update mocks base method.
func (m *MockMinerManageAPI) Update(arg0 context.Context, arg1, arg2 int64) (map[address.Address]*types.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[address.Address]*types.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMinerManageAPIMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMinerManageAPI)(nil).Update), arg0, arg1, arg2)
}
