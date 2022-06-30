package order_test

import (
	"context"
	"errors"
	"gokomodo/domain/entity"
	order_usecase "gokomodo/internal/usecase/order"
	"gokomodo/mocks"
	"gokomodo/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListOrder(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepository)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)
	orders := []*entity.Order{order}

	err := errors.New("error")
	expectedErr := []error{
		err,
	}

	orderRepo.
		On("GetOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
		Return(orders, nil).
		Twice()
	orderRepo.
		On("CountOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
		Return(int32(len(orders)), nil).
		Once()
	orderRepo.
		On("GetOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
		Return(orders, nil).
		Twice()
	orderRepo.
		On("CountOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
		Return(int32(len(orders)), nil).
		Once()
	orderRepo.
		On("CountOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
		Return(int32(0), err).
		Once()
	orderRepo.
		On("GetOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
		Return(orders, err).
		Once()
	orderRepo.
		On("CountOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
		Return(int32(0), err).
		Once()
	orderRepo.
		On("GetOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
		Return(orders, err).
		Once()

	t.Run("ListOrderSeller", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, err := useCase.ListOrder(ctx, order.Seller.Id, "SELLER", nil)

		assert.Nil(t, err)
		assert.Equal(t, orders, res)
		assert.Equal(t, int32(len(orders)), count)
	})

	t.Run("ListOrderBuyer", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, err := useCase.ListOrder(ctx, order.Buyer.Id, "BUYER", nil)

		assert.Nil(t, err)
		assert.Equal(t, orders, res)
		assert.Equal(t, int32(len(orders)), count)
	})

	t.Run("ErrorRole", func(t *testing.T) {
		err := errors.New("role not found")
		expectedErr := []error{
			err,
		}

		useCase := order_usecase.NewOrderInteractor(nil, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Seller.Id, "GUEST", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorCountOrderSeller", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Seller.Id, "SELLER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorGetOrderSeller", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Seller.Id, "SELLER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorCountOrderBuyer", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Buyer.Id, "BUYER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorGetOrderBuyer", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Buyer.Id, "BUYER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
