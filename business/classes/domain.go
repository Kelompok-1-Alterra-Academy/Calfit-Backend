package classes

import (
	context "context"
	"time"
)

type Domain struct {
	Id            	   uint
	Name     		   string `validate:"required"`
	Description	 	   string `validate:"required"`
	Banner_picture_url string `validate:"required"`
	Card_picture_url   string `validate:"required"`
	Category           string `validate:"required"`
	Status             string `validate:"required"`
	// Membership_typeID  uint `validate:"required"`
	GymID              uint `validate:"required"`
	// Booking_details    []booking_details.Domain
	// Schedules          []schedules.Domain `gorm:"many2many:class_schedules"`
	Created_at     		 time.Time
	Updated_at     		 time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, id string, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id string) (error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, id string, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id string) (error)
}
