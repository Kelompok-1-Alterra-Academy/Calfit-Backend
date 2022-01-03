package schedules

import (
	"CalFit/business/schedules"
	"time"
)

type Schedule struct {
	Id           int `gorm:"primaryKey"`
	TimeSchedule string
	Duration     int
	SessionID    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func FromDomain(domain schedules.Domain) Schedule {
	return Schedule{
		Id:           domain.Id,
		TimeSchedule: domain.TimeSchedule,
		Duration:     domain.Duration,
		SessionID:    domain.SessionID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func (s Schedule) toDomain() schedules.Domain {
	return schedules.Domain{
		Id:           s.Id,
		TimeSchedule: s.TimeSchedule,
		Duration:     s.Duration,
		SessionID:    s.SessionID,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}
