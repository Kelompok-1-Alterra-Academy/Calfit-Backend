package superadmins

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Username  string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	// LoginOAuth(ctx context.Context, superadmins Domain) (Domain, error)
	Register(ctx context.Context, superadmins Domain) (Domain, error)
	// Login(ctx context.Context, superadmins Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, superadmins Domain) (Domain, error)
	Register(ctx context.Context, superadmins Domain) (Domain, error)
	GetByUsername(ctx context.Context, username string) (Domain, error)
}
