package classes

import (
	"CalFit/repository/mysql/booking_details"
	"CalFit/repository/mysql/schedules"
	"time"
)

type Class struct {
	Id                int `gorm:"primaryKey"`
	Name              string
	Description       string
	Picture_url       string
	Category          string
	Status            string
	Membership_typeID int
	GymID             int
	Booking_details   []booking_details.Booking_detail
	Schedules         []schedules.Schedule `gorm:"many2many:class_schedules"`
	Created_at        time.Time
	Updated_at        time.Time
}
