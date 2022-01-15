package request

import bookingdetails "CalFit/business/booking_details"

type Booking_details struct {
	Amount             int    `json:"amount" form:"amount"`
	Status             string `json:"status" form:"status"`
	UserID             int    `json:"user_id" form:"user_id"`
	OperationalAdminID int    `json:"operational_admin_id" form:"operational_admin_id"`
	PaymentID          int    `json:"payment_id" form:"payment_id"`
	ClassID            int    `json:"class_id" form:"class_id"`
}

func (request *Booking_details) ToDomain() bookingdetails.Domain {
	return bookingdetails.Domain{
		Amount:             request.Amount,
		Status:             request.Status,
		UserID:             request.UserID,
		OperationalAdminID: request.OperationalAdminID,
		PaymentID:          request.PaymentID,
		ClassID:            request.ClassID,
	}
}
