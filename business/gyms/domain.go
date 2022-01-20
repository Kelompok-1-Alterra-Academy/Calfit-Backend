package gyms

import (
	"CalFit/business/addresses"
	"CalFit/business/paginations"
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
	Membership_typeID  uint
	// Booking_details    []booking_details.Domain
	// Schedules          []schedules.Domain `gorm:"many2many:class_schedules"`
	Created_at time.Time
	Updated_at time.Time
}

type Domain struct {
	Id                   uint
	Name                 string `validate:"required"`
	Description          string `validate:"required"`
	Telephone            string `validate:"required,min=7,max=20,numeric"`
	Picture              string `validate:"required"`
	Operational_admin_ID uint   `validate:"required"`
	Address_ID           uint
	Operational_admin    uint
	Address              addresses.Domain
	Classes              []ClassDomain
	Created_at           time.Time
	Updated_at           time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}

type DomainService interface {
	GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}
