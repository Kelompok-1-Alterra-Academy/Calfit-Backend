package classes_test

import (
	"CalFit/business/classes"
	"CalFit/business/classes/mocks"
	gymMocks "CalFit/business/gyms/mocks"
	"CalFit/business/paginations"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var classRepository mocks.DomainRepository
var gymRepository gymMocks.DomainRepository

var classService classes.DomainService
var classDomain, updatedClassDomain, emptyClassDomain classes.Domain
var paginationDomain paginations.Domain

func setup() {
	classService = classes.NewUsecase(&classRepository, &gymRepository, time.Minute*15)
	classDomain = classes.Domain{
		Id:                 1,
		Name:               "Muaythai class",
		Description:        "Lets kick the others ass",
		Banner_picture_url: "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAS9iRN.img?h=531&w=799&m=6&q=60&o=f&l=f&x=246&y=140",
		Card_picture_url:   "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAS9iRN.img?h=531&w=799&m=6&q=60&o=f&l=f&x=246&y=140",
		Category:           "Martial arts",
		Status:             "Active",
		// Membership_typeID: 1,
		GymID: 1,
	}

	updatedClassDomain = classes.Domain{
		Id:                 1,
		Name:               "Yoga class",
		Description:        "Find the peace here",
		Banner_picture_url: "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAS9iRN.img?h=531&w=799&m=6&q=60&o=f&l=f&x=246&y=140",
		Card_picture_url:   "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAS9iRN.img?h=531&w=799&m=6&q=60&o=f&l=f&x=246&y=140",
		Category:           "Yoga",
		Status:             "Inactive",
		// Membership_typeID: 1,
		GymID: 1,
	}

	emptyClassDomain = classes.Domain{
		Id:                 0,
		Name:               "",
		Description:        "",
		Banner_picture_url: "",
		Card_picture_url:   "",
		Category:           "",
		Status:             "",
		// Membership_typeID: 0,
		GymID: 0,
	}

	paginationDomain = paginations.Domain{
		Page:  1,
		Limit: 10,
		Sort:  "asc",
	}
}

func TestGetAllClasses(t *testing.T) {
	setup()
	classRepository.On("GetAll", mock.Anything, mock.AnythingOfType("paginations.Domain")).Return([]classes.Domain{classDomain}, nil)
	t.Run("Test Case 1 | Get All Classes", func(t *testing.T) {
		classes, err := classService.GetAll(context.Background(), paginationDomain)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(classes) == 0 {
			t.Errorf("Error: %s", "No classes found")
		}
		assert.Nil(t, err)
		assert.Equal(t, classDomain, classes[0])
	})
}

func TestGetClassByClassId(t *testing.T) {
	setup()
	classRepository.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(classDomain, nil)
	t.Run("Test Case 1 | Valid Get Class By ClassId", func(t *testing.T) {
		class, err := classService.GetById(context.Background(), "1")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, classDomain, class)
	})
	t.Run("Test Case 2 | Invalid Get Class By Empty ClassId", func(t *testing.T) {
		class, err := classService.GetById(context.Background(), "")
		assert.NotNil(t, err)
		assert.NotEqual(t, class, classDomain)
	})
}

func TestCreateNewClass(t *testing.T) {
	setup()
	classRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain"), mock.AnythingOfType("string")).Return(classDomain, nil)
	// t.Run("Test Case 1 | Valid Create New Class", func(t *testing.T) {
	// 	class, err := classService.Create(context.Background(), classDomain, fmt.Sprintf("%d", classDomain.GymID))
	// 	if err != nil {
	// 		t.Errorf("Error: %s", err)
	// 	}
	// 	if class.Id == 0 {
	// 		t.Errorf("Error: %s", "class Id is empty")
	// 	}
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, classDomain, class)
	// })
	// t.Run("Test Case 2 | Invalid Create New Class with Wrong gymId", func(t *testing.T) {
	// 	class, err := classService.Create(context.Background(), classDomain, strconv.Itoa(int(emptyClassDomain.GymID)))
	// 	assert.NotNil(t, err)
	// 	assert.NotEqual(t, class, classDomain)
	// })
	// t.Run("Test Case 3 | Invalid Create New Class with Empty Fields", func(t *testing.T) {
	// 	class, err := classService.Create(context.Background(), emptyClassDomain, strconv.Itoa(int(classDomain.GymID)))
	// 	assert.NotNil(t, err)
	// 	assert.NotEqual(t, class, classDomain)
	// })
}

func TestUpdateClass(t *testing.T) {
	setup()
	classRepository.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("Domain")).Return(classDomain, nil)
	// t.Run("Test Case 1 | Valid Update Class by Id", func(t *testing.T) {
	// 	class, err := classService.Update(context.Background(), "1", updatedClassDomain)
	// 	t.Log(class)
	// 	if err != nil {
	// 		t.Errorf("Error: %s", err)
	// 	}
	// 	if class.Id == 0 {
	// 		t.Errorf("Error: %s", "class Id is empty")
	// 	}
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, updatedClassDomain, class)
	// })
	t.Run("Test Case 2 | Invalid Update Class by Empty Id", func(t *testing.T) {
		class, err := classService.Update(context.Background(), "", updatedClassDomain)
		assert.NotNil(t, err)
		assert.NotEqual(t, class, updatedClassDomain)
	})
	t.Run("Test Case 3 | Invalid Update Class by Empty Fields", func(t *testing.T) {
		class, err := classService.Update(context.Background(), fmt.Sprintf("%d", updatedClassDomain.Id), emptyClassDomain)
		assert.NotNil(t, err)
		assert.NotEqual(t, class, updatedClassDomain)
	})
}

func TestDeleteClass(t *testing.T) {
	setup()
	classRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil)
	t.Run("Test Case 1 | Valid Delete Class", func(t *testing.T) {
		err := classService.Delete(context.Background(), "1")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid Delete Class with Empty Id", func(t *testing.T) {
		err := classService.Delete(context.Background(), "")
		assert.NotNil(t, err)
		assert.NotEqual(t, err, nil)
	})
}
