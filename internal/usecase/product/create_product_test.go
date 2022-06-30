package product_test

import (
	"context"
	"errors"
	product_usecase "gokomodo/internal/usecase/product"
	"gokomodo/mocks"
	"gokomodo/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepository)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	err := errors.New("error")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("StoreProduct", mock.Anything, product).
		Return(nil).
		Once()
	productRepo.
		On("StoreProduct", mock.Anything, product).
		Return(err).
		Once()

	t.Run("OK", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, errUseCase := useCase.CreateProduct(ctx, product)

		assert.Nil(t, errUseCase)
		assert.Equal(t, product, res)
	})

	t.Run("ErrorStoreProduct", func(t *testing.T) {
		useCase := product_usecase.NewProductInteractor(productRepo)

		res, errUseCase := useCase.CreateProduct(ctx, product)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
