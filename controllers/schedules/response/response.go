package response

import (
	"CalFit/business/schedules"
	"time"
)

type Schedules struct {
	ID           int       `json:"id"`
	TimeSchedule string    `json:"time_schedule"`
	Duration     int       `json:"duration"`
	SessionID    int       `json:"session_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(s schedules.Domain) Schedules {
	return Schedules{
		ID:           s.Id,
		TimeSchedule: s.TimeSchedule,
		Duration:     s.Duration,
		SessionID:    s.SessionID,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}
