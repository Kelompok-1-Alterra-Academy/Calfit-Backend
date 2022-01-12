package bookingdetails

import "time"

type Booking_detail struct {
	Id                 int `gorm:"primaryKey"`
	Amount             int
	Status             string
	UserID             int
	OperationalAdminID int
	PaymentID          int
	ClassID            int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
