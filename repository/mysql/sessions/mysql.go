package sessions

import (
	"CalFit/business/sessions"
	"context"

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
	if err := repo.DBConn.Debug().Create(&data); err != nil {
		return sessions.Domain{}, nil
	}
	return data.toDomain(), nil
}

func (repo *SessionsRepo) GetAll(ctx context.Context, domain sessions.Domain) ([]sessions.Domain, error) {
	return []sessions.Domain{}, nil

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
