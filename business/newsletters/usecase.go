package newsletters

import (
	"CalFit/business/paginations"
	"CalFit/exceptions"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

type NewsletterUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewNewsletterUsecase(repo Repository, timeout time.Duration) *NewsletterUsecase {
	return &NewsletterUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (n *NewsletterUsecase) GetAll(ctx context.Context, pagination paginations.Domain) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ContextTimeout)
	defer cancel()

	return n.Repo.GetAll(ctx, pagination)
}

func (n *NewsletterUsecase) CountAll(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ContextTimeout)
	defer cancel()

	return n.Repo.CountAll(ctx)
}

func (n *NewsletterUsecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return n.Repo.GetById(ctx, id)
}

func (n *NewsletterUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ContextTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return n.Repo.Create(ctx, domain)
}

func (n *NewsletterUsecase) Update(ctx context.Context, id string, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return n.Repo.Update(ctx, id, domain)
}

func (n *NewsletterUsecase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, n.ContextTimeout)
	defer cancel()

	if id == "" {
		return exceptions.ErrEmptyInput
	}

	return n.Repo.Delete(ctx, id)
}
