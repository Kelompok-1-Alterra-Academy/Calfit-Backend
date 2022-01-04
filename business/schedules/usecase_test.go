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
		TimeSchedule: "domain7.00-09.00",
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

func TestGet(t *testing.T) {
	testSetup()
	t.Run("Test case 1 || Valid data", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]schedules.Domain{domain}, nil).Once()
		schedules, err := usecase.Get(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, domain, schedules[0])
	})
	t.Run("Test case 2 || Server error", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]schedules.Domain{domain}, errors.New("Internal server error")).Once()
		_, err := usecase.Get(context.Background())
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	testSetup()
	t.Run("Test case 1 || Valid data", func(t *testing.T) {
		repo.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		schedules, err := usecase.GetById(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, domain, schedules)
	})
	t.Run("Test case 2 || Empty input param", func(t *testing.T) {
		repo.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(domain, errors.New("Empty input")).Once()
		_, err := usecase.GetById(context.Background(), 0)
		assert.NotNil(t, err)
	})
	t.Run("Test case 3 || Server error", func(t *testing.T) {
		repo.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(domain, errors.New("Internal server error")).Once()
		_, err := usecase.GetById(context.Background(), 1)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	testSetup()
	t.Run("Test case 1 || Valid data", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
		schedules, err := usecase.Update(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, domain, schedules)
	})
	t.Run("Test case 2 || Empty input param", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, errors.New("Empty input")).Once()
		_, err := usecase.Update(context.Background(), schedules.Domain{Id: 0})
		assert.NotNil(t, err)
	})
	t.Run("Test case 3 || Server error", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, errors.New("Internal server error")).Once()
		_, err := usecase.Update(context.Background(), domain)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	testSetup()
	t.Run("Test case 1 || Valid data", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		schedules, err := usecase.Delete(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, domain, schedules)
	})
	t.Run("Test case 2 || Empty input param", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(domain, errors.New("Empty input")).Once()
		_, err := usecase.Delete(context.Background(), 0)
		assert.NotNil(t, err)
	})
	t.Run("Test case 3 || Server error", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(domain, errors.New("Internal server error")).Once()
		_, err := usecase.Delete(context.Background(), 1)
		assert.NotNil(t, err)
	})
}
