// Code generated by mockery. DO NOT EDIT.

package mock_l2_shared

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/fiwallets/zkevm-node/state"
)

// statePostClosedBatchCheckL2Block is an autogenerated mock type for the statePostClosedBatchCheckL2Block type
type statePostClosedBatchCheckL2Block struct {
	mock.Mock
}

type statePostClosedBatchCheckL2Block_Expecter struct {
	mock *mock.Mock
}

func (_m *statePostClosedBatchCheckL2Block) EXPECT() *statePostClosedBatchCheckL2Block_Expecter {
	return &statePostClosedBatchCheckL2Block_Expecter{mock: &_m.Mock}
}

// GetLastL2BlockByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *statePostClosedBatchCheckL2Block) GetLastL2BlockByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.L2Block, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastL2BlockByBatchNumber")
	}

	var r0 *state.L2Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.L2Block, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.L2Block); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.L2Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastL2BlockByBatchNumber'
type statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call struct {
	*mock.Call
}

// GetLastL2BlockByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx pgx.Tx
func (_e *statePostClosedBatchCheckL2Block_Expecter) GetLastL2BlockByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call {
	return &statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call{Call: _e.mock.On("GetLastL2BlockByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx pgx.Tx)) *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call) Return(_a0 *state.L2Block, _a1 error) *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, pgx.Tx) (*state.L2Block, error)) *statePostClosedBatchCheckL2Block_GetLastL2BlockByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// newStatePostClosedBatchCheckL2Block creates a new instance of statePostClosedBatchCheckL2Block. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStatePostClosedBatchCheckL2Block(t interface {
	mock.TestingT
	Cleanup(func())
}) *statePostClosedBatchCheckL2Block {
	mock := &statePostClosedBatchCheckL2Block{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
