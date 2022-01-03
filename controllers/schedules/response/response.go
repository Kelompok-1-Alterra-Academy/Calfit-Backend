package response

import (
	"CalFit/business/schedules"
)

type Schedules struct {
	TimeSchedule string `json:"time_schedule"`
	Duration     int    `json:"duration"`
}

func FromDomain(s schedules.Domain) Schedules {
	return Schedules{
		TimeSchedule: s.TimeSchedule,
		Duration:     s.Duration,
	}
}
