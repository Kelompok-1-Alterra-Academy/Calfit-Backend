package memberships

import (
	"CalFit/business/addresses"
	bookingdetails "CalFit/business/booking_details"
	"CalFit/business/schedules"
	"context"
	"time"
)

type ClassDomain struct {
	Id                 uint
	Name               string
	Description        string
	Banner_picture_url string
	Card_picture_url   string
	Online             bool
	Link               string
	Category           string
	Status             string
	Price              int
	Membership_typeID  uint
	Booking_details    []bookingdetails.Domain
	Schedules          []schedules.Domain `gorm:"many2many:class_schedules"`
	Created_at         time.Time
	Updated_at         time.Time
}

type UserDomain struct {
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

type Domain struct {
	Id          int
	Name        string `validate:"required"`
	Description string `validate:"required,min=20"`
	Price       int    `validate:"required"`
	Classes     []ClassDomain
	Users       []UserDomain
	Created_at  time.Time
	Updated_at  time.Time
}

type DomainService interface {
	Insert(ctx context.Context, memberships Domain) (Domain, error)
	Get(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, id string, memberships Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}

type Repository interface {
	Insert(ctx context.Context, memberships Domain) (Domain, error)
	Get(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, id string, memberships Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}
