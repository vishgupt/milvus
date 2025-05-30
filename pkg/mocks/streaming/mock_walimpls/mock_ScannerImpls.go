// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock_walimpls

import (
	message "github.com/milvus-io/milvus/pkg/v2/streaming/util/message"
	mock "github.com/stretchr/testify/mock"
)

// MockScannerImpls is an autogenerated mock type for the ScannerImpls type
type MockScannerImpls struct {
	mock.Mock
}

type MockScannerImpls_Expecter struct {
	mock *mock.Mock
}

func (_m *MockScannerImpls) EXPECT() *MockScannerImpls_Expecter {
	return &MockScannerImpls_Expecter{mock: &_m.Mock}
}

// Chan provides a mock function with no fields
func (_m *MockScannerImpls) Chan() <-chan message.ImmutableMessage {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Chan")
	}

	var r0 <-chan message.ImmutableMessage
	if rf, ok := ret.Get(0).(func() <-chan message.ImmutableMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan message.ImmutableMessage)
		}
	}

	return r0
}

// MockScannerImpls_Chan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chan'
type MockScannerImpls_Chan_Call struct {
	*mock.Call
}

// Chan is a helper method to define mock.On call
func (_e *MockScannerImpls_Expecter) Chan() *MockScannerImpls_Chan_Call {
	return &MockScannerImpls_Chan_Call{Call: _e.mock.On("Chan")}
}

func (_c *MockScannerImpls_Chan_Call) Run(run func()) *MockScannerImpls_Chan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScannerImpls_Chan_Call) Return(_a0 <-chan message.ImmutableMessage) *MockScannerImpls_Chan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScannerImpls_Chan_Call) RunAndReturn(run func() <-chan message.ImmutableMessage) *MockScannerImpls_Chan_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with no fields
func (_m *MockScannerImpls) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScannerImpls_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockScannerImpls_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockScannerImpls_Expecter) Close() *MockScannerImpls_Close_Call {
	return &MockScannerImpls_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockScannerImpls_Close_Call) Run(run func()) *MockScannerImpls_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScannerImpls_Close_Call) Return(_a0 error) *MockScannerImpls_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScannerImpls_Close_Call) RunAndReturn(run func() error) *MockScannerImpls_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Done provides a mock function with no fields
func (_m *MockScannerImpls) Done() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Done")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// MockScannerImpls_Done_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Done'
type MockScannerImpls_Done_Call struct {
	*mock.Call
}

// Done is a helper method to define mock.On call
func (_e *MockScannerImpls_Expecter) Done() *MockScannerImpls_Done_Call {
	return &MockScannerImpls_Done_Call{Call: _e.mock.On("Done")}
}

func (_c *MockScannerImpls_Done_Call) Run(run func()) *MockScannerImpls_Done_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScannerImpls_Done_Call) Return(_a0 <-chan struct{}) *MockScannerImpls_Done_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScannerImpls_Done_Call) RunAndReturn(run func() <-chan struct{}) *MockScannerImpls_Done_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with no fields
func (_m *MockScannerImpls) Error() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Error")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScannerImpls_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockScannerImpls_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
func (_e *MockScannerImpls_Expecter) Error() *MockScannerImpls_Error_Call {
	return &MockScannerImpls_Error_Call{Call: _e.mock.On("Error")}
}

func (_c *MockScannerImpls_Error_Call) Run(run func()) *MockScannerImpls_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScannerImpls_Error_Call) Return(_a0 error) *MockScannerImpls_Error_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScannerImpls_Error_Call) RunAndReturn(run func() error) *MockScannerImpls_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *MockScannerImpls) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockScannerImpls_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockScannerImpls_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockScannerImpls_Expecter) Name() *MockScannerImpls_Name_Call {
	return &MockScannerImpls_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockScannerImpls_Name_Call) Run(run func()) *MockScannerImpls_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScannerImpls_Name_Call) Return(_a0 string) *MockScannerImpls_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScannerImpls_Name_Call) RunAndReturn(run func() string) *MockScannerImpls_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockScannerImpls creates a new instance of MockScannerImpls. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScannerImpls(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScannerImpls {
	mock := &MockScannerImpls{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
