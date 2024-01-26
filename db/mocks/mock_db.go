// Code generated by MockGen. DO NOT EDIT.
// Source: Ewallet/db (interfaces: WalletDatabase)

// Package mock_db is a generated GoMock package.
package mock_db

import (
        db "Ewallet/db"
        models "Ewallet/models"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockWalletDatabase is a mock of WalletDatabase interface.
type MockWalletDatabase struct {
        ctrl     *gomock.Controller
        recorder *MockWalletDatabaseMockRecorder
}

// MockWalletDatabaseMockRecorder is the mock recorder for MockWalletDatabase.
type MockWalletDatabaseMockRecorder struct {
        mock *MockWalletDatabase
}

// NewMockWalletDatabase creates a new mock instance.
func NewMockWalletDatabase(ctrl *gomock.Controller) *MockWalletDatabase {
        mock := &MockWalletDatabase{ctrl: ctrl}
        mock.recorder = &MockWalletDatabaseMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletDatabase) EXPECT() *MockWalletDatabaseMockRecorder {
        return m.recorder
}

// Begin mocks base method.
func (m *MockWalletDatabase) Begin() db.WalletDatabase {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Begin")
        ret0, _ := ret[0].(db.WalletDatabase)
        return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockWalletDatabaseMockRecorder) Begin() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockWalletDatabase)(nil).Begin))
}

// Commit mocks base method.
func (m *MockWalletDatabase) Commit() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Commit")
        ret0, _ := ret[0].(error)
        return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockWalletDatabaseMockRecorder) Commit() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockWalletDatabase)(nil).Commit))
}

// CreateTransaction mocks base method.
func (m *MockWalletDatabase) CreateTransaction(arg0 *models.Transaction) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateTransaction", arg0)
        ret0, _ := ret[0].(error)
        return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockWalletDatabaseMockRecorder) CreateTransaction(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockWalletDatabase)(nil).CreateTransaction), arg0)
}

// CreateWallet mocks base method.
func (m *MockWalletDatabase) CreateWallet(arg0 *models.Wallet) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateWallet", arg0)
        ret0, _ := ret[0].(error)
        return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletDatabaseMockRecorder) CreateWallet(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletDatabase)(nil).CreateWallet), arg0)
}

// Error mocks base method.
func (m *MockWalletDatabase) Error() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Error")
        ret0, _ := ret[0].(error)
        return ret0
}

// Error indicates an expected call of Error.
func (mr *MockWalletDatabaseMockRecorder) Error() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockWalletDatabase)(nil).Error))
}

// GetTransactionHistory mocks base method.
func (m *MockWalletDatabase) GetTransactionHistory(arg0 string) ([]models.Transaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionHistory", arg0)
        ret0, _ := ret[0].([]models.Transaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetTransactionHistory indicates an expected call of GetTransactionHistory.
func (mr *MockWalletDatabaseMockRecorder) GetTransactionHistory(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionHistory", reflect.TypeOf((*MockWalletDatabase)(nil).GetTransactionHistory), arg0)
}

// GetWalletByID mocks base method.
func (m *MockWalletDatabase) GetWalletByID(arg0 string) (*models.Wallet, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetWalletByID", arg0)
        ret0, _ := ret[0].(*models.Wallet)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetWalletByID indicates an expected call of GetWalletByID.
func (mr *MockWalletDatabaseMockRecorder) GetWalletByID(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletByID", reflect.TypeOf((*MockWalletDatabase)(nil).GetWalletByID), arg0)
}

// Rollback mocks base method.
func (m *MockWalletDatabase) Rollback() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Rollback")
        ret0, _ := ret[0].(error)
        return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockWalletDatabaseMockRecorder) Rollback() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockWalletDatabase)(nil).Rollback))
}

// UpdateWalletBalance mocks base method.
func (m *MockWalletDatabase) UpdateWalletBalance(arg0 *models.Wallet) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateWalletBalance", arg0)
        ret0, _ := ret[0].(error)
        return ret0
}

// UpdateWalletBalance indicates an expected call of UpdateWalletBalance.
func (mr *MockWalletDatabaseMockRecorder) UpdateWalletBalance(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWalletBalance", reflect.TypeOf((*MockWalletDatabase)(nil).UpdateWalletBalance), arg0)
}