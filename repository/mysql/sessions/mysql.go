package sessions

import (
	"CalFit/business/sessions"
	"context"
	"errors"

	"gorm.io/gorm"
)

type SessionsRepo struct {
	DBConn *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) sessions.Repository {
	return &SessionsRepo{
		DBConn: db,
	}
}

func (repo *SessionsRepo) Insert(ctx context.Context, domain sessions.Domain) (sessions.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Create(&data).Error; err != nil {
		return sessions.Domain{}, nil
	}
	return data.toDomain(), nil
}

func (repo *SessionsRepo) GetAll(ctx context.Context) ([]sessions.Domain, error) {
	data := []Session{}
	if err := repo.DBConn.Debug().Preload("Schedules").Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []sessions.Domain{}, errors.New("record not found")
		}
		return []sessions.Domain{}, err
	}
	var domainSession []sessions.Domain
	for _, val := range data {
		domainSession = append(domainSession, val.toDomain())
	}
	return domainSession, nil

}

func (repo *SessionsRepo) GetById(ctx context.Context, domain sessions.Domain) (sessions.Domain, error) {
	return sessions.Domain{}, nil

}

func (repo *SessionsRepo) Update(ctx context.Context, domain sessions.Domain) (sessions.Domain, error) {
	return sessions.Domain{}, nil

}

func (repo *SessionsRepo) Delete(ctx context.Context, domain sessions.Domain) (sessions.Domain, error) {

	return sessions.Domain{}, nil
}
