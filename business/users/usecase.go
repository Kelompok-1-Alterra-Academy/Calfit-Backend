package users

import (
	"CalFit/app/middlewares"
	"CalFit/exceptions"
	"CalFit/helpers"
	"context"
	"time"
)

type UsersUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JWTAuth        *middlewares.ConfigJWT
}

func NewUsersUsecase(repo Repository, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UsersUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
		JWTAuth:        jwtAuth,
	}
}

func (uu *UsersUsecase) LoginOAuth(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := uu.Repo.LoginOAuth(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	res.Token, _ = uu.JWTAuth.GenerateToken(res.Email)
	return res, nil
}

func (uu *UsersUsecase) Register(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	var err error
	users.Password, err = helpers.Hash(users.Password)
	if err != nil {
		return Domain{}, err
	}
	res, err := uu.Repo.Register(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (uu *UsersUsecase) Login(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := uu.Repo.GetByUsername(ctx, users.Email)
	if err != nil {
		return Domain{}, err
	}
	if !helpers.ValidateHash(users.Password, res.Password) {
		return Domain{}, exceptions.ErrValidationFailed
	}
	res.Token, _ = uu.JWTAuth.GenerateToken(res.Email)
	return res, nil
}
