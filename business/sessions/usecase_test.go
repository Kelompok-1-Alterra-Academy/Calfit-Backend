package sessions_test

import (
	"CalFit/business/schedules"
	"CalFit/business/sessions"
	"CalFit/business/sessions/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo mocks.Repository
var domain sessions.Domain
var usecase sessions.Usecase

func testSetup() {
	domain = sessions.Domain{
		Id:          1,
		Name:        "morning",
		Description: "morning",
		Schedules:   []schedules.Domain{},
	}

	usecase = sessions.NewSessionsUsecase(&repo, time.Minute*1)
}

func TestInsert(t *testing.T) {
	testSetup()
	t.Run("Valid data", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, nil).Once()
		_, err := usecase.Insert(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("Server error", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.AnythingOfType("Domain")).Return(domain, errors.New("internal server error")).Once()
		_, err := usecase.Insert(context.Background(), domain)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	testSetup()
	t.Run("Valid get", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]sessions.Domain{domain}, nil).Once()
		session, err := usecase.GetAll(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, domain, session[0])
	})
	t.Run("Server error", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]sessions.Domain{}, errors.New("Internal server error")).Once()
		session, err := usecase.GetAll(context.Background())
		assert.NotNil(t, err)
		assert.NotEqual(t, domain, session)
	})
}
