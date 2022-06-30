package product_test

import (
	"context"
	"errors"
	"gokomodo/domain/entity"
	product_usecase "gokomodo/internal/usecase/product"
	"gokomodo/mocks"
	"gokomodo/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListProductBuyer(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepository)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)
	products := []*entity.Product{product}

	err := errors.New("error")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("GetProduct", mock.Anything, mock.Anything).
		Return(products, nil).
		Twice()
	productRepo.
		On("CountProduct", mock.Anything, mock.Anything).
		Return(int32(len(products)), nil).
		Once()
	productRepo.
		On("GetProduct", mock.Anything, mock.Anything).
		Return(products, err).
		Once()
	productRepo.
		On("CountProduct", mock.Anything, mock.Anything).
		Return(int32(0), err).
		Twice()

	t.Run("OK", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, count, err := useCase.ListProductBuyer(ctx, nil)

		assert.Nil(t, err)
		assert.Equal(t, products, res)
		assert.Equal(t, int32(len(products)), count)
	})
	t.Run("ErrorGetProduct", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, count, errUseCase := useCase.ListProductBuyer(ctx, nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
	t.Run("ErrorCountProduct", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, count, errUseCase := useCase.ListProductBuyer(ctx, nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
