// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	zap "go.uber.org/zap"
)

// logKvFunc is an autogenerated mock type for the logKvFunc type
type logKvFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, msg, keyAndValues
func (_m *logKvFunc) Execute(_a0 *zap.SugaredLogger, msg string, keyAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0, msg)
	_ca = append(_ca, keyAndValues...)
	_m.Called(_ca...)
}

type mockConstructorTestingTnewLogKvFunc interface {
	mock.TestingT
	Cleanup(func())
}

// newLogKvFunc creates a new instance of logKvFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newLogKvFunc(t mockConstructorTestingTnewLogKvFunc) *logKvFunc {
	mock := &logKvFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
