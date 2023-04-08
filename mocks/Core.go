// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	zapper "github.com/GoFarsi/zapper"
	mock "github.com/stretchr/testify/mock"
)

// Core is an autogenerated mock type for the Core type
type Core struct {
	mock.Mock
}

// init provides a mock function with given fields: _a0
func (_m *Core) init(_a0 *zapper.Zap) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*zapper.Zap) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCore interface {
	mock.TestingT
	Cleanup(func())
}

// NewCore creates a new instance of Core. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCore(t mockConstructorTestingTNewCore) *Core {
	mock := &Core{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
