// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/paularah/wallet/pkg/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/paularah/wallet/pkg/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddWalletBalance mocks base method.
func (m *MockStore) AddWalletBalance(arg0 context.Context, arg1 db.AddWalletBalanceParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWalletBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWalletBalance indicates an expected call of AddWalletBalance.
func (mr *MockStoreMockRecorder) AddWalletBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWalletBalance", reflect.TypeOf((*MockStore)(nil).AddWalletBalance), arg0, arg1)
}

// CreateFunding mocks base method.
func (m *MockStore) CreateFunding(arg0 context.Context, arg1 db.CreateFundingParams) (db.Funding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFunding", arg0, arg1)
	ret0, _ := ret[0].(db.Funding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFunding indicates an expected call of CreateFunding.
func (mr *MockStoreMockRecorder) CreateFunding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFunding", reflect.TypeOf((*MockStore)(nil).CreateFunding), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateWallet mocks base method.
func (m *MockStore) CreateWallet(arg0 context.Context, arg1 db.CreateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockStoreMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockStore)(nil).CreateWallet), arg0, arg1)
}

// CreateWalletEntry mocks base method.
func (m *MockStore) CreateWalletEntry(arg0 context.Context, arg1 db.CreateWalletEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWalletEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWalletEntry indicates an expected call of CreateWalletEntry.
func (mr *MockStoreMockRecorder) CreateWalletEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWalletEntry", reflect.TypeOf((*MockStore)(nil).CreateWalletEntry), arg0, arg1)
}

// DeleteWallet mocks base method.
func (m *MockStore) DeleteWallet(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet.
func (mr *MockStoreMockRecorder) DeleteWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockStore)(nil).DeleteWallet), arg0, arg1)
}

// GetFunding mocks base method.
func (m *MockStore) GetFunding(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunding", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunding indicates an expected call of GetFunding.
func (mr *MockStoreMockRecorder) GetFunding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunding", reflect.TypeOf((*MockStore)(nil).GetFunding), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserFromEmail mocks base method.
func (m *MockStore) GetUserFromEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFromEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFromEmail indicates an expected call of GetUserFromEmail.
func (mr *MockStoreMockRecorder) GetUserFromEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFromEmail", reflect.TypeOf((*MockStore)(nil).GetUserFromEmail), arg0, arg1)
}

// GetWallet mocks base method.
func (m *MockStore) GetWallet(arg0 context.Context, arg1 int64) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockStoreMockRecorder) GetWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockStore)(nil).GetWallet), arg0, arg1)
}

// GetWalletEntry mocks base method.
func (m *MockStore) GetWalletEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletEntry indicates an expected call of GetWalletEntry.
func (mr *MockStoreMockRecorder) GetWalletEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletEntry", reflect.TypeOf((*MockStore)(nil).GetWalletEntry), arg0, arg1)
}

// GetWalletForUpdate mocks base method.
func (m *MockStore) GetWalletForUpdate(arg0 context.Context, arg1 int64) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletForUpdate indicates an expected call of GetWalletForUpdate.
func (mr *MockStoreMockRecorder) GetWalletForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletForUpdate", reflect.TypeOf((*MockStore)(nil).GetWalletForUpdate), arg0, arg1)
}

// GetWalletFunding mocks base method.
func (m *MockStore) GetWalletFunding(arg0 context.Context, arg1 db.GetWalletFundingParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletFunding", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletFunding indicates an expected call of GetWalletFunding.
func (mr *MockStoreMockRecorder) GetWalletFunding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletFunding", reflect.TypeOf((*MockStore)(nil).GetWalletFunding), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), arg0, arg1)
}

// ListWalletEntries mocks base method.
func (m *MockStore) ListWalletEntries(arg0 context.Context, arg1 db.ListWalletEntriesParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWalletEntries", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWalletEntries indicates an expected call of ListWalletEntries.
func (mr *MockStoreMockRecorder) ListWalletEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWalletEntries", reflect.TypeOf((*MockStore)(nil).ListWalletEntries), arg0, arg1)
}

// ListWalletFunding mocks base method.
func (m *MockStore) ListWalletFunding(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWalletFunding", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWalletFunding indicates an expected call of ListWalletFunding.
func (mr *MockStoreMockRecorder) ListWalletFunding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWalletFunding", reflect.TypeOf((*MockStore)(nil).ListWalletFunding), arg0, arg1)
}

// ListWallets mocks base method.
func (m *MockStore) ListWallets(arg0 context.Context, arg1 db.ListWalletsParams) ([]db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWallets", arg0, arg1)
	ret0, _ := ret[0].([]db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWallets indicates an expected call of ListWallets.
func (mr *MockStoreMockRecorder) ListWallets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWallets", reflect.TypeOf((*MockStore)(nil).ListWallets), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TranferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateWallet mocks base method.
func (m *MockStore) UpdateWallet(arg0 context.Context, arg1 db.UpdateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWallet indicates an expected call of UpdateWallet.
func (mr *MockStoreMockRecorder) UpdateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWallet", reflect.TypeOf((*MockStore)(nil).UpdateWallet), arg0, arg1)
}
