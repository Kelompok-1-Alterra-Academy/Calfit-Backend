package payments

import (
	"CalFit/repository/mysql/booking_details"
	"time"
)

type Payment struct {
	Id              int `gorm:"primaryKey"`
	Name            string
	Method          string
	Booking_details []booking_details.Booking_detail
	Created_at      time.Time
	Updated_at      time.Time
}
