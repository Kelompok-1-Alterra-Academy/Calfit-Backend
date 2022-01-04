package schedules_test

import (
	"CalFit/business/schedules"
	"CalFit/business/schedules/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo mocks.Repository
var domain schedules.Domain
var usecase schedules.Usecase

func testSetup() {
	domain = schedules.Domain{
		Id:           1,
		TimeSchedule: "07.00-09.00",
		Duration:     120,
		SessionID:    1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	usecase = schedules.NewSchedulesUsecase(&repo, time.Minute*1)
}

func TestInsert(t *testing.T) {
	testSetup()
	t.Run("Test case 1 || Valid data", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
		schedules, err := usecase.Insert(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, domain, schedules)
	})
	t.Run("Test case 2 || Server error", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, errors.New("Internal server error")).Once()
		_, err := usecase.Insert(context.Background(), domain)
		assert.NotNil(t, err)
	})
}
