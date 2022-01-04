package schedules

import (
	"context"
	"time"
)

type SchedulesUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewSchedulesUsecase(repo Repository, timeout time.Duration) Usecase {
	return &SchedulesUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (schedulesUseCase *SchedulesUsecase) Insert(ctx context.Context, schedules Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, schedulesUseCase.ContextTimeout)
	defer cancel()
	res, err := schedulesUseCase.Repo.Insert(ctx, schedules)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (schedulesUseCase *SchedulesUsecase) Get(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, schedulesUseCase.ContextTimeout)
	defer cancel()
	res, err := schedulesUseCase.Repo.Get(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}

func (schedulesUseCase *SchedulesUsecase) Update(ctx context.Context, schedules Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, schedulesUseCase.ContextTimeout)
	defer cancel()
	res, err := schedulesUseCase.Repo.Update(ctx, schedules)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (schedulesUseCase *SchedulesUsecase) Delete(ctx context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, schedulesUseCase.ContextTimeout)
	defer cancel()
	res, err := schedulesUseCase.Repo.Delete(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
