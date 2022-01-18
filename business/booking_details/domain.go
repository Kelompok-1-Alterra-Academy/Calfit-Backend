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
	ClassName          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Usecase interface {
	Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
}
