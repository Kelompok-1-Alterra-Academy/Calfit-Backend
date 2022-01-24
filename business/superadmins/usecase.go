package superadmins

import (
	"CalFit/app/middlewares"
	"CalFit/exceptions"
	"CalFit/helpers"
	"context"
	"time"
)

type SuperadminsUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JWTAuth        *middlewares.ConfigJWT
}

func NewSuperadminsUsecase(repo Repository, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &SuperadminsUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
		JWTAuth:        jwtAuth,
	}
}

func (su *SuperadminsUsecase) Register(ctx context.Context, superadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, su.ContextTimeout)
	defer cancel()
	if superadmin.Username == "" || superadmin.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}

	superadmins, err := su.Repo.GetAll(ctx)
	if err != nil {
		return Domain{}, err
	}
	if superadmins != nil {
		return Domain{}, exceptions.ErrSuperadminExists
	}

	superadmin.Password, err = helpers.Hash(superadmin.Password)
	if err != nil {
		return Domain{}, err
	}
	res, err := su.Repo.Register(ctx, superadmin)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (su *SuperadminsUsecase) Login(ctx context.Context, superadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, su.ContextTimeout)
	defer cancel()
	if superadmin.Username == "" || superadmin.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := su.Repo.GetByUsername(ctx, superadmin.Username)
	if err != nil {
		return Domain{}, err
	}
	if !helpers.ValidateHash(superadmin.Password, res.Password) {
		return Domain{}, exceptions.ErrValidationFailed
	}
	res.Token, _ = su.JWTAuth.GenerateToken(res.Id, res.Username, false, false, true)
	return res, nil
}

func (su *SuperadminsUsecase) UpdatePassword(ctx context.Context, superadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, su.ContextTimeout)
	defer cancel()
	if superadmin.Username == "" || superadmin.Password == "" || superadmin.NewPassword == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := su.Repo.Login(ctx, superadmin)
	if err != nil {
		return Domain{}, err
	}
	if !helpers.ValidateHash(superadmin.Password, res.Password) {
		return Domain{}, exceptions.ErrValidationFailed
	}
	superadmin.Password, err = helpers.Hash(superadmin.NewPassword)
	if err != nil {
		return Domain{}, err
	}
	res, err = su.Repo.UpdatePassword(ctx, superadmin)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
