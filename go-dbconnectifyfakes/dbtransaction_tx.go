// Code generated by MockGen. DO NOT EDIT.
// Source: dbtransaction.go
//
// Generated by this command:
//
//	mockgen -source=dbtransaction.go -destination=go-dbconnectifyfakes/dbtransaction_tx.go -package=godbconnectifyfakes
//

// Package godbconnectifyfakes is a generated GoMock package.
package godbconnectifyfakes

import (
	context "context"
	reflect "reflect"

	go_dbconnectify "github.com/SyaibanAhmadRamadhan/go-dbconnectify"
	gomock "go.uber.org/mock/gomock"
)

// MockTx is a mock of Tx interface.
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx.
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance.
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// DoTransaction mocks base method.
func (m *MockTx) DoTransaction(ctx context.Context, opt *go_dbconnectify.TxOption, fn func(context.Context) (bool, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTransaction", ctx, opt, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoTransaction indicates an expected call of DoTransaction.
func (mr *MockTxMockRecorder) DoTransaction(ctx, opt, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTransaction", reflect.TypeOf((*MockTx)(nil).DoTransaction), ctx, opt, fn)
}

// DoTransactionx mocks base method.
func (m *MockTx) DoTransactionx(ctx context.Context, opt *go_dbconnectify.TxOption, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTransactionx", ctx, opt, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoTransactionx indicates an expected call of DoTransactionx.
func (mr *MockTxMockRecorder) DoTransactionx(ctx, opt, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTransactionx", reflect.TypeOf((*MockTx)(nil).DoTransactionx), ctx, opt, fn)
}
