package request

import "CalFit/bussiness/schedules"

type Schedules struct {
	Time_schedule string `json:"time_schedule" form:"time_schedule"`
	Duration      int    `json:"duration" form:"duration"`
	SessionID     int    `json:"session_id" form:"session_id"`
}

func ToDomain(s Schedules) schedules.Domain {
	return schedules.Domain{
		Time_schedule: s.Time_schedule,
		Duration:      s.Duration,
		SessionID:     s.SessionID,
	}
}