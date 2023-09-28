// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-user/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryProgress is an autogenerated mock type for the IRepositoryProgress type
type IRepositoryProgress struct {
	mock.Mock
}

// FindProgressByUser provides a mock function with given fields: idUser
func (_m *IRepositoryProgress) FindProgressByUser(idUser uint64) ([]entity.Progress, error) {
	ret := _m.Called(idUser)

	var r0 []entity.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]entity.Progress, error)); ok {
		return rf(idUser)
	}
	if rf, ok := ret.Get(0).(func(uint64) []entity.Progress); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProgressByUserAndCategory provides a mock function with given fields: idUser, idCategory
func (_m *IRepositoryProgress) FindProgressByUserAndCategory(idUser uint64, idCategory uint64) (*entity.Progress, error) {
	ret := _m.Called(idUser, idCategory)

	var r0 *entity.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) (*entity.Progress, error)); ok {
		return rf(idUser, idCategory)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64) *entity.Progress); ok {
		r0 = rf(idUser, idCategory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(idUser, idCategory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterProgress provides a mock function with given fields: progress
func (_m *IRepositoryProgress) RegisterProgress(progress entity.Progress) (*entity.Progress, error) {
	ret := _m.Called(progress)

	var r0 *entity.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Progress) (*entity.Progress, error)); ok {
		return rf(progress)
	}
	if rf, ok := ret.Get(0).(func(entity.Progress) *entity.Progress); ok {
		r0 = rf(progress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Progress) error); ok {
		r1 = rf(progress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProgress provides a mock function with given fields: progress
func (_m *IRepositoryProgress) UpdateProgress(progress entity.Progress) (*entity.Progress, error) {
	ret := _m.Called(progress)

	var r0 *entity.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Progress) (*entity.Progress, error)); ok {
		return rf(progress)
	}
	if rf, ok := ret.Get(0).(func(entity.Progress) *entity.Progress); ok {
		r0 = rf(progress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Progress) error); ok {
		r1 = rf(progress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryProgress interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryProgress creates a new instance of IRepositoryProgress. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryProgress(t mockConstructorTestingTNewIRepositoryProgress) *IRepositoryProgress {
	mock := &IRepositoryProgress{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}