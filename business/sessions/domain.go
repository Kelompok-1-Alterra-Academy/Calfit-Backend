package sessions

import (
	"CalFit/business/schedules"
	"context"
	"time"
)

type Domain struct {
	Id          int
	Name        string
	Description string
	Schedules   []schedules.Domain
	Created_at  time.Time
	Updated_at  time.Time
}

type Usecase interface {
	Insert(ctx context.Context, sessions Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, sessions Domain) (Domain, error)
	Update(ctx context.Context, sessions Domain) (Domain, error)
	Delete(ctx context.Context, sessions Domain) (Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, sessions Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, sessions Domain) (Domain, error)
	Update(ctx context.Context, sessions Domain) (Domain, error)
	Delete(ctx context.Context, sessions Domain) (Domain, error)
}
