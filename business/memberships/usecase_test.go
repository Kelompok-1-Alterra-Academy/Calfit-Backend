package memberships_test

import (
	"CalFit/business/memberships"
	"CalFit/business/memberships/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var membershipRepository mocks.Repository

var membershipService memberships.DomainService
var membershipDomain, updatedMembershipDomain, emptyMembershipDomain memberships.Domain

func setup() {
	membershipService = memberships.NewMembershipsUsecase(&membershipRepository, time.Minute*15)
	membershipDomain = memberships.Domain{
		Id:          1,
		Name:        "Basic",
		Description: "Get basic membership for free to all member",
	}
	membershipDomain = memberships.Domain{
		Id:          1,
		Name:        "Basic",
		Description: "Get basic membership for free to all member",
	}
	emptyMembershipDomain = memberships.Domain{
		Id:          1,
		Name:        "Basic",
		Description: "Get basic membership for free to all member",
	}
}

func TestGetMemberships(t *testing.T) {
	setup()
	membershipRepository.On("Get", mock.Anything, mock.AnythingOfType("memberships.Domain")).Return([]memberships.Domain{membershipDomain}, nil)
	t.Run("Test Case 1 | Get All Memberships", func(t *testing.T) {
		memberships, err := membershipService.Get(context.Background())
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(memberships) == 0 {
			t.Errorf("Error: %s", "No Memberships found")
		}
		assert.Nil(t, err)
		assert.Equal(t, membershipDomain, memberships[0])
	})
}

func TestGetMembershipByMembershipId(t *testing.T) {
	setup()
	membershipRepository.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(membershipDomain, nil)
	t.Run("Test Case 1 | Valid Get Membership By MembershipId", func(t *testing.T) {
		gym, err := membershipService.GetById(context.Background(), "1")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, membershipDomain, gym)
	})
	t.Run("Test Case 2 | Invalid Get Membership By Empty MembershipId", func(t *testing.T) {
		gym, err := membershipService.GetById(context.Background(), "")
		assert.NotNil(t, err)
		assert.NotEqual(t, gym, membershipDomain)
	})
}

func TestCreateNewMembership(t *testing.T) {
	setup()
	membershipRepository.On("Insert", mock.Anything, mock.AnythingOfType("Domain")).Return(membershipDomain, nil)
	t.Run("Test Case 1 | Valid Insert New Membership", func(t *testing.T) {
		gym, err := membershipService.Insert(context.Background(), membershipDomain)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if gym.Id == 0 {
			t.Errorf("Error: %s", "gym Id is empty")
		}
		assert.Nil(t, err)
		assert.Equal(t, membershipDomain, gym)
	})
	t.Run("Test Case 2 | Invalid Insert New Membership with Empty Fields", func(t *testing.T) {
		gym, err := membershipService.Insert(context.Background(), emptyMembershipDomain)
		assert.NotNil(t, err)
		assert.NotEqual(t, gym, membershipDomain)
	})
}

func TestUpdateMembershipByMembershipId(t *testing.T) {
	setup()
	membershipRepository.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("Domain")).Return(updatedMembershipDomain, nil)

	t.Run("Test Case 1 | Valid Update Membership", func(t *testing.T) {
		gym, err := membershipService.Update(context.Background(), "1", updatedMembershipDomain)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, updatedMembershipDomain, gym)
	})
	t.Run("Test Case 2 | Invalid Update Membership with Empty MembershipId", func(t *testing.T) {
		gym, err := membershipService.Update(context.Background(), "0", emptyMembershipDomain)
		assert.NotNil(t, err)
		assert.NotEqual(t, gym, membershipDomain)
	})
}

func TestDeleteMembershipByMembershipId(t *testing.T) {
	setup()
	membershipRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil)
	t.Run("Test Case 1 | Valid Delete Order", func(t *testing.T) {
		err := membershipService.Delete(context.Background(), "1")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid Delete Membership with Empty MembershipId", func(t *testing.T) {
		err := membershipService.Delete(context.Background(), "")
		if err.Error() != "empty input" {
			t.Errorf("Error: %s", err)
		}
		assert.NotNil(t, err)
	})
}