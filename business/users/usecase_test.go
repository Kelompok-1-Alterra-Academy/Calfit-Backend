package users_test

import (
	"CalFit/app/middlewares"
	"CalFit/business/users"
	"CalFit/business/users/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo mocks.Repository
var domain users.Domain
var usecase users.Usecase
var configJWT middlewares.ConfigJWT

func testSetup() {
	configJWT = middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("secret"),
		ExpiresDuration: viper.GetInt("expired"),
	}
	domain = users.Domain{
		Id:               1,
		Email:            "dummy@gmail.com",
		Password:         "dummy123",
		Photo:            "dummy",
		MembershipTypeID: 1,
		AddressID:        1,
		Token:            "dummy",
		BookingDetails:   domain.BookingDetails,
		Address:          domain.Address,
	}
	usecase = users.NewUsersUsecase(&repo, time.Minute*1, &configJWT)
}

func TestLogin(t *testing.T) {
	testSetup()
	t.Run("Test case 1 | Valid get", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
		user, err := usecase.Login(context.Background(), domain)
		user.Token = "dummy"
		assert.Nil(t, err)
		assert.Equal(t, domain, user)
	})
	t.Run("Test case 2 | Server error", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.AnythingOfType("Domain")).Return(users.Domain{}, errors.New("Internal server error")).Once()
		user, err := usecase.Login(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
	t.Run("Test case 3 | Empty email or password", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.AnythingOfType("Domain")).Return(users.Domain{}, errors.New("Email or password is empty")).Once()
		domain.Email = ""
		domain.Password = ""
		user, err := usecase.Login(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
}
