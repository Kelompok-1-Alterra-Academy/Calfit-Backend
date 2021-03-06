package response

import (
	bookingdetails "CalFit/business/booking_details"
	classResponse "CalFit/controllers/classes/response"
	"CalFit/controllers/schedules/response"
	"time"
)

type Booking_details struct {
	ID                 int                         `json:"id"`
	Amount             int                         `json:"amount"`
	Status             string                      `json:"status"`
	UserID             int                         `json:"user_id"`
	OperationalAdminID int                         `json:"operational_admin_id"`
	PaymentID          int                         `json:"payment_id"`
	Class              classResponse.ClassResponse `json:"class"`
	PaymentProof       string                      `json:"payment_proof,omitempty"`
	CreatedAt          *time.Time                  `json:"created_at"`
	UpdatedAt          *time.Time                  `json:"updated_at"`
}

func FromDomain(domain bookingdetails.Domain) Booking_details {
	class := classResponse.ClassResponse{
		ID:   uint(domain.Id),
		Name: domain.ClassName,
		Schedules: []response.Schedules{
			{
				TimeSchedule: domain.TimeSchedule,
			},
		},
		GymName:          domain.GymName,
		Card_picture_url: domain.CardPictureUrl,
		Online:           domain.Online,
		Link:             domain.Link,
	}
	return Booking_details{
		ID:                 domain.Id,
		Amount:             domain.Amount,
		Status:             domain.Status,
		UserID:             domain.UserID,
		OperationalAdminID: domain.OperationalAdminID,
		PaymentID:          domain.PaymentID,
		Class:              class,
		PaymentProof:       domain.PaymentProof,
		CreatedAt:          &domain.CreatedAt,
		UpdatedAt:          &domain.UpdatedAt,
	}
}
