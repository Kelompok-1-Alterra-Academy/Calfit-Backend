package booking_details

import "time"

type Booking_detail struct {
	Id                  int `gorm:"primaryKey"`
	Amount              int
	Status              string
	UserID              int
	Operational_adminID int
	PaymentID           int
	ClassID             int
	Created_at          time.Time
	Updated_at          time.Time
}
