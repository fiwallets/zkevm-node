// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
	mock "github.com/stretchr/testify/mock"

	state "github.com/fiwallets/zkevm-node/state"
)

// StateGetBatchByNumberInterface is an autogenerated mock type for the StateGetBatchByNumberInterface type
type StateGetBatchByNumberInterface struct {
	mock.Mock
}

type StateGetBatchByNumberInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StateGetBatchByNumberInterface) EXPECT() *StateGetBatchByNumberInterface_Expecter {
	return &StateGetBatchByNumberInterface_Expecter{mock: &_m.Mock}
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateGetBatchByNumberInterface) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetBatchByNumber")
	}

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateGetBatchByNumberInterface_GetBatchByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBatchByNumber'
type StateGetBatchByNumberInterface_GetBatchByNumber_Call struct {
	*mock.Call
}

// GetBatchByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx pgx.Tx
func (_e *StateGetBatchByNumberInterface_Expecter) GetBatchByNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *StateGetBatchByNumberInterface_GetBatchByNumber_Call {
	return &StateGetBatchByNumberInterface_GetBatchByNumber_Call{Call: _e.mock.On("GetBatchByNumber", ctx, batchNumber, dbTx)}
}

func (_c *StateGetBatchByNumberInterface_GetBatchByNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx pgx.Tx)) *StateGetBatchByNumberInterface_GetBatchByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateGetBatchByNumberInterface_GetBatchByNumber_Call) Return(_a0 *state.Batch, _a1 error) *StateGetBatchByNumberInterface_GetBatchByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateGetBatchByNumberInterface_GetBatchByNumber_Call) RunAndReturn(run func(context.Context, uint64, pgx.Tx) (*state.Batch, error)) *StateGetBatchByNumberInterface_GetBatchByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewStateGetBatchByNumberInterface creates a new instance of StateGetBatchByNumberInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateGetBatchByNumberInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateGetBatchByNumberInterface {
	mock := &StateGetBatchByNumberInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
