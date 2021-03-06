// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "gokomodo/domain/entity"

	mock "github.com/stretchr/testify/mock"

	request "gokomodo/internal/delivery/request"

	uuid "github.com/google/uuid"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// CountProduct provides a mock function with given fields: ctx, options
func (_m *ProductRepository) CountProduct(ctx context.Context, options *request.Option) (int32, error) {
	ret := _m.Called(ctx, options)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, *request.Option) int32); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.Option) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountProductSeller provides a mock function with given fields: ctx, userId, options
func (_m *ProductRepository) CountProductSeller(ctx context.Context, userId uuid.UUID, options *request.Option) (int32, error) {
	ret := _m.Called(ctx, userId, options)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *request.Option) int32); ok {
		r0 = rf(ctx, userId, options)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, *request.Option) error); ok {
		r1 = rf(ctx, userId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProductById provides a mock function with given fields: ctx, productId
func (_m *ProductRepository) FindProductById(ctx context.Context, productId uuid.UUID) (*entity.Product, error) {
	ret := _m.Called(ctx, productId)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Product); ok {
		r0 = rf(ctx, productId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: ctx, options
func (_m *ProductRepository) GetProduct(ctx context.Context, options *request.Option) ([]*entity.Product, error) {
	ret := _m.Called(ctx, options)

	var r0 []*entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, *request.Option) []*entity.Product); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.Option) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductSeller provides a mock function with given fields: ctx, userId, options
func (_m *ProductRepository) GetProductSeller(ctx context.Context, userId uuid.UUID, options *request.Option) ([]*entity.Product, error) {
	ret := _m.Called(ctx, userId, options)

	var r0 []*entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *request.Option) []*entity.Product); ok {
		r0 = rf(ctx, userId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, *request.Option) error); ok {
		r1 = rf(ctx, userId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreProduct provides a mock function with given fields: ctx, product
func (_m *ProductRepository) StoreProduct(ctx context.Context, product *entity.Product) error {
	ret := _m.Called(ctx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
