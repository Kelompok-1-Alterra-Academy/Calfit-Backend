package schedules

import (
	"CalFit/business/schedules"
	"time"
)

type Schedule struct {
	Id            int `gorm:"primaryKey"`
	Time_schedule string
	Duration      int
	SessionID     int
	Created_at    time.Time
	Updated_at    time.Time
}

func fromDomain(domain schedules.Domain) Schedule {
	return Schedule{
		Id:            domain.Id,
		Time_schedule: domain.Time_schedule,
		Duration:      domain.Duration,
		SessionID:     domain.SessionID,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}
}

func (s Schedule) toDomain() schedules.Domain {
	return schedules.Domain{
		Id:            s.Id,
		Time_schedule: s.Time_schedule,
		Duration:      s.Duration,
		SessionID:     s.SessionID,
		Created_at:    s.Created_at,
		Updated_at:    s.Updated_at,
	}
}
