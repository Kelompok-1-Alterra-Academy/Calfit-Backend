package response

import (
	bookingdetails "CalFit/business/booking_details"
	"time"
)

type Booking_details struct {
	Amount             int       `json:"amount" form:"amount"`
	Status             string    `json:"status" form:"status"`
	UserID             int       `json:"user_id" form:"user_id"`
	OperationalAdminID int       `json:"operational_admin_id" form:"operational_admin_id"`
	PaymentID          int       `json:"payment_id" form:"payment_id"`
	ClassID            int       `json:"class_id" form:"class_id"`
	ClassName          string    `json:"class_name" form:"class_name"`
	CreatedAt          time.Time `json:"created_at" form:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" form:"updated_at"`
}

func FromDomain(domain bookingdetails.Domain) Booking_details {
	return Booking_details{
		Amount:             domain.Amount,
		Status:             domain.Status,
		UserID:             domain.UserID,
		OperationalAdminID: domain.OperationalAdminID,
		PaymentID:          domain.PaymentID,
		ClassID:            domain.ClassID,
		ClassName:          domain.ClassName,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}
