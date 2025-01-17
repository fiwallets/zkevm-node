// Code generated by mockery. DO NOT EDIT.

package mock_elderberry

import (
	context "context"

	etherman "github.com/fiwallets/zkevm-node/etherman"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	time "time"
)

// PreviousProcessor is an autogenerated mock type for the PreviousProcessor type
type PreviousProcessor struct {
	mock.Mock
}

type PreviousProcessor_Expecter struct {
	mock *mock.Mock
}

func (_m *PreviousProcessor) EXPECT() *PreviousProcessor_Expecter {
	return &PreviousProcessor_Expecter{mock: &_m.Mock}
}

// Process provides a mock function with given fields: ctx, order, l1Block, dbTx
func (_m *PreviousProcessor) Process(ctx context.Context, order etherman.Order, l1Block *etherman.Block, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, order, l1Block, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Process")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, etherman.Order, *etherman.Block, pgx.Tx) error); ok {
		r0 = rf(ctx, order, l1Block, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreviousProcessor_Process_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Process'
type PreviousProcessor_Process_Call struct {
	*mock.Call
}

// Process is a helper method to define mock.On call
//   - ctx context.Context
//   - order etherman.Order
//   - l1Block *etherman.Block
//   - dbTx pgx.Tx
func (_e *PreviousProcessor_Expecter) Process(ctx interface{}, order interface{}, l1Block interface{}, dbTx interface{}) *PreviousProcessor_Process_Call {
	return &PreviousProcessor_Process_Call{Call: _e.mock.On("Process", ctx, order, l1Block, dbTx)}
}

func (_c *PreviousProcessor_Process_Call) Run(run func(ctx context.Context, order etherman.Order, l1Block *etherman.Block, dbTx pgx.Tx)) *PreviousProcessor_Process_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(etherman.Order), args[2].(*etherman.Block), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *PreviousProcessor_Process_Call) Return(_a0 error) *PreviousProcessor_Process_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PreviousProcessor_Process_Call) RunAndReturn(run func(context.Context, etherman.Order, *etherman.Block, pgx.Tx) error) *PreviousProcessor_Process_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessSequenceBatches provides a mock function with given fields: ctx, sequencedBatches, blockNumber, l1BlockTimestamp, dbTx
func (_m *PreviousProcessor) ProcessSequenceBatches(ctx context.Context, sequencedBatches []etherman.SequencedBatch, blockNumber uint64, l1BlockTimestamp time.Time, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, sequencedBatches, blockNumber, l1BlockTimestamp, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for ProcessSequenceBatches")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []etherman.SequencedBatch, uint64, time.Time, pgx.Tx) error); ok {
		r0 = rf(ctx, sequencedBatches, blockNumber, l1BlockTimestamp, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreviousProcessor_ProcessSequenceBatches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessSequenceBatches'
type PreviousProcessor_ProcessSequenceBatches_Call struct {
	*mock.Call
}

// ProcessSequenceBatches is a helper method to define mock.On call
//   - ctx context.Context
//   - sequencedBatches []etherman.SequencedBatch
//   - blockNumber uint64
//   - l1BlockTimestamp time.Time
//   - dbTx pgx.Tx
func (_e *PreviousProcessor_Expecter) ProcessSequenceBatches(ctx interface{}, sequencedBatches interface{}, blockNumber interface{}, l1BlockTimestamp interface{}, dbTx interface{}) *PreviousProcessor_ProcessSequenceBatches_Call {
	return &PreviousProcessor_ProcessSequenceBatches_Call{Call: _e.mock.On("ProcessSequenceBatches", ctx, sequencedBatches, blockNumber, l1BlockTimestamp, dbTx)}
}

func (_c *PreviousProcessor_ProcessSequenceBatches_Call) Run(run func(ctx context.Context, sequencedBatches []etherman.SequencedBatch, blockNumber uint64, l1BlockTimestamp time.Time, dbTx pgx.Tx)) *PreviousProcessor_ProcessSequenceBatches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]etherman.SequencedBatch), args[2].(uint64), args[3].(time.Time), args[4].(pgx.Tx))
	})
	return _c
}

func (_c *PreviousProcessor_ProcessSequenceBatches_Call) Return(_a0 error) *PreviousProcessor_ProcessSequenceBatches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PreviousProcessor_ProcessSequenceBatches_Call) RunAndReturn(run func(context.Context, []etherman.SequencedBatch, uint64, time.Time, pgx.Tx) error) *PreviousProcessor_ProcessSequenceBatches_Call {
	_c.Call.Return(run)
	return _c
}

// NewPreviousProcessor creates a new instance of PreviousProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPreviousProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *PreviousProcessor {
	mock := &PreviousProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
