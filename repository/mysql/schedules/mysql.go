package schedules

import (
	"CalFit/business/schedules"
	"CalFit/exceptions"
	"context"
	"errors"
	"time"

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

func (repo *SchedulesRepo) Insert(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	data := FromDomain(domain)
	data.CreatedAt = time.Now()
	if err := repo.DBConn.Create(&data).Error; err != nil {
		return schedules.Domain{}, err
	}
	return data.toDomain(), nil
}

func (repo *SchedulesRepo) Get(ctx context.Context) ([]schedules.Domain, error) {
	data := []Schedule{}
	if err := repo.DBConn.Find(&data).Error; err != nil {
		return []schedules.Domain{}, err
	}
	var domainSchedules []schedules.Domain
	for _, val := range data {
		domainSchedules = append(domainSchedules, val.toDomain())
	}
	return domainSchedules, nil
}

func (repo *SchedulesRepo) GetById(ctx context.Context, id int) (schedules.Domain, error) {
	data := Schedule{}
	if err := repo.DBConn.Where("id=?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schedules.Domain{}, exceptions.ErrNotFound
		}
		return schedules.Domain{}, err
	}
	return data.toDomain(), nil
}

func (repo *SchedulesRepo) Update(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("id=?", data.Id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schedules.Domain{}, exceptions.ErrNotFound
		}
		return schedules.Domain{}, err
	}
	data.Duration = domain.Duration
	data.TimeSchedule = domain.TimeSchedule
	data.SessionID = domain.SessionID
	data.UpdatedAt = time.Now()
	if err := repo.DBConn.Where("id=?", data.Id).Save(&data).Error; err != nil {
		return schedules.Domain{}, err
	}
	return data.toDomain(), nil
}

func (repo *SchedulesRepo) Delete(ctx context.Context, id int) (schedules.Domain, error) {
	data := Schedule{}
	if err := repo.DBConn.Where("id=?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schedules.Domain{}, exceptions.ErrNotFound
		}
		return schedules.Domain{}, err
	}
	repo.DBConn.Where("id=?", data.Id).Delete(&data)
	return data.toDomain(), nil
}
