package users

import (
	"CalFit/business/addresses"
	bookingdetails "CalFit/business/booking_details"
	"context"
	"time"
)

type Domain struct {
	Id               int
	Email            string
	Photo            string
	Password         string
	MembershipTypeID int
	AddressID        uint
	Token            string
	BookingDetails   []bookingdetails.Domain
	Address          addresses.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	Login(ctx context.Context, users Domain) (Domain, error)
	Register(ctx context.Context, users Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, users Domain) (Domain, error)
	Register(ctx context.Context, users Domain) (Domain, error)
}
