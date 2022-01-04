package request

import (
	"CalFit/business/schedules"
)

type Schedules struct {
	TimeSchedule string `json:"time_schedule" form:"time_schedule"`
	Duration     int    `json:"duration" form:"duration"`
	SessionID    int    `json:"session_id" form:"session_id"`
}

func ToDomain(s Schedules) schedules.Domain {
	return schedules.Domain{
		TimeSchedule: s.TimeSchedule,
		Duration:     s.Duration,
		SessionID:    s.SessionID,
	}
}
