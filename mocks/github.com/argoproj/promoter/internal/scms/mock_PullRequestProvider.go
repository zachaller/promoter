// Code generated by mockery v2.42.2. DO NOT EDIT.

package scms

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/zachaller/promoter/api/v1alpha1"
)

// MockPullRequestProvider is an autogenerated mock type for the PullRequestProvider type
type MockPullRequestProvider struct {
	mock.Mock
}

type MockPullRequestProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPullRequestProvider) EXPECT() *MockPullRequestProvider_Expecter {
	return &MockPullRequestProvider_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx, pullRequest
func (_m *MockPullRequestProvider) Close(ctx context.Context, pullRequest *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error) {
	ret := _m.Called(ctx, pullRequest)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 *v1alpha1.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)); ok {
		return rf(ctx, pullRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.PullRequest) *v1alpha1.PullRequest); ok {
		r0 = rf(ctx, pullRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.PullRequest) error); ok {
		r1 = rf(ctx, pullRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPullRequestProvider_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPullRequestProvider_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
//   - pullRequest *v1alpha1.PullRequest
func (_e *MockPullRequestProvider_Expecter) Close(ctx interface{}, pullRequest interface{}) *MockPullRequestProvider_Close_Call {
	return &MockPullRequestProvider_Close_Call{Call: _e.mock.On("Close", ctx, pullRequest)}
}

func (_c *MockPullRequestProvider_Close_Call) Run(run func(ctx context.Context, pullRequest *v1alpha1.PullRequest)) *MockPullRequestProvider_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1alpha1.PullRequest))
	})
	return _c
}

func (_c *MockPullRequestProvider_Close_Call) Return(_a0 *v1alpha1.PullRequest, _a1 error) *MockPullRequestProvider_Close_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPullRequestProvider_Close_Call) RunAndReturn(run func(context.Context, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)) *MockPullRequestProvider_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, title, head, base, description, pullRequest
func (_m *MockPullRequestProvider) Create(ctx context.Context, title string, head string, base string, description string, pullRequest *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error) {
	ret := _m.Called(ctx, title, head, base, description, pullRequest)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *v1alpha1.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)); ok {
		return rf(ctx, title, head, base, description, pullRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, *v1alpha1.PullRequest) *v1alpha1.PullRequest); ok {
		r0 = rf(ctx, title, head, base, description, pullRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, *v1alpha1.PullRequest) error); ok {
		r1 = rf(ctx, title, head, base, description, pullRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPullRequestProvider_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPullRequestProvider_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - title string
//   - head string
//   - base string
//   - description string
//   - pullRequest *v1alpha1.PullRequest
func (_e *MockPullRequestProvider_Expecter) Create(ctx interface{}, title interface{}, head interface{}, base interface{}, description interface{}, pullRequest interface{}) *MockPullRequestProvider_Create_Call {
	return &MockPullRequestProvider_Create_Call{Call: _e.mock.On("Create", ctx, title, head, base, description, pullRequest)}
}

func (_c *MockPullRequestProvider_Create_Call) Run(run func(ctx context.Context, title string, head string, base string, description string, pullRequest *v1alpha1.PullRequest)) *MockPullRequestProvider_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(*v1alpha1.PullRequest))
	})
	return _c
}

func (_c *MockPullRequestProvider_Create_Call) Return(_a0 *v1alpha1.PullRequest, _a1 error) *MockPullRequestProvider_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPullRequestProvider_Create_Call) RunAndReturn(run func(context.Context, string, string, string, string, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)) *MockPullRequestProvider_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Merge provides a mock function with given fields: ctx, commitMessage, pullRequest
func (_m *MockPullRequestProvider) Merge(ctx context.Context, commitMessage string, pullRequest *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error) {
	ret := _m.Called(ctx, commitMessage, pullRequest)

	if len(ret) == 0 {
		panic("no return value specified for Merge")
	}

	var r0 *v1alpha1.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)); ok {
		return rf(ctx, commitMessage, pullRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *v1alpha1.PullRequest) *v1alpha1.PullRequest); ok {
		r0 = rf(ctx, commitMessage, pullRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *v1alpha1.PullRequest) error); ok {
		r1 = rf(ctx, commitMessage, pullRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPullRequestProvider_Merge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Merge'
type MockPullRequestProvider_Merge_Call struct {
	*mock.Call
}

// Merge is a helper method to define mock.On call
//   - ctx context.Context
//   - commitMessage string
//   - pullRequest *v1alpha1.PullRequest
func (_e *MockPullRequestProvider_Expecter) Merge(ctx interface{}, commitMessage interface{}, pullRequest interface{}) *MockPullRequestProvider_Merge_Call {
	return &MockPullRequestProvider_Merge_Call{Call: _e.mock.On("Merge", ctx, commitMessage, pullRequest)}
}

func (_c *MockPullRequestProvider_Merge_Call) Run(run func(ctx context.Context, commitMessage string, pullRequest *v1alpha1.PullRequest)) *MockPullRequestProvider_Merge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*v1alpha1.PullRequest))
	})
	return _c
}

func (_c *MockPullRequestProvider_Merge_Call) Return(_a0 *v1alpha1.PullRequest, _a1 error) *MockPullRequestProvider_Merge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPullRequestProvider_Merge_Call) RunAndReturn(run func(context.Context, string, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)) *MockPullRequestProvider_Merge_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, title, description, pullRequest
func (_m *MockPullRequestProvider) Update(ctx context.Context, title string, description string, pullRequest *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error) {
	ret := _m.Called(ctx, title, description, pullRequest)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *v1alpha1.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)); ok {
		return rf(ctx, title, description, pullRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *v1alpha1.PullRequest) *v1alpha1.PullRequest); ok {
		r0 = rf(ctx, title, description, pullRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *v1alpha1.PullRequest) error); ok {
		r1 = rf(ctx, title, description, pullRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPullRequestProvider_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPullRequestProvider_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - title string
//   - description string
//   - pullRequest *v1alpha1.PullRequest
func (_e *MockPullRequestProvider_Expecter) Update(ctx interface{}, title interface{}, description interface{}, pullRequest interface{}) *MockPullRequestProvider_Update_Call {
	return &MockPullRequestProvider_Update_Call{Call: _e.mock.On("Update", ctx, title, description, pullRequest)}
}

func (_c *MockPullRequestProvider_Update_Call) Run(run func(ctx context.Context, title string, description string, pullRequest *v1alpha1.PullRequest)) *MockPullRequestProvider_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*v1alpha1.PullRequest))
	})
	return _c
}

func (_c *MockPullRequestProvider_Update_Call) Return(_a0 *v1alpha1.PullRequest, _a1 error) *MockPullRequestProvider_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPullRequestProvider_Update_Call) RunAndReturn(run func(context.Context, string, string, *v1alpha1.PullRequest) (*v1alpha1.PullRequest, error)) *MockPullRequestProvider_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPullRequestProvider creates a new instance of MockPullRequestProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPullRequestProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPullRequestProvider {
	mock := &MockPullRequestProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
