package gyms

import (
	context "context"
	"time"
)

type Domain struct {
	Id            uint
	Name     string `validate:"required"`
	Telephone     string `validate:"required"`
	Picture     string `validate:"required"`
	Operational_admin_ID        string `validate:"required"`
	Address_ID        string `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// GetById(ctx context.Context, id string) (Domain, error)
	// GetByISBN(ctx context.Context, isbn string) (Domain, error)
	// Create(ctx context.Context, domain Domain) (Domain, error)
	// UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// GetById(ctx context.Context, id string) (Domain, error)
	// GetByISBN(ctx context.Context, isbn string) (Domain, error)
	// Create(ctx context.Context, domain Domain) (Domain, error)
	// UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}
