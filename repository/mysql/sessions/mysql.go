package sessions

import (
	"CalFit/business/sessions"
	"CalFit/exceptions"
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
	if err := repo.DBConn.Create(&data).Error; err != nil {
		return sessions.Domain{}, nil
	}
	return data.ToDomain(), nil
}

func (repo *SessionsRepo) GetAll(ctx context.Context) ([]sessions.Domain, error) {
	data := []Session{}
	if err := repo.DBConn.Preload("Schedules").Find(&data).Error; err != nil {
		return []sessions.Domain{}, err
	}
	var domainSession []sessions.Domain
	for _, val := range data {
		domainSession = append(domainSession, val.ToDomain())
	}
	return domainSession, nil

}

func (repo *SessionsRepo) GetById(ctx context.Context, id int) (sessions.Domain, error) {
	data := Session{}
	if err := repo.DBConn.Where("id=?", id).Preload("Schedules").First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sessions.Domain{}, exceptions.ErrNotFound
		}
		return sessions.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *SessionsRepo) Update(ctx context.Context, domain sessions.Domain) (sessions.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("id=?", data.Id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sessions.Domain{}, exceptions.ErrNotFound
		}
		return sessions.Domain{}, err
	}
	if err := repo.DBConn.Save(&data).Error; err != nil {
		return sessions.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *SessionsRepo) Delete(ctx context.Context, id int) (sessions.Domain, error) {
	data := Session{}
	if err := repo.DBConn.Where("id=?", id).Preload("Schedules").First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sessions.Domain{}, exceptions.ErrNotFound
		}
		return sessions.Domain{}, err
	}
	if err := repo.DBConn.Delete(&data).Error; err != nil {
		return sessions.Domain{}, err
	}
	return data.ToDomain(), nil
}
