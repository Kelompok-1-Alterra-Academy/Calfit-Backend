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

// func (uu *SuperadminsUsecase) LoginOAuth(ctx context.Context, users Domain) (Domain, error) {
// 	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
// 	defer cancel()
// 	if users.Email == "" || users.Password == "" {
// 		return Domain{}, exceptions.ErrInvalidCredentials
// 	}
// 	res, err := uu.Repo.LoginOAuth(ctx, users)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	res.Token, _ = uu.JWTAuth.GenerateToken(res.Id, res.Email, true, false, false)
// 	return res, nil
// }

func (su *SuperadminsUsecase) Register(ctx context.Context, superadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, su.ContextTimeout)
	defer cancel()
	if superadmin.Username == "" || superadmin.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	var err error
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

// func (uu *SuperadminsUsecase) Login(ctx context.Context, users Domain) (Domain, error) {
// 	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
// 	defer cancel()
// 	if users.Email == "" || users.Password == "" {
// 		return Domain{}, exceptions.ErrInvalidCredentials
// 	}
// 	res, err := uu.Repo.GetByUsername(ctx, users.Email)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	if !helpers.ValidateHash(users.Password, res.Password) {
// 		return Domain{}, exceptions.ErrValidationFailed
// 	}
// 	res.Token, _ = uu.JWTAuth.GenerateToken(res.Id, res.Email, true, false, false)
// 	return res, nil
// }
