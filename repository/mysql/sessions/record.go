package sessions

import (
	"CalFit/repository/mysql/schedules"
	"time"
)

type Session struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Schedules   []schedules.Schedule
	Created_at  time.Time
	Updated_at  time.Time
}
