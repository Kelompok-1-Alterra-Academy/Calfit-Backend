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

func (usecase *BookingDetailsUsecase) Insert(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	res, err := usecase.Repo.Insert(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
