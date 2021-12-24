package request

import (
	"CalFit/business/schedules"
	"time"
)

type Schedules struct {
	Id            int    `json:"id" form:"id"`
	Time_schedule string `json:"time_schedule" form:"time_schedule"`
	Duration      int    `json:"duration" form:"duration"`
	SessionID     int    `json:"session_id" form:"session_id"`
}

func ToDomain(s Schedules) schedules.Domain {
	if s.Id == 0 {
		return schedules.Domain{
			Time_schedule: s.Time_schedule,
			Duration:      s.Duration,
			SessionID:     s.SessionID,
			Created_at:    time.Now(),
		}
	}
	return schedules.Domain{
		Id: s.Id,
	}
}
