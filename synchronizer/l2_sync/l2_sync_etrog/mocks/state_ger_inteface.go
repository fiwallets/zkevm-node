// Code generated by mockery. DO NOT EDIT.

package mock_l2_sync_etrog

import (
	context "context"

	common "github.com/fiwallets/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/fiwallets/zkevm-node/state"
)

// StateGERInteface is an autogenerated mock type for the StateGERInteface type
type StateGERInteface struct {
	mock.Mock
}

type StateGERInteface_Expecter struct {
	mock *mock.Mock
}

func (_m *StateGERInteface) EXPECT() *StateGERInteface_Expecter {
	return &StateGERInteface_Expecter{mock: &_m.Mock}
}

// GetExitRootByGlobalExitRoot provides a mock function with given fields: ctx, ger, dbTx
func (_m *StateGERInteface) GetExitRootByGlobalExitRoot(ctx context.Context, ger common.Hash, dbTx pgx.Tx) (*state.GlobalExitRoot, error) {
	ret := _m.Called(ctx, ger, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetExitRootByGlobalExitRoot")
	}

	var r0 *state.GlobalExitRoot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) (*state.GlobalExitRoot, error)); ok {
		return rf(ctx, ger, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *state.GlobalExitRoot); ok {
		r0 = rf(ctx, ger, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.GlobalExitRoot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, ger, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateGERInteface_GetExitRootByGlobalExitRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExitRootByGlobalExitRoot'
type StateGERInteface_GetExitRootByGlobalExitRoot_Call struct {
	*mock.Call
}

// GetExitRootByGlobalExitRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - ger common.Hash
//   - dbTx pgx.Tx
func (_e *StateGERInteface_Expecter) GetExitRootByGlobalExitRoot(ctx interface{}, ger interface{}, dbTx interface{}) *StateGERInteface_GetExitRootByGlobalExitRoot_Call {
	return &StateGERInteface_GetExitRootByGlobalExitRoot_Call{Call: _e.mock.On("GetExitRootByGlobalExitRoot", ctx, ger, dbTx)}
}

func (_c *StateGERInteface_GetExitRootByGlobalExitRoot_Call) Run(run func(ctx context.Context, ger common.Hash, dbTx pgx.Tx)) *StateGERInteface_GetExitRootByGlobalExitRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateGERInteface_GetExitRootByGlobalExitRoot_Call) Return(_a0 *state.GlobalExitRoot, _a1 error) *StateGERInteface_GetExitRootByGlobalExitRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateGERInteface_GetExitRootByGlobalExitRoot_Call) RunAndReturn(run func(context.Context, common.Hash, pgx.Tx) (*state.GlobalExitRoot, error)) *StateGERInteface_GetExitRootByGlobalExitRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *StateGERInteface) GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastBlock")
	}

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateGERInteface_GetLastBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastBlock'
type StateGERInteface_GetLastBlock_Call struct {
	*mock.Call
}

// GetLastBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx pgx.Tx
func (_e *StateGERInteface_Expecter) GetLastBlock(ctx interface{}, dbTx interface{}) *StateGERInteface_GetLastBlock_Call {
	return &StateGERInteface_GetLastBlock_Call{Call: _e.mock.On("GetLastBlock", ctx, dbTx)}
}

func (_c *StateGERInteface_GetLastBlock_Call) Run(run func(ctx context.Context, dbTx pgx.Tx)) *StateGERInteface_GetLastBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Tx))
	})
	return _c
}

func (_c *StateGERInteface_GetLastBlock_Call) Return(_a0 *state.Block, _a1 error) *StateGERInteface_GetLastBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateGERInteface_GetLastBlock_Call) RunAndReturn(run func(context.Context, pgx.Tx) (*state.Block, error)) *StateGERInteface_GetLastBlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewStateGERInteface creates a new instance of StateGERInteface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateGERInteface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateGERInteface {
	mock := &StateGERInteface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
