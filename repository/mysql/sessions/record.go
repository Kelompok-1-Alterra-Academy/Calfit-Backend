package sessions

import (
	"CalFit/business/schedules"
	"CalFit/business/sessions"
	schedulesRepo "CalFit/repository/mysql/schedules"
	"time"
)

type Session struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Schedules   []schedulesRepo.Schedule
	Created_at  time.Time
	Updated_at  time.Time
}

func FromDomain(domain sessions.Domain) Session {
	return Session{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
}

func (s Session) toDomain() sessions.Domain {
	return sessions.Domain{
		Id:          s.Id,
		Name:        s.Name,
		Description: s.Description,
		Schedules:   convertToArray(s.Schedules),
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
}

func convertToArray(schedule []schedulesRepo.Schedule) []schedules.Domain {
	schedulesDomain := []schedules.Domain{}
	for _, val := range schedule {
		schedulesDomain = append(schedulesDomain, toScheduleDomain(val))
	}
	return schedulesDomain
}

func toScheduleDomain(schedule schedulesRepo.Schedule) schedules.Domain {
	return schedules.Domain{
		Id:            schedule.Id,
		Time_schedule: schedule.Time_schedule,
		Duration:      schedule.Duration,
		SessionID:     schedule.SessionID,
		Created_at:    schedule.Created_at,
		Updated_at:    schedule.Updated_at,
	}
}
