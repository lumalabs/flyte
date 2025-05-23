// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	flyteidlcore "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"
)

// TaskReader is an autogenerated mock type for the TaskReader type
type TaskReader struct {
	mock.Mock
}

type TaskReader_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskReader) EXPECT() *TaskReader_Expecter {
	return &TaskReader_Expecter{mock: &_m.Mock}
}

// Path provides a mock function with given fields: ctx
func (_m *TaskReader) Path(ctx context.Context) (storage.DataReference, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Path")
	}

	var r0 storage.DataReference
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (storage.DataReference, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) storage.DataReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskReader_Path_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Path'
type TaskReader_Path_Call struct {
	*mock.Call
}

// Path is a helper method to define mock.On call
//   - ctx context.Context
func (_e *TaskReader_Expecter) Path(ctx interface{}) *TaskReader_Path_Call {
	return &TaskReader_Path_Call{Call: _e.mock.On("Path", ctx)}
}

func (_c *TaskReader_Path_Call) Run(run func(ctx context.Context)) *TaskReader_Path_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *TaskReader_Path_Call) Return(_a0 storage.DataReference, _a1 error) *TaskReader_Path_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TaskReader_Path_Call) RunAndReturn(run func(context.Context) (storage.DataReference, error)) *TaskReader_Path_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx
func (_m *TaskReader) Read(ctx context.Context) (*flyteidlcore.TaskTemplate, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *flyteidlcore.TaskTemplate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*flyteidlcore.TaskTemplate, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *flyteidlcore.TaskTemplate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flyteidlcore.TaskTemplate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskReader_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type TaskReader_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
func (_e *TaskReader_Expecter) Read(ctx interface{}) *TaskReader_Read_Call {
	return &TaskReader_Read_Call{Call: _e.mock.On("Read", ctx)}
}

func (_c *TaskReader_Read_Call) Run(run func(ctx context.Context)) *TaskReader_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *TaskReader_Read_Call) Return(_a0 *flyteidlcore.TaskTemplate, _a1 error) *TaskReader_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TaskReader_Read_Call) RunAndReturn(run func(context.Context) (*flyteidlcore.TaskTemplate, error)) *TaskReader_Read_Call {
	_c.Call.Return(run)
	return _c
}

// NewTaskReader creates a new instance of TaskReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskReader {
	mock := &TaskReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
