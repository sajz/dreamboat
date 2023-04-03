// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/blocknative/dreamboat/beacon/client (interfaces: BeaconNode)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	client "github.com/blocknative/dreamboat/beacon/client"
	structs "github.com/blocknative/dreamboat/structs"
	gomock "github.com/golang/mock/gomock"
)

// MockBeaconNode is a mock of BeaconNode interface.
type MockBeaconNode struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeMockRecorder
}

// MockBeaconNodeMockRecorder is the mock recorder for MockBeaconNode.
type MockBeaconNodeMockRecorder struct {
	mock *MockBeaconNode
}

// NewMockBeaconNode creates a new mock instance.
func NewMockBeaconNode(ctrl *gomock.Controller) *MockBeaconNode {
	mock := &MockBeaconNode{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconNode) EXPECT() *MockBeaconNodeMockRecorder {
	return m.recorder
}

// Endpoint mocks base method.
func (m *MockBeaconNode) Endpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Endpoint indicates an expected call of Endpoint.
func (mr *MockBeaconNodeMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockBeaconNode)(nil).Endpoint))
}

// Genesis mocks base method.
func (m *MockBeaconNode) Genesis() (structs.GenesisInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(structs.GenesisInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Genesis indicates an expected call of Genesis.
func (mr *MockBeaconNodeMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockBeaconNode)(nil).Genesis))
}

// GetForkSchedule mocks base method.
func (m *MockBeaconNode) GetForkSchedule() (*client.GetForkScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForkSchedule")
	ret0, _ := ret[0].(*client.GetForkScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForkSchedule indicates an expected call of GetForkSchedule.
func (mr *MockBeaconNodeMockRecorder) GetForkSchedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForkSchedule", reflect.TypeOf((*MockBeaconNode)(nil).GetForkSchedule))
}

// GetProposerDuties mocks base method.
func (m *MockBeaconNode) GetProposerDuties(arg0 structs.Epoch) (*client.RegisteredProposersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposerDuties", arg0)
	ret0, _ := ret[0].(*client.RegisteredProposersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposerDuties indicates an expected call of GetProposerDuties.
func (mr *MockBeaconNodeMockRecorder) GetProposerDuties(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposerDuties", reflect.TypeOf((*MockBeaconNode)(nil).GetProposerDuties), arg0)
}

// GetWithdrawals mocks base method.
func (m *MockBeaconNode) GetWithdrawals(arg0 structs.Slot) (*client.GetWithdrawalsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithdrawals", arg0)
	ret0, _ := ret[0].(*client.GetWithdrawalsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithdrawals indicates an expected call of GetWithdrawals.
func (mr *MockBeaconNodeMockRecorder) GetWithdrawals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithdrawals", reflect.TypeOf((*MockBeaconNode)(nil).GetWithdrawals), arg0)
}

// KnownValidators mocks base method.
func (m *MockBeaconNode) KnownValidators(arg0 structs.Slot) (client.AllValidatorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KnownValidators", arg0)
	ret0, _ := ret[0].(client.AllValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KnownValidators indicates an expected call of KnownValidators.
func (mr *MockBeaconNodeMockRecorder) KnownValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownValidators", reflect.TypeOf((*MockBeaconNode)(nil).KnownValidators), arg0)
}

// PublishBlock mocks base method.
func (m *MockBeaconNode) PublishBlock(arg0 context.Context, arg1 structs.SignedBeaconBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishBlock indicates an expected call of PublishBlock.
func (mr *MockBeaconNodeMockRecorder) PublishBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishBlock", reflect.TypeOf((*MockBeaconNode)(nil).PublishBlock), arg0, arg1)
}

// Randao mocks base method.
func (m *MockBeaconNode) Randao(arg0 structs.Slot) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Randao", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Randao indicates an expected call of Randao.
func (mr *MockBeaconNodeMockRecorder) Randao(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Randao", reflect.TypeOf((*MockBeaconNode)(nil).Randao), arg0)
}

// SubscribeToHeadEvents mocks base method.
func (m *MockBeaconNode) SubscribeToHeadEvents(arg0 context.Context, arg1 chan client.HeadEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubscribeToHeadEvents", arg0, arg1)
}

// SubscribeToHeadEvents indicates an expected call of SubscribeToHeadEvents.
func (mr *MockBeaconNodeMockRecorder) SubscribeToHeadEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToHeadEvents", reflect.TypeOf((*MockBeaconNode)(nil).SubscribeToHeadEvents), arg0, arg1)
}

// SyncStatus mocks base method.
func (m *MockBeaconNode) SyncStatus() (*client.SyncStatusPayloadData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus")
	ret0, _ := ret[0].(*client.SyncStatusPayloadData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockBeaconNodeMockRecorder) SyncStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockBeaconNode)(nil).SyncStatus))
}
