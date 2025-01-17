// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	executor "github.com/fiwallets/zkevm-node/state/runtime/executor"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ExecutorServiceClientMock is an autogenerated mock type for the ExecutorServiceClient type
type ExecutorServiceClientMock struct {
	mock.Mock
}

type ExecutorServiceClientMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ExecutorServiceClientMock) EXPECT() *ExecutorServiceClientMock_Expecter {
	return &ExecutorServiceClientMock_Expecter{mock: &_m.Mock}
}

// GetFlushStatus provides a mock function with given fields: ctx, in, opts
func (_m *ExecutorServiceClientMock) GetFlushStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*executor.GetFlushStatusResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFlushStatus")
	}

	var r0 *executor.GetFlushStatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*executor.GetFlushStatusResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) *executor.GetFlushStatusResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executor.GetFlushStatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutorServiceClientMock_GetFlushStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlushStatus'
type ExecutorServiceClientMock_GetFlushStatus_Call struct {
	*mock.Call
}

// GetFlushStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - in *emptypb.Empty
//   - opts ...grpc.CallOption
func (_e *ExecutorServiceClientMock_Expecter) GetFlushStatus(ctx interface{}, in interface{}, opts ...interface{}) *ExecutorServiceClientMock_GetFlushStatus_Call {
	return &ExecutorServiceClientMock_GetFlushStatus_Call{Call: _e.mock.On("GetFlushStatus",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ExecutorServiceClientMock_GetFlushStatus_Call) Run(run func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption)) *ExecutorServiceClientMock_GetFlushStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*emptypb.Empty), variadicArgs...)
	})
	return _c
}

func (_c *ExecutorServiceClientMock_GetFlushStatus_Call) Return(_a0 *executor.GetFlushStatusResponse, _a1 error) *ExecutorServiceClientMock_GetFlushStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExecutorServiceClientMock_GetFlushStatus_Call) RunAndReturn(run func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*executor.GetFlushStatusResponse, error)) *ExecutorServiceClientMock_GetFlushStatus_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessBatch provides a mock function with given fields: ctx, in, opts
func (_m *ExecutorServiceClientMock) ProcessBatch(ctx context.Context, in *executor.ProcessBatchRequest, opts ...grpc.CallOption) (*executor.ProcessBatchResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ProcessBatch")
	}

	var r0 *executor.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *executor.ProcessBatchRequest, ...grpc.CallOption) (*executor.ProcessBatchResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *executor.ProcessBatchRequest, ...grpc.CallOption) *executor.ProcessBatchResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executor.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *executor.ProcessBatchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutorServiceClientMock_ProcessBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessBatch'
type ExecutorServiceClientMock_ProcessBatch_Call struct {
	*mock.Call
}

// ProcessBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - in *executor.ProcessBatchRequest
//   - opts ...grpc.CallOption
func (_e *ExecutorServiceClientMock_Expecter) ProcessBatch(ctx interface{}, in interface{}, opts ...interface{}) *ExecutorServiceClientMock_ProcessBatch_Call {
	return &ExecutorServiceClientMock_ProcessBatch_Call{Call: _e.mock.On("ProcessBatch",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ExecutorServiceClientMock_ProcessBatch_Call) Run(run func(ctx context.Context, in *executor.ProcessBatchRequest, opts ...grpc.CallOption)) *ExecutorServiceClientMock_ProcessBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*executor.ProcessBatchRequest), variadicArgs...)
	})
	return _c
}

func (_c *ExecutorServiceClientMock_ProcessBatch_Call) Return(_a0 *executor.ProcessBatchResponse, _a1 error) *ExecutorServiceClientMock_ProcessBatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExecutorServiceClientMock_ProcessBatch_Call) RunAndReturn(run func(context.Context, *executor.ProcessBatchRequest, ...grpc.CallOption) (*executor.ProcessBatchResponse, error)) *ExecutorServiceClientMock_ProcessBatch_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessBatchV2 provides a mock function with given fields: ctx, in, opts
func (_m *ExecutorServiceClientMock) ProcessBatchV2(ctx context.Context, in *executor.ProcessBatchRequestV2, opts ...grpc.CallOption) (*executor.ProcessBatchResponseV2, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ProcessBatchV2")
	}

	var r0 *executor.ProcessBatchResponseV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *executor.ProcessBatchRequestV2, ...grpc.CallOption) (*executor.ProcessBatchResponseV2, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *executor.ProcessBatchRequestV2, ...grpc.CallOption) *executor.ProcessBatchResponseV2); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executor.ProcessBatchResponseV2)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *executor.ProcessBatchRequestV2, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutorServiceClientMock_ProcessBatchV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessBatchV2'
type ExecutorServiceClientMock_ProcessBatchV2_Call struct {
	*mock.Call
}

// ProcessBatchV2 is a helper method to define mock.On call
//   - ctx context.Context
//   - in *executor.ProcessBatchRequestV2
//   - opts ...grpc.CallOption
func (_e *ExecutorServiceClientMock_Expecter) ProcessBatchV2(ctx interface{}, in interface{}, opts ...interface{}) *ExecutorServiceClientMock_ProcessBatchV2_Call {
	return &ExecutorServiceClientMock_ProcessBatchV2_Call{Call: _e.mock.On("ProcessBatchV2",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ExecutorServiceClientMock_ProcessBatchV2_Call) Run(run func(ctx context.Context, in *executor.ProcessBatchRequestV2, opts ...grpc.CallOption)) *ExecutorServiceClientMock_ProcessBatchV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*executor.ProcessBatchRequestV2), variadicArgs...)
	})
	return _c
}

func (_c *ExecutorServiceClientMock_ProcessBatchV2_Call) Return(_a0 *executor.ProcessBatchResponseV2, _a1 error) *ExecutorServiceClientMock_ProcessBatchV2_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExecutorServiceClientMock_ProcessBatchV2_Call) RunAndReturn(run func(context.Context, *executor.ProcessBatchRequestV2, ...grpc.CallOption) (*executor.ProcessBatchResponseV2, error)) *ExecutorServiceClientMock_ProcessBatchV2_Call {
	_c.Call.Return(run)
	return _c
}

// NewExecutorServiceClientMock creates a new instance of ExecutorServiceClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutorServiceClientMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExecutorServiceClientMock {
	mock := &ExecutorServiceClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
