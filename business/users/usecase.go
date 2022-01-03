package users

import (
	"CalFit/exceptions"
	"context"
	"time"
)

type UsersUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewUsersUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UsersUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (uu *UsersUsecase) Login(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := uu.Repo.Login(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
