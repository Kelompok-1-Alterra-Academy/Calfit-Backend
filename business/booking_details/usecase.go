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

func (bdu *BookingDetailsUsecase) Insert(ctx context.Context, bookingDetails Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, bdu.ContextTimeout)
	defer cancel()
	return Domain{}, nil
}
