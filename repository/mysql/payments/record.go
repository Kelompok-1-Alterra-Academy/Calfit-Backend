package payments

import (
	bookingdetails "CalFit/repository/mysql/booking_details"
	"time"
)

type Payment struct {
	Id              int `gorm:"primaryKey"`
	Name            string
	Method          string
	Booking_details []bookingdetails.Booking_detail
	Created_at      time.Time
	Updated_at      time.Time
}
