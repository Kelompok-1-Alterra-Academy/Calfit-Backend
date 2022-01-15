package bookingdetails

import (
	"context"
	"time"
)

type Domain struct {
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
	Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
}
