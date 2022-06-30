package user_test

import (
	"context"
	"errors"
	"gokomodo/internal/mocks"
	user_usecase "gokomodo/internal/usecase/user"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	ctx := context.TODO()

	userRepo := new(mocks.UserRepositoryMock)

	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)
	user.Password = "qweasd123"
	user1 := testdata.NewUser(userDTO)

	err := errors.New("error")
	expectedErr := []error{
		err,
	}

	userRepo.
		On("FindUserByEmail", mock.Anything, user.Email).
		Return(user1, nil).
		Once()
	userRepo.
		On("FindUserByEmail", mock.Anything, user.Email).
		Return(user, err).
		Once()

	t.Run("OK", func(t *testing.T) {
		useCase := user_usecase.NewUserInteractor(userRepo)

		res, errUseCase := useCase.Login(ctx, user)

		assert.Nil(t, errUseCase)
		assert.Equal(t, user.Email, res.Email)
		assert.Equal(t, user1.Password, res.Password)
	})

	t.Run("ErrorFindUser", func(t *testing.T) {
		useCase := user_usecase.NewUserInteractor(userRepo)

		res, errUseCase := useCase.Login(ctx, user)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}
