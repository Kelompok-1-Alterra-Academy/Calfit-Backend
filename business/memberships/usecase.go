package memberships

import (
	"CalFit/exceptions"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

type MembershipsUsecase struct {
	membershipsRepo Repository
	ContextTimeout  time.Duration
}

func NewMembershipsUsecase(repo Repository, timeout time.Duration) *MembershipsUsecase {
	return &MembershipsUsecase{
		membershipsRepo: repo,
		ContextTimeout:  timeout,
	}
}

func (u *MembershipsUsecase) Insert(ctx context.Context, memberships Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(memberships)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.membershipsRepo.Insert(ctx, memberships)
}

func (u *MembershipsUsecase) Get(ctx context.Context, memberships Domain) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.membershipsRepo.Get(ctx, memberships)
}

func (u *MembershipsUsecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return u.membershipsRepo.GetById(ctx, id)
}

func (u *MembershipsUsecase) Update(ctx context.Context, id string, memberships Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}
	validate := validator.New()
	err := validate.Struct(memberships)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.membershipsRepo.Update(ctx, id, memberships)
}

func (u *MembershipsUsecase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return exceptions.ErrEmptyInput
	}
	return u.membershipsRepo.Delete(ctx, id)
}
