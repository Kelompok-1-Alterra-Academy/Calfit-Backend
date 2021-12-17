package schedules

import (
	"time"
)

type Schedule struct {
	Id            int `gorm:"primaryKey"`
	Time_schedule string
	Duration      int
	SessionID     int
	Created_at    time.Time
	Updated_at    time.Time
}

type Class struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Picture_url string
	Category    string
	Status      string
	Schedules   []Schedule `gorm:"many2many:class_schedules"`
	Created_at  time.Time
	Updated_at  time.Time
}
