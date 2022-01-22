package users

import (
	"CalFit/business/users"
	"CalFit/repository/mysql/addresses"
	bookingDetailsRepo "CalFit/repository/mysql/booking_details"
	"time"
)

type User struct {
	Id               int    `gorm:"primaryKey"`
	Email            string `gorm:"not null"`
	Photo            string
	Password         string
	FullName         string
	MembershipTypeID int
	AddressID        uint `gorm:"not null"`
	BookingDetails   []bookingDetailsRepo.Booking_detail
	Address          addresses.Address `gorm:"foreignkey:AddressID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:               domain.Id,
		Email:            domain.Email,
		Photo:            domain.Photo,
		Password:         domain.Password,
		FullName:         domain.FullName,
		MembershipTypeID: domain.MembershipTypeID,
		AddressID:        domain.AddressID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func (u User) ToDomain() users.Domain {
	return users.Domain{
		Id:               u.Id,
		Email:            u.Email,
		Photo:            u.Photo,
		Password:         u.Password,
		FullName:         u.FullName,
		MembershipTypeID: u.MembershipTypeID,
		AddressID:        u.AddressID,
		BookingDetails:   convertToArray(u.BookingDetails),
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

func convertToArray(bookingDetails []bookingDetailsRepo.Booking_detail) []users.BookingDetailDomain {
	var bookingDetailsDomain []users.BookingDetailDomain
	for _, val := range bookingDetails {
		bookingDetailsDomain = append(bookingDetailsDomain, toBookingDetailsDomain(val))
	}
	return bookingDetailsDomain
}

func toBookingDetailsDomain(bookingDetails bookingDetailsRepo.Booking_detail) users.BookingDetailDomain {
	return users.BookingDetailDomain{
		Id:                 bookingDetails.Id,
		Amount:             bookingDetails.Amount,
		Status:             bookingDetails.Status,
		UserID:             bookingDetails.UserID,
		OperationalAdminID: bookingDetails.OperationalAdminID,
		PaymentID:          bookingDetails.PaymentID,
		ClassID:            bookingDetails.ClassID,
		CreatedAt:          bookingDetails.CreatedAt,
		UpdatedAt:          bookingDetails.UpdatedAt,
	}
}
