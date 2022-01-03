package bookingdetails

import "time"

type Domain struct {
	Id                 int
	Amount             int
	Status             string
	UserID             int
	OperationalAdminID int
	PaymentID          int
	ClassID            int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
