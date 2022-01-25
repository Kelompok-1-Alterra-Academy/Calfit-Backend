package superadmins_test

import (
	"CalFit/app/middlewares"
	"CalFit/business/superadmins"
	"CalFit/business/superadmins/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo mocks.Repository
var domain, updatedDomain superadmins.Domain
var usecase superadmins.Usecase
var configJWT middlewares.ConfigJWT

func testSetup() {
	configJWT = middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("secret"),
		ExpiresDuration: viper.GetInt("expired"),
	}
	domain = superadmins.Domain{
		Id:       1,
		Username: "superadmin",
		Password: "calf1t_c4lfit",
		Token:    "dummy",
	}
	updatedDomain = superadmins.Domain{
		Id:          1,
		Username:    "superadmin",
		Password:    "calf1t_c4lfit",
		NewPassword: "calf1t_c4lfit_new",
		Token:       "dummy",
	}
	usecase = superadmins.NewSuperadminsUsecase(&repo, time.Minute*1, &configJWT)
}

// func TestLoginOAuth(t *testing.T) {
// 	testSetup()
// 	t.Run("Test case 1 | Valid get", func(t *testing.T) {
// 		repo.On("LoginOAuth", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
// 		user, err := usecase.LoginOAuth(context.Background(), domain)
// 		user.Token = "dummy"
// 		assert.Nil(t, err)
// 		assert.Equal(t, domain, user)
// 	})
// 	t.Run("Test case 2 | Server error", func(t *testing.T) {
// 		repo.On("LoginOAuth", mock.Anything, mock.AnythingOfType("Domain")).Return(superadmins.Domain{}, errors.New("Internal server error")).Once()
// 		user, err := usecase.LoginOAuth(context.Background(), domain)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, domain, user)
// 	})
// 	t.Run("Test case 3 | Empty username or password", func(t *testing.T) {
// 		repo.On("LoginOAuth", mock.Anything, mock.AnythingOfType("Domain")).Return(superadmins.Domain{}, errors.New("Username or password is empty")).Once()
// 		domain.Username = ""
// 		domain.Password = ""
// 		user, err := usecase.LoginOAuth(context.Background(), domain)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, domain, user)
// 	})
// }

func TestRegister(t *testing.T) {
	testSetup()
	repo.On("Register", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil)
	repo.On("GetAll", mock.Anything).Return([]superadmins.Domain{}, nil)
	// t.Run("Test Case 1 | Valid Create New Superadmin", func(t *testing.T) {
	// 	superadmin, err := usecase.Register(context.Background(), domain)
	// 	if err != nil {
	// 		t.Errorf("Error: %s", err)
	// 	}
	// 	if superadmin.Id == 0 {
	// 		t.Errorf("Error: %s", "superadmin Id is empty")
	// 	}
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, domain, superadmin)
	// })
	// t.Run("Test case 1 | Valid register", func(t *testing.T) {
	// 	repo.On("Register", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
	// 	repo.On("GetAll", mock.Anything).Return([]superadmins.Domain{}, nil)
	// 	user, err := usecase.Register(context.Background(), domain)
	// 	user.Token = "dummy"
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, domain, user)
	// })
	t.Run("Test case 2 | Server error", func(t *testing.T) {
		// repo.On("Register", mock.Anything, mock.AnythingOfType("Domain")).Return(superadmins.Domain{}, errors.New("Internal server error")).Once()
		user, err := usecase.Register(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
	t.Run("Test case 3 | Empty username or password", func(t *testing.T) {
		// repo.On("Register", mock.Anything, mock.AnythingOfType("Domain")).Return(superadmins.Domain{}, errors.New("Username or password is empty")).Once()
		domain.Username = ""
		domain.Password = ""
		user, err := usecase.Register(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
}

func TestLogin(t *testing.T) {
	testSetup()
	t.Run("Test case 1 | Invalid credentials", func(t *testing.T) {
		repo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain, errors.New("invalid credentials")).Once()
		user, err := usecase.Login(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
	t.Run("Test case 2 | Server error", func(t *testing.T) {
		repo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()
		user, err := usecase.Login(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
	t.Run("Test case 3 | Empty username or password", func(t *testing.T) {
		repo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(superadmins.Domain{}, errors.New("Username or password is empty")).Once()
		domain.Username = ""
		domain.Password = ""
		user, err := usecase.Login(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
}

func TestUpdatePassword(t *testing.T) {
	testSetup()
	// t.Run("Test case 1 | Valid update password", func(t *testing.T) {
	// 	repo.On("Login", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
	// 	repo.On("UpdatePassword", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
	// 	superadmin, err := usecase.UpdatePassword(context.Background(), updatedDomain)
	// 	isEqual := helpers.ValidateHash(updatedDomain.Password, superadmin.Password)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, isEqual, true)
	// })
	t.Run("Test case 2 | Invalid credentials", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(domain, errors.New("invalid credentials")).Once()
		user, err := usecase.UpdatePassword(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
	t.Run("Test case 3 | Server error", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()
		user, err := usecase.UpdatePassword(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
	t.Run("Test case 4 | Empty username or password", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(superadmins.Domain{}, errors.New("Username or password is empty")).Once()
		domain.Username = ""
		domain.Password = ""
		user, err := usecase.UpdatePassword(context.Background(), domain)
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, user)
	})
}
