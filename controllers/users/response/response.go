package request

import (
	"CalFit/business/users"
	addresses "CalFit/controllers/addresses/response"
	"time"
)

type UsersResponse struct {
	Id               int    `json:"id"`
	Email            string `json:"email"`
	Photo            string `json:"photo"`
	Password         string `json:"password"`
	MembershipTypeID int    `json:"membershipTypeId"`
	// AddressID        	uint					`json:"email"`
	Token          string                    `json:"token"`
	FullName       string                    `json:"fullname"`
	Address        addresses.AddressResponse `json:"address"`
	BookingDetails []BookingDetailsResponse
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BookingDetailsResponse struct {
	Id                 int       `json:"id"`
	Amount             int       `json:"amount"`
	Status             string    `json:"status"`
	UserID             int       `json:"userId"`
	OperationalAdminID int       `json:"operationalAdminId"`
	PaymentID          int       `json:"paymentId"`
	ClassID            int       `json:"classId"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

func FromBookingDetailsDomain(domain users.BookingDetailDomain) BookingDetailsResponse {
	return BookingDetailsResponse{
		Id:                 domain.Id,
		Amount:             domain.Amount,
		Status:             domain.Status,
		UserID:             domain.UserID,
		OperationalAdminID: domain.OperationalAdminID,
		PaymentID:          domain.PaymentID,
		ClassID:            domain.ClassID,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}

func FromBookingDetailsDomainList(domain []users.BookingDetailDomain) []BookingDetailsResponse {
	var response []BookingDetailsResponse
	for _, item := range domain {
		response = append(response, FromBookingDetailsDomain(item))
	}
	return response
}

func FromDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		Id:               domain.Id,
		Email:            domain.Email,
		Photo:            domain.Photo,
		Password:         domain.Password,
		MembershipTypeID: domain.MembershipTypeID,
		Token:            domain.Token,
		FullName:         domain.FullName,
		BookingDetails:   FromBookingDetailsDomainList(domain.BookingDetails),
		Address:          addresses.FromDomain(domain.Address),
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
