package product_test

import (
	"context"
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/internal/mocks"
	product_usecase "gokomodo/internal/usecase/product"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListProductSeller(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)
	products := []*entity.Product{product}

	err := errors.New("error")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("GetProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(products, nil).
		Twice()
	productRepo.
		On("CountProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(int32(len(products)), nil).
		Once()
	productRepo.
		On("GetProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(products, err).
		Once()
	productRepo.
		On("CountProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(int32(len(products)), err).
		Once()

	t.Run("OK", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, count, errUseCase := useCase.ListProductSeller(ctx, product.Seller.Id, nil)

		assert.Nil(t, errUseCase)
		assert.Equal(t, products, res)
		assert.Equal(t, int32(len(products)), count)
	})

	t.Run("ErrorGetProductSeller", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, count, errUseCase := useCase.ListProductSeller(ctx, product.Seller.Id, nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ErrorCountProductSeller", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, count, errUseCase := useCase.ListProductSeller(ctx, product.Seller.Id, nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
