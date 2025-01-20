// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	context "context"

	event "github.com/fiwallets/zkevm-node/event"
	mock "github.com/stretchr/testify/mock"
)

// EventLogInterface is an autogenerated mock type for the EventLogInterface type
type EventLogInterface struct {
	mock.Mock
}

type EventLogInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EventLogInterface) EXPECT() *EventLogInterface_Expecter {
	return &EventLogInterface_Expecter{mock: &_m.Mock}
}

// LogEvent provides a mock function with given fields: ctx, _a1
func (_m *EventLogInterface) LogEvent(ctx context.Context, _a1 *event.Event) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LogEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventLogInterface_LogEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogEvent'
type EventLogInterface_LogEvent_Call struct {
	*mock.Call
}

// LogEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *event.Event
func (_e *EventLogInterface_Expecter) LogEvent(ctx interface{}, _a1 interface{}) *EventLogInterface_LogEvent_Call {
	return &EventLogInterface_LogEvent_Call{Call: _e.mock.On("LogEvent", ctx, _a1)}
}

func (_c *EventLogInterface_LogEvent_Call) Run(run func(ctx context.Context, _a1 *event.Event)) *EventLogInterface_LogEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*event.Event))
	})
	return _c
}

func (_c *EventLogInterface_LogEvent_Call) Return(_a0 error) *EventLogInterface_LogEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventLogInterface_LogEvent_Call) RunAndReturn(run func(context.Context, *event.Event) error) *EventLogInterface_LogEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewEventLogInterface creates a new instance of EventLogInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventLogInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventLogInterface {
	mock := &EventLogInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
