package bookingdetails

import (
	"context"
	"time"
)

type BookingDetailsUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewBookingDetailsUsecase(repo Repository, timeout time.Duration) Usecase {
	return &BookingDetailsUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (usecase *BookingDetailsUsecase) Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
