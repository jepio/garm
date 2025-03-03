// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	params "github.com/cloudbase/garm/params"
	mock "github.com/stretchr/testify/mock"
)

// PoolManager is an autogenerated mock type for the PoolManager type
type PoolManager struct {
	mock.Mock
}

// DeleteRunner provides a mock function with given fields: runner, forceRemove
func (_m *PoolManager) DeleteRunner(runner params.Instance, forceRemove bool) error {
	ret := _m.Called(runner, forceRemove)

	var r0 error
	if rf, ok := ret.Get(0).(func(params.Instance, bool) error); ok {
		r0 = rf(runner, forceRemove)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDeleteRunner provides a mock function with given fields: runner
func (_m *PoolManager) ForceDeleteRunner(runner params.Instance) error {
	ret := _m.Called(runner)

	var r0 error
	if rf, ok := ret.Get(0).(func(params.Instance) error); ok {
		r0 = rf(runner)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWebhookInfo provides a mock function with given fields: ctx
func (_m *PoolManager) GetWebhookInfo(ctx context.Context) (params.HookInfo, error) {
	ret := _m.Called(ctx)

	var r0 params.HookInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (params.HookInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) params.HookInfo); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(params.HookInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GithubRunnerRegistrationToken provides a mock function with given fields:
func (_m *PoolManager) GithubRunnerRegistrationToken() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleWorkflowJob provides a mock function with given fields: job
func (_m *PoolManager) HandleWorkflowJob(job params.WorkflowJob) error {
	ret := _m.Called(job)

	var r0 error
	if rf, ok := ret.Get(0).(func(params.WorkflowJob) error); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *PoolManager) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InstallWebhook provides a mock function with given fields: ctx, param
func (_m *PoolManager) InstallWebhook(ctx context.Context, param params.InstallWebhookParams) (params.HookInfo, error) {
	ret := _m.Called(ctx, param)

	var r0 params.HookInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, params.InstallWebhookParams) (params.HookInfo, error)); ok {
		return rf(ctx, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, params.InstallWebhookParams) params.HookInfo); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(params.HookInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, params.InstallWebhookParams) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshState provides a mock function with given fields: param
func (_m *PoolManager) RefreshState(param params.UpdatePoolStateParams) error {
	ret := _m.Called(param)

	var r0 error
	if rf, ok := ret.Get(0).(func(params.UpdatePoolStateParams) error); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RootCABundle provides a mock function with given fields:
func (_m *PoolManager) RootCABundle() (params.CertificateBundle, error) {
	ret := _m.Called()

	var r0 params.CertificateBundle
	var r1 error
	if rf, ok := ret.Get(0).(func() (params.CertificateBundle, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() params.CertificateBundle); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(params.CertificateBundle)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *PoolManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *PoolManager) Status() params.PoolManagerStatus {
	ret := _m.Called()

	var r0 params.PoolManagerStatus
	if rf, ok := ret.Get(0).(func() params.PoolManagerStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(params.PoolManagerStatus)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *PoolManager) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UninstallWebhook provides a mock function with given fields: ctx
func (_m *PoolManager) UninstallWebhook(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wait provides a mock function with given fields:
func (_m *PoolManager) Wait() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebhookSecret provides a mock function with given fields:
func (_m *PoolManager) WebhookSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewPoolManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewPoolManager creates a new instance of PoolManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPoolManager(t mockConstructorTestingTNewPoolManager) *PoolManager {
	mock := &PoolManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
