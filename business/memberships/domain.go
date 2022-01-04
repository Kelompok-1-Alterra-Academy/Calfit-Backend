package memberships

import (
	"time"
)

type Domain struct {
	Id          int
	Name        string
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}

type Usecase interface {
	Insert(memberships Domain) (Domain, error)
	Get(memberships Domain) ([]Domain, error)
	Update(memberships Domain) (Domain, error)
	Delete(memberships Domain) (Domain, error)
}

type Repository interface {
	Insert(memberships Domain) (Domain, error)
	Get(memberships Domain) ([]Domain, error)
	Update(memberships Domain) (Domain, error)
	Delete(memberships Domain) (Domain, error)
}
