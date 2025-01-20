// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	context "context"

	state "github.com/fiwallets/zkevm-node/state"
	mock "github.com/stretchr/testify/mock"
)

// L1BlockCheckerIntegrator is an autogenerated mock type for the L1BlockCheckerIntegrator type
type L1BlockCheckerIntegrator struct {
	mock.Mock
}

type L1BlockCheckerIntegrator_Expecter struct {
	mock *mock.Mock
}

func (_m *L1BlockCheckerIntegrator) EXPECT() *L1BlockCheckerIntegrator_Expecter {
	return &L1BlockCheckerIntegrator_Expecter{mock: &_m.Mock}
}

// CheckReorgWrapper provides a mock function with given fields: ctx, reorgFirstBlockOk, errReportedByReorgFunc
func (_m *L1BlockCheckerIntegrator) CheckReorgWrapper(ctx context.Context, reorgFirstBlockOk *state.Block, errReportedByReorgFunc error) (*state.Block, error) {
	ret := _m.Called(ctx, reorgFirstBlockOk, errReportedByReorgFunc)

	if len(ret) == 0 {
		panic("no return value specified for CheckReorgWrapper")
	}

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Block, error) (*state.Block, error)); ok {
		return rf(ctx, reorgFirstBlockOk, errReportedByReorgFunc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Block, error) *state.Block); ok {
		r0 = rf(ctx, reorgFirstBlockOk, errReportedByReorgFunc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Block, error) error); ok {
		r1 = rf(ctx, reorgFirstBlockOk, errReportedByReorgFunc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L1BlockCheckerIntegrator_CheckReorgWrapper_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckReorgWrapper'
type L1BlockCheckerIntegrator_CheckReorgWrapper_Call struct {
	*mock.Call
}

// CheckReorgWrapper is a helper method to define mock.On call
//   - ctx context.Context
//   - reorgFirstBlockOk *state.Block
//   - errReportedByReorgFunc error
func (_e *L1BlockCheckerIntegrator_Expecter) CheckReorgWrapper(ctx interface{}, reorgFirstBlockOk interface{}, errReportedByReorgFunc interface{}) *L1BlockCheckerIntegrator_CheckReorgWrapper_Call {
	return &L1BlockCheckerIntegrator_CheckReorgWrapper_Call{Call: _e.mock.On("CheckReorgWrapper", ctx, reorgFirstBlockOk, errReportedByReorgFunc)}
}

func (_c *L1BlockCheckerIntegrator_CheckReorgWrapper_Call) Run(run func(ctx context.Context, reorgFirstBlockOk *state.Block, errReportedByReorgFunc error)) *L1BlockCheckerIntegrator_CheckReorgWrapper_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Block), args[2].(error))
	})
	return _c
}

func (_c *L1BlockCheckerIntegrator_CheckReorgWrapper_Call) Return(_a0 *state.Block, _a1 error) *L1BlockCheckerIntegrator_CheckReorgWrapper_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L1BlockCheckerIntegrator_CheckReorgWrapper_Call) RunAndReturn(run func(context.Context, *state.Block, error) (*state.Block, error)) *L1BlockCheckerIntegrator_CheckReorgWrapper_Call {
	_c.Call.Return(run)
	return _c
}

// OnResetState provides a mock function with given fields: ctx
func (_m *L1BlockCheckerIntegrator) OnResetState(ctx context.Context) {
	_m.Called(ctx)
}

// L1BlockCheckerIntegrator_OnResetState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnResetState'
type L1BlockCheckerIntegrator_OnResetState_Call struct {
	*mock.Call
}

// OnResetState is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L1BlockCheckerIntegrator_Expecter) OnResetState(ctx interface{}) *L1BlockCheckerIntegrator_OnResetState_Call {
	return &L1BlockCheckerIntegrator_OnResetState_Call{Call: _e.mock.On("OnResetState", ctx)}
}

func (_c *L1BlockCheckerIntegrator_OnResetState_Call) Run(run func(ctx context.Context)) *L1BlockCheckerIntegrator_OnResetState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L1BlockCheckerIntegrator_OnResetState_Call) Return() *L1BlockCheckerIntegrator_OnResetState_Call {
	_c.Call.Return()
	return _c
}

func (_c *L1BlockCheckerIntegrator_OnResetState_Call) RunAndReturn(run func(context.Context)) *L1BlockCheckerIntegrator_OnResetState_Call {
	_c.Call.Return(run)
	return _c
}

// OnStart provides a mock function with given fields: ctx
func (_m *L1BlockCheckerIntegrator) OnStart(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for OnStart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// L1BlockCheckerIntegrator_OnStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnStart'
type L1BlockCheckerIntegrator_OnStart_Call struct {
	*mock.Call
}

// OnStart is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L1BlockCheckerIntegrator_Expecter) OnStart(ctx interface{}) *L1BlockCheckerIntegrator_OnStart_Call {
	return &L1BlockCheckerIntegrator_OnStart_Call{Call: _e.mock.On("OnStart", ctx)}
}

func (_c *L1BlockCheckerIntegrator_OnStart_Call) Run(run func(ctx context.Context)) *L1BlockCheckerIntegrator_OnStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L1BlockCheckerIntegrator_OnStart_Call) Return(_a0 error) *L1BlockCheckerIntegrator_OnStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1BlockCheckerIntegrator_OnStart_Call) RunAndReturn(run func(context.Context) error) *L1BlockCheckerIntegrator_OnStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewL1BlockCheckerIntegrator creates a new instance of L1BlockCheckerIntegrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewL1BlockCheckerIntegrator(t interface {
	mock.TestingT
	Cleanup(func())
}) *L1BlockCheckerIntegrator {
	mock := &L1BlockCheckerIntegrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
