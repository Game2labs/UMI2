// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	txmgr "github.com/ethereum-optimism/optimism/umi-service/txmgr"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TxManager is an autogenerated mock type for the TxManager type
type TxManager struct {
	mock.Mock
}

// From provides a mock function with given fields:
func (_m *TxManager) From() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// Send provides a mock function with given fields: ctx, candidate
func (_m *TxManager) Send(ctx context.Context, candidate txmgr.TxCandidate) (*types.Receipt, error) {
	ret := _m.Called(ctx, candidate)

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgr.TxCandidate) (*types.Receipt, error)); ok {
		return rf(ctx, candidate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgr.TxCandidate) *types.Receipt); ok {
		r0 = rf(ctx, candidate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgr.TxCandidate) error); ok {
		r1 = rf(ctx, candidate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTxManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxManager creates a new instance of TxManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxManager(t mockConstructorTestingTNewTxManager) *TxManager {
	mock := &TxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
