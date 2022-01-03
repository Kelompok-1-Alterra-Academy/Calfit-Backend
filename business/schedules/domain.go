package schedules

import (
	"time"
)

type Domain struct {
	Id           int
	TimeSchedule string
	Duration     int
	SessionID    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Insert(schedules Domain) (Domain, error)
	Get(schedules Domain) ([]Domain, error)
	Update(schedules Domain) (Domain, error)
	Delete(schedules Domain) (Domain, error)
}

type Repository interface {
	Insert(schedules Domain) (Domain, error)
	Get(schedules Domain) ([]Domain, error)
	Update(schedules Domain) (Domain, error)
	Delete(schedules Domain) (Domain, error)
}
