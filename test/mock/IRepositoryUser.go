// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-user/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryUser is an autogenerated mock type for the IRepositoryUser type
type IRepositoryUser struct {
	mock.Mock
}

// FindUserByEmailAndPassword provides a mock function with given fields: email, pass
func (_m *IRepositoryUser) FindUserByEmailAndPassword(email string, pass string) *entity.User {
	ret := _m.Called(email, pass)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string, string) *entity.User); ok {
		r0 = rf(email, pass)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	return r0
}

// RegisterUser provides a mock function with given fields: user
func (_m *IRepositoryUser) RegisterUser(user entity.User) (*entity.User, error) {
	ret := _m.Called(user)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (*entity.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entity.User) *entity.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: user
func (_m *IRepositoryUser) UpdateUser(user entity.User) (*entity.User, error) {
	ret := _m.Called(user)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (*entity.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entity.User) *entity.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryUser creates a new instance of IRepositoryUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryUser(t mockConstructorTestingTNewIRepositoryUser) *IRepositoryUser {
	mock := &IRepositoryUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}