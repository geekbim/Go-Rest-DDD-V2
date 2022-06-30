package order_test

import (
	"context"
	"errors"
	order_usecase "gokomodo/internal/usecase/order"
	"gokomodo/mocks"
	"gokomodo/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAcceptOrder(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepository)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	err := errors.New("update order status failed")
	expectedErr := []error{
		err,
	}

	orderRepo.
		On("FindOrderBySellerIdAndOrderId", mock.Anything, order.Seller.Id, order.Id).
		Return(order, nil).
		Twice()
	orderRepo.
		On("UpdateOrderStatus", mock.Anything, order).
		Return(nil).
		Once()
	orderRepo.
		On("FindOrderBySellerIdAndOrderId", mock.Anything, order.Seller.Id, order.Id).
		Return(order, err).
		Once()
	orderRepo.
		On("UpdateOrderStatus", mock.Anything, order).
		Return(err).
		Once()

	t.Run("OK", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, err := useCase.AcceptOrder(ctx, order.Seller.Id, order.Id)

		assert.Nil(t, err)
		assert.Equal(t, order, res)
	})

	t.Run("ErrorFindOrder", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, errUseCase := useCase.AcceptOrder(ctx, order.Seller.Id, order.Id)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorUpdateOrderStatus", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, errUseCase := useCase.AcceptOrder(ctx, order.Seller.Id, order.Id)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
