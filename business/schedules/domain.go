package schedules

import (
	"time"
)

type Domain struct {
	Id            int
	Time_schedule string
	Duration      int
	SessionID     int
	Created_at    time.Time
	Updated_at    time.Time
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
