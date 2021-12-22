package schedules

import (
	"CalFit/bussiness/schedules"

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

func (schedulesRepo *SchedulesRepo) Insert(schedulesDomain schedules.Domain) (schedules.Domain, error) {
	return schedules.Domain{}, nil
}

func (schedulesRepo *SchedulesRepo) Get(schedulesDomain schedules.Domain) ([]schedules.Domain, error) {
	return []schedules.Domain{}, nil
}

func (schedulesRepo *SchedulesRepo) Delete(schedulesDomain schedules.Domain) (schedules.Domain, error) {
	return schedules.Domain{}, nil
}
