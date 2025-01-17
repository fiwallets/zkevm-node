// Code generated by mockery v2.39.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/fiwallets/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pool "github.com/0xPolygonHermez/zkevm-node/pool"

	time "time"

	types "github.com/fiwallets/go-ethereum/core/types"
)

// PoolMock is an autogenerated mock type for the PoolInterface type
type PoolMock struct {
	mock.Mock
}

// AddTx provides a mock function with given fields: ctx, tx, ip
func (_m *PoolMock) AddTx(ctx context.Context, tx types.Transaction, ip string) error {
	ret := _m.Called(ctx, tx, ip)

	if len(ret) == 0 {
		panic("no return value specified for AddTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Transaction, string) error); ok {
		r0 = rf(ctx, tx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CalculateEffectiveGasPrice provides a mock function with given fields: rawTx, txGasPrice, txGasUsed, l1GasPrice, l2GasPrice
func (_m *PoolMock) CalculateEffectiveGasPrice(rawTx []byte, txGasPrice *big.Int, txGasUsed uint64, l1GasPrice uint64, l2GasPrice uint64) (*big.Int, error) {
	ret := _m.Called(rawTx, txGasPrice, txGasUsed, l1GasPrice, l2GasPrice)

	if len(ret) == 0 {
		panic("no return value specified for CalculateEffectiveGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, *big.Int, uint64, uint64, uint64) (*big.Int, error)); ok {
		return rf(rawTx, txGasPrice, txGasUsed, l1GasPrice, l2GasPrice)
	}
	if rf, ok := ret.Get(0).(func([]byte, *big.Int, uint64, uint64, uint64) *big.Int); ok {
		r0 = rf(rawTx, txGasPrice, txGasUsed, l1GasPrice, l2GasPrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, *big.Int, uint64, uint64, uint64) error); ok {
		r1 = rf(rawTx, txGasPrice, txGasUsed, l1GasPrice, l2GasPrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountPendingTransactions provides a mock function with given fields: ctx
func (_m *PoolMock) CountPendingTransactions(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CountPendingTransactions")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EffectiveGasPriceEnabled provides a mock function with given fields:
func (_m *PoolMock) EffectiveGasPriceEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EffectiveGasPriceEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetGasPrices provides a mock function with given fields: ctx
func (_m *PoolMock) GetGasPrices(ctx context.Context) (pool.GasPrices, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPrices")
	}

	var r0 pool.GasPrices
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pool.GasPrices, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pool.GasPrices); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(pool.GasPrices)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, address
func (_m *PoolMock) GetNonce(ctx context.Context, address common.Address) (uint64, error) {
	ret := _m.Called(ctx, address)

	if len(ret) == 0 {
		panic("no return value specified for GetNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingTxHashesSince provides a mock function with given fields: ctx, since
func (_m *PoolMock) GetPendingTxHashesSince(ctx context.Context, since time.Time) ([]common.Hash, error) {
	ret := _m.Called(ctx, since)

	if len(ret) == 0 {
		panic("no return value specified for GetPendingTxHashesSince")
	}

	var r0 []common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) ([]common.Hash, error)); ok {
		return rf(ctx, since)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []common.Hash); ok {
		r0 = rf(ctx, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingTxs provides a mock function with given fields: ctx, limit
func (_m *PoolMock) GetPendingTxs(ctx context.Context, limit uint64) ([]pool.Transaction, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetPendingTxs")
	}

	var r0 []pool.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]pool.Transaction, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []pool.Transaction); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pool.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByHash provides a mock function with given fields: ctx, hash
func (_m *PoolMock) GetTransactionByHash(ctx context.Context, hash common.Hash) (*pool.Transaction, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByHash")
	}

	var r0 *pool.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*pool.Transaction, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *pool.Transaction); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pool.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByL2Hash provides a mock function with given fields: ctx, hash
func (_m *PoolMock) GetTransactionByL2Hash(ctx context.Context, hash common.Hash) (*pool.Transaction, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByL2Hash")
	}

	var r0 *pool.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*pool.Transaction, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *pool.Transaction); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pool.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPoolMock creates a new instance of PoolMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPoolMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *PoolMock {
	mock := &PoolMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
