package bookingdetails

import (
	"CalFit/business/classes"
	"CalFit/business/schedules"
	"context"
	"strconv"
	"time"
)

type BookingDetailsUsecase struct {
	Repo           Repository
	ClassRepo      classes.DomainRepository
	SchedulesRepo  schedules.Repository
	ContextTimeout time.Duration
}

func NewBookingDetailsUsecase(repo Repository, classRepo classes.DomainRepository, schedulesRepo schedules.Repository, timeout time.Duration) Usecase {
	return &BookingDetailsUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
		ClassRepo:      classRepo,
		SchedulesRepo:  schedulesRepo,
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
		classes, _ := usecase.ClassRepo.GetById(ctx, strconv.Itoa(val.ClassID))
		schedules, _ := usecase.SchedulesRepo.GetById(ctx, val.ScheduleID)
		res[i].ClassName = classes.Name
		res[i].TimeSchedule = schedules.TimeSchedule
	}
	return res, nil
}
