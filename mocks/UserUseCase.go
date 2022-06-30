// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "gokomodo/domain/entity"

	exceptions "gokomodo/pkg/exceptions"

	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, user
func (_m *UserUseCase) Login(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomerError) {
	ret := _m.Called(ctx, user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 *exceptions.CustomerError
	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) *exceptions.CustomerError); ok {
		r1 = rf(ctx, user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exceptions.CustomerError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}