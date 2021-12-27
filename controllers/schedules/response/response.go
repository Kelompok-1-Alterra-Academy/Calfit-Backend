package response

import (
	"CalFit/business/schedules"
)

type Schedules struct {
	Time_schedule string `json:"time_schedule"`
	Duration      int    `json:"duration"`
}

func FromDomain(s schedules.Domain) Schedules {
	return Schedules{
		Time_schedule: s.Time_schedule,
		Duration:      s.Duration,
	}
}
