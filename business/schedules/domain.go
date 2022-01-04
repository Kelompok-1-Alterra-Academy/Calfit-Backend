package schedules

import (
	"context"
	"time"
)

type Domain struct {
	Id           int
	TimeSchedule string
	Duration     int
	SessionID    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Insert(ctx context.Context, schedules Domain) (Domain, error)
	Get(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, schedules Domain) (Domain, error)
	Delete(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, schedules Domain) (Domain, error)
	Get(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, schedules Domain) (Domain, error)
	Delete(ctx context.Context, id int) (Domain, error)
}
