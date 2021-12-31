package addresses

import (
	context "context"
	"time"
)

type Domain struct {
	Id            		 uint
	Address        		 string `validate:"required"`
	District        	 string `validate:"required"`
	City        		 string `validate:"required"`
	Postal_code        	 string `validate:"required"`
	Created_at     		 time.Time
	Updated_at     		 time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	// GetByISBN(ctx context.Context, isbn string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	// GetByISBN(ctx context.Context, isbn string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}
