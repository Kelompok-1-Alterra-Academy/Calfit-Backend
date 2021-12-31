package schedules

import (
	"CalFit/business/schedules"
	"CalFit/exceptions"
	"errors"

	"gorm.io/gorm"
)

type SchedulesRepo struct {
	DBConn *gorm.DB
}

func NewSchedulesRepo(db *gorm.DB) schedules.Repository {
	return &SchedulesRepo{
		DBConn: db,
	}
}

func (repo *SchedulesRepo) Insert(domain schedules.Domain) (schedules.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Create(&data).Error; err != nil {
		return schedules.Domain{}, err
	}
	return data.toDomain(), nil
}

func (repo *SchedulesRepo) Get(domain schedules.Domain) ([]schedules.Domain, error) {
	data := []Schedule{}
	if err := repo.DBConn.Debug().Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []schedules.Domain{}, exceptions.ErrNotFound
		}
		return []schedules.Domain{}, err
	}
	var domainSchedules []schedules.Domain
	for _, val := range data {
		domainSchedules = append(domainSchedules, val.toDomain())
	}
	return domainSchedules, nil
}

func (repo *SchedulesRepo) Update(domain schedules.Domain) (schedules.Domain, error) {
	data := FromDomain(domain)
	sessionId := Schedule{}
	repo.DBConn.Debug().Where("id=?", domain.Id).First(&sessionId)
	data.SessionID = sessionId.SessionID
	if err := repo.DBConn.Debug().Where("id=?", data.Id).Save(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schedules.Domain{}, exceptions.ErrNotFound
		}
		return schedules.Domain{}, err
	}
	return data.toDomain(), nil
}

func (repo *SchedulesRepo) Delete(domain schedules.Domain) (schedules.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Where("id=?", data.Id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schedules.Domain{}, exceptions.ErrNotFound
		}
	}
	repo.DBConn.Where("id=?", data.Id).Delete(&data)
	return data.toDomain(), nil
}
