package users

import (
	"CalFit/business/users"
	addressesRepo "CalFit/repository/mysql/addresses"
	bookingDetailsRepo "CalFit/repository/mysql/booking_details"
	"time"
)

type User struct {
	Id               int `gorm:"primaryKey"`
	Username         string
	Email            string `gorm:"not null"`
	Photo            string
	Password         string
	FullName         string
	MembershipTypeID int
	AddressID        uint `gorm:"not null"`
	BookingDetails   []bookingDetailsRepo.Booking_detail
	Address          addressesRepo.Address
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:               domain.Id,
		Username:         domain.Username,
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
		Username:         u.Username,
		Email:            u.Email,
		Photo:            u.Photo,
		Password:         u.Password,
		FullName:         u.FullName,
		MembershipTypeID: u.MembershipTypeID,
		AddressID:        u.AddressID,
		BookingDetails:   ConvertToBookingDetailsArray(u.BookingDetails),
		Address:          u.Address.ToDomain(),
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

func ConvertToBookingDetailsArray(bookingDetails []bookingDetailsRepo.Booking_detail) []users.BookingDetailDomain {
	var bookingDetailsDomain []users.BookingDetailDomain
	for _, val := range bookingDetails {
		bookingDetailsDomain = append(bookingDetailsDomain, ToBookingDetailsDomain(val))
	}
	return bookingDetailsDomain
}

func ToBookingDetailsDomain(bookingDetails bookingDetailsRepo.Booking_detail) users.BookingDetailDomain {
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

func ToListDomain(data []User) []users.Domain {
	var listDomain []users.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}

	return listDomain
}
