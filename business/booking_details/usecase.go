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

func (usecase *BookingDetailsUsecase) GetByUserID(ctx context.Context, userID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	res, err := usecase.Repo.GetByUserID(ctx, userID)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}

func (usecase *BookingDetailsUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	res, err := usecase.Repo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (usecase *BookingDetailsUsecase) GetAll(ctx context.Context, id int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	res, err := usecase.Repo.GetAll(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}
