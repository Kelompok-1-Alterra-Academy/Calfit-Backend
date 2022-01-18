package bookingdetails

import (
	"CalFit/business/classes"
	"context"
	"strconv"
	"time"
)

type BookingDetailsUsecase struct {
	Repo           Repository
	ClassRepo      classes.DomainRepository
	ContextTimeout time.Duration
}

func NewBookingDetailsUsecase(repo Repository, classRepo classes.DomainRepository, timeout time.Duration) Usecase {
	return &BookingDetailsUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
		ClassRepo:      classRepo,
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
	for i, val := range res {
		domain, _ := usecase.ClassRepo.GetById(ctx, strconv.Itoa(val.ClassID))
		res[i].ClassName = domain.Name
	}
	return res, nil
}
