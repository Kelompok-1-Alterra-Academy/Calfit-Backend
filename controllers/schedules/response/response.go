package response

import (
	"CalFit/business/schedules"
	"time"
)

type Schedules struct {
	ID           int        `json:"id,omitempty"`
	TimeSchedule string     `json:"time_schedule,omitempty"`
	Duration     int        `json:"duration,omitempty"`
	SessionID    int        `json:"session_id,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func FromDomain(s schedules.Domain) Schedules {
	return Schedules{
		ID:           s.Id,
		TimeSchedule: s.TimeSchedule,
		Duration:     s.Duration,
		SessionID:    s.SessionID,
		CreatedAt:    &s.CreatedAt,
		UpdatedAt:    &s.UpdatedAt,
	}
}
