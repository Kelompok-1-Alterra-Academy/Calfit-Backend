package gyms

import (
	"CalFit/exceptions"
	context "context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type Usecase struct {
	Repo           DomainRepository
	ContextTimeout time.Duration
}

func NewUsecase(repo DomainRepository, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (u *Usecase) GetAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return u.Repo.GetById(ctx, id)
}

func (u *Usecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()
	
	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		log.Println(err)
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.Repo.Create(ctx, domain)
}

func (u *Usecase) Update(ctx context.Context, id string, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.Repo.Update(ctx, id, domain)
}

func (u *Usecase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return exceptions.ErrEmptyInput
	}

	return u.Repo.Delete(ctx, id)
}
