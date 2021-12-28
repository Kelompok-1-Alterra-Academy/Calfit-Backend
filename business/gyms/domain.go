package gyms

import (
	"CalFit/business/addresses"
	context "context"
	"time"
)

type Domain struct {
	Id            		 uint
	Name     			 string `validate:"required"`
	Telephone     		 string `validate:"required,min=7,max=20,numeric"`
	Picture     		 string `validate:"required"`
	Operational_admin_ID uint `validate:"required"`
	Address_ID        	 uint
	Operational_admin    uint
	Address        		 addresses.Domain
	Created_at     		 time.Time
	Updated_at     		 time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) (error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) (error)
}
