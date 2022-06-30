package order_test

import (
	"context"
	"errors"
	"gokomodo/internal/mocks"
	order_usecase "gokomodo/internal/usecase/order"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)
	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	err := errors.New("error")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("FindProductById", mock.Anything, order.Product.Id).
		Return(product, nil).
		Twice()
	orderRepo.
		On("StoreOrder", mock.Anything, order).
		Return(nil).
		Once()
	productRepo.
		On("FindProductById", mock.Anything, order.Product.Id).
		Return(product, err).
		Once()
	orderRepo.
		On("StoreOrder", mock.Anything, order).
		Return(err).
		Once()

	t.Run("OK", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, productRepo)

		res, err := useCase.CreateOrder(ctx, order)

		assert.Nil(t, err)
		assert.Equal(t, order, res)
	})

	t.Run("ErrorFindProduct", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, productRepo)

		res, errUseCase := useCase.CreateOrder(ctx, order)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorStoreProduct", func(t *testing.T) {
		useCase := order_usecase.NewOrderInteractor(orderRepo, productRepo)

		res, errUseCase := useCase.CreateOrder(ctx, order)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
