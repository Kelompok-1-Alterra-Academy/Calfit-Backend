package sessions

import (
	"context"
	"time"
)

type SessionsUsecase struct {
	sessionsRepo   Repository
	contextTimeout time.Duration
}

func NewSessionsUsecase(repo Repository, timeout time.Duration) Usecase {
	return &SessionsUsecase{
		sessionsRepo: repo,
	}
}

func (su *SessionsUsecase) Insert(ctx context.Context, sessions Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	res, err := su.sessionsRepo.Insert(ctx, sessions)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (su *SessionsUsecase) GetById(ctx context.Context, sessions Domain) (Domain, error) {
	return Domain{}, nil
}

func (su *SessionsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	res, err := su.sessionsRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}

func (su *SessionsUsecase) Update(ctx context.Context, sessions Domain) (Domain, error) {
	return Domain{}, nil
}

func (su *SessionsUsecase) Delete(ctx context.Context, sessions Domain) (Domain, error) {
	return Domain{}, nil
}
