package gyms_test

import (
	"CalFit/business/gyms"
	"CalFit/business/gyms/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var gymRepository mocks.DomainRepository

var gymService gyms.DomainService
var gymDomain gyms.Domain

func setup() {
	gymService = gyms.NewUsecase(&gymRepository, time.Minute*15)
	gymDomain = gyms.Domain{
		Id:          		  1,
		Name: 				  "Gelud Gym",
		Telephone:   		  "08123456789",
		Picture:	     	  "",
		Operational_admin_ID: 1,
		Address_ID: 		  1,
	}
}

func TestGetAllGyms(t *testing.T) {
	setup()
	gymRepository.On("GetAll", mock.Anything).Return([]gyms.Domain{gymDomain}, nil)
	t.Run("Test Case 1 | Get All Gyms", func(t *testing.T) {
		gyms, err := gymService.GetAll(context.Background())
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(gyms) == 0 {
			t.Errorf("Error: %s", "No gyms found")
		}
		assert.Nil(t, err)
		assert.Equal(t, gymDomain, gyms[0])
	})

}

func TestGetGymByGymId(t *testing.T) {
	setup()
	gymRepository.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(gymDomain, nil)
	t.Run("Test Case 1 | Valid Get Gym By GymId", func(t *testing.T) {
		gym, err := gymService.GetById(context.Background(), "5wBQEp6ruIAC")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, gymDomain, gym)
	})
	t.Run("Test Case 2 | Invalid Get Gym By Empty GymId", func(t *testing.T) {
		gym, err := gymService.GetById(context.Background(), "")
		assert.NotNil(t, err)
		assert.NotEqual(t, gym, gymDomain)
	})
}

func TestCreateNewGym(t *testing.T) {
	setup()
	gymRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(gymDomain, nil)
	t.Run("Test Case 1 | Valid Create New gym", func(t *testing.T) {
		gym, err := gymService.Create(context.Background(), gymDomain)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if gym.Id == 0 {
			t.Errorf("Error: %s", "gym Id is empty")
		}
		assert.Nil(t, err)
		assert.Equal(t, gymDomain, gym)
	})
	t.Run("Test Case 2 | Invalid Create New gym with Empty Fields", func(t *testing.T) {
		gym, err := gymService.Create(context.Background(), gyms.Domain{
			Name: 				  "",
			Telephone:   		  "",
			Picture:	     	  "",
			Operational_admin_ID: 0,
			Address_ID: 		  0,
		})
		assert.NotNil(t, err)
		assert.NotEqual(t, gym, gymDomain)
	})
}

// func TestUpdategymstatusByGymId(t *testing.T) {
// 	setup()
// 	gymRepository.On("UpdateStatus", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("bool")).Return(gymDomain, nil)
// 	t.Run("Test Case 1 | Valid Update gym Status", func(t *testing.T) {
// 		gym, err := gymService.UpdateStatus(context.Background(), "5wBQEp6ruIAC", true)
// 		if err != nil {
// 			t.Errorf("Error: %s", err)
// 		}
// 		assert.Nil(t, err)
// 		assert.Equal(t, gymDomain, gym)
// 	})
// 	t.Run("Test Case 2 | Invalid Update gym Status with Empty GymId", func(t *testing.T) {
// 		gym, err := gymService.UpdateStatus(context.Background(), "", true)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, gym, gymDomain)
// 	})
// }
