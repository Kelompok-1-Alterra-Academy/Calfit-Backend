package users

import (
	"CalFit/business/addresses"
	"CalFit/business/paginations"
	"context"
	"time"
)

type Domain struct {
	Id               int
	Email            string
	Photo            string
	Password         string
	MembershipTypeID int
	MembershipName   string
	AddressID        uint
	Token            string
	FullName         string
	BookingDetails   []BookingDetailDomain
	Address          addresses.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type BookingDetailDomain struct {
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

type Usecase interface {
	LoginOAuth(ctx context.Context, users Domain) (Domain, error)
	Register(ctx context.Context, users Domain) (Domain, error)
	Login(ctx context.Context, users Domain) (Domain, error)
	GetByUsername(ctx context.Context, email string) (Domain, error)
	Update(ctx context.Context, users Domain) (Domain, error)
}

type ProfileUsecase interface {
	GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
}

type Repository interface {
	LoginOAuth(ctx context.Context, users Domain) (Domain, error)
	Register(ctx context.Context, users Domain) (Domain, error)
	GetByUsername(ctx context.Context, email string) (Domain, error)
	Update(ctx context.Context, users Domain) (Domain, error)
}

type ProfileRepository interface {
	GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
}
