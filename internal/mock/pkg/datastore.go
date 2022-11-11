// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mock_relay is a generated GoMock package.
package mock_relay

import (
	context "context"
	reflect "reflect"
	time "time"

	relay "github.com/blocknative/dreamboat/pkg"
	types "github.com/flashbots/go-boost-utils/types"
	gomock "github.com/golang/mock/gomock"
	datastore "github.com/ipfs/go-datastore"
)

// MockDatastore is a mock of Datastore interface.
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore.
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance.
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// GetDelivered mocks base method.
func (m *MockDatastore) GetDelivered(arg0 context.Context, arg1 relay.Query) (relay.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelivered", arg0, arg1)
	ret0, _ := ret[0].(relay.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelivered indicates an expected call of GetDelivered.
func (mr *MockDatastoreMockRecorder) GetDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelivered", reflect.TypeOf((*MockDatastore)(nil).GetDelivered), arg0, arg1)
}

// GetDeliveredBatch mocks base method.
func (m *MockDatastore) GetDeliveredBatch(arg0 context.Context, arg1 []relay.Query) ([]relay.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveredBatch", arg0, arg1)
	ret0, _ := ret[0].([]relay.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveredBatch indicates an expected call of GetDeliveredBatch.
func (mr *MockDatastoreMockRecorder) GetDeliveredBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveredBatch", reflect.TypeOf((*MockDatastore)(nil).GetDeliveredBatch), arg0, arg1)
}

// GetHeaderBatch mocks base method.
func (m *MockDatastore) GetHeaderBatch(arg0 context.Context, arg1 []relay.Query) ([]relay.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderBatch", arg0, arg1)
	ret0, _ := ret[0].([]relay.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeaderBatch indicates an expected call of GetHeaderBatch.
func (mr *MockDatastoreMockRecorder) GetHeaderBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderBatch", reflect.TypeOf((*MockDatastore)(nil).GetHeaderBatch), arg0, arg1)
}

// GetHeaders mocks base method.
func (m *MockDatastore) GetHeaders(arg0 context.Context, arg1 relay.Query) ([]relay.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaders", arg0, arg1)
	ret0, _ := ret[0].([]relay.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeaders indicates an expected call of GetHeaders.
func (mr *MockDatastoreMockRecorder) GetHeaders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaders", reflect.TypeOf((*MockDatastore)(nil).GetHeaders), arg0, arg1)
}

// GetMaxProfitHeadersDesc mocks base method.
func (m *MockDatastore) GetMaxProfitHeadersDesc(arg0 context.Context, arg1 relay.Slot) ([]relay.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxProfitHeadersDesc", arg0, arg1)
	ret0, _ := ret[0].([]relay.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxProfitHeadersDesc indicates an expected call of GetMaxProfitHeadersDesc.
func (mr *MockDatastoreMockRecorder) GetMaxProfitHeadersDesc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxProfitHeadersDesc", reflect.TypeOf((*MockDatastore)(nil).GetMaxProfitHeadersDesc), arg0, arg1)
}

// GetPayload mocks base method.
func (m *MockDatastore) GetPayload(arg0 context.Context, arg1 relay.PayloadKey) (*relay.BlockBidAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", arg0, arg1)
	ret0, _ := ret[0].(*relay.BlockBidAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayload indicates an expected call of GetPayload.
func (mr *MockDatastoreMockRecorder) GetPayload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockDatastore)(nil).GetPayload), arg0, arg1)
}

// GetRegistration mocks base method.
func (m *MockDatastore) GetRegistration(arg0 context.Context, arg1 relay.PubKey) (types.SignedValidatorRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistration", arg0, arg1)
	ret0, _ := ret[0].(types.SignedValidatorRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistration indicates an expected call of GetRegistration.
func (mr *MockDatastoreMockRecorder) GetRegistration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistration", reflect.TypeOf((*MockDatastore)(nil).GetRegistration), arg0, arg1)
}

// PutDelivered mocks base method.
func (m *MockDatastore) PutDelivered(arg0 context.Context, arg1 relay.Slot, arg2 relay.DeliveredTrace, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutDelivered", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutDelivered indicates an expected call of PutDelivered.
func (mr *MockDatastoreMockRecorder) PutDelivered(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutDelivered", reflect.TypeOf((*MockDatastore)(nil).PutDelivered), arg0, arg1, arg2, arg3)
}

// PutHeader mocks base method.
func (m *MockDatastore) PutHeader(arg0 context.Context, arg1 relay.Slot, arg2 relay.HeaderAndTrace, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutHeader", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutHeader indicates an expected call of PutHeader.
func (mr *MockDatastoreMockRecorder) PutHeader(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutHeader", reflect.TypeOf((*MockDatastore)(nil).PutHeader), arg0, arg1, arg2, arg3)
}

// PutPayload mocks base method.
func (m *MockDatastore) PutPayload(arg0 context.Context, arg1 relay.PayloadKey, arg2 *relay.BlockBidAndTrace, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPayload", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutPayload indicates an expected call of PutPayload.
func (mr *MockDatastoreMockRecorder) PutPayload(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPayload", reflect.TypeOf((*MockDatastore)(nil).PutPayload), arg0, arg1, arg2, arg3)
}

// PutRegistration mocks base method.
func (m *MockDatastore) PutRegistration(arg0 context.Context, arg1 relay.PubKey, arg2 types.SignedValidatorRegistration, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRegistration", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutRegistration indicates an expected call of PutRegistration.
func (mr *MockDatastoreMockRecorder) PutRegistration(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRegistration", reflect.TypeOf((*MockDatastore)(nil).PutRegistration), arg0, arg1, arg2, arg3)
}

// MockTTLStorage is a mock of TTLStorage interface.
type MockTTLStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTTLStorageMockRecorder
}

// MockTTLStorageMockRecorder is the mock recorder for MockTTLStorage.
type MockTTLStorageMockRecorder struct {
	mock *MockTTLStorage
}

// NewMockTTLStorage creates a new mock instance.
func NewMockTTLStorage(ctrl *gomock.Controller) *MockTTLStorage {
	mock := &MockTTLStorage{ctrl: ctrl}
	mock.recorder = &MockTTLStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTTLStorage) EXPECT() *MockTTLStorageMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTTLStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTTLStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTTLStorage)(nil).Close))
}

// Get mocks base method.
func (m *MockTTLStorage) Get(arg0 context.Context, arg1 datastore.Key) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTTLStorageMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTTLStorage)(nil).Get), arg0, arg1)
}

// GetBatch mocks base method.
func (m *MockTTLStorage) GetBatch(ctx context.Context, keys []datastore.Key) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBatch", ctx, keys)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBatch indicates an expected call of GetBatch.
func (mr *MockTTLStorageMockRecorder) GetBatch(ctx, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBatch", reflect.TypeOf((*MockTTLStorage)(nil).GetBatch), ctx, keys)
}

// PutWithTTL mocks base method.
func (m *MockTTLStorage) PutWithTTL(arg0 context.Context, arg1 datastore.Key, arg2 []byte, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutWithTTL", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutWithTTL indicates an expected call of PutWithTTL.
func (mr *MockTTLStorageMockRecorder) PutWithTTL(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutWithTTL", reflect.TypeOf((*MockTTLStorage)(nil).PutWithTTL), arg0, arg1, arg2, arg3)
}
