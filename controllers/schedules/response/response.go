package response

import (
	"CalFit/business/schedules"
	"time"
)

type Schedules struct {
	ID           int       `json:"id"`
	TimeSchedule string    `json:"time_schedule"`
	Duration     int       `json:"duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(s schedules.Domain) Schedules {
	return Schedules{
		ID:           s.Id,
		TimeSchedule: s.TimeSchedule,
		Duration:     s.Duration,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}
