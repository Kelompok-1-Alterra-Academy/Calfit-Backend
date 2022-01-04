package response

import (
	"CalFit/business/memberships"
	"time"
)

type Memberships struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func FromDomain(s memberships.Domain) Memberships {
	return Memberships{
		Id:          s.Id,
		Name:        s.Name,
		Description: s.Description,
		Created_at:  s.Created_at,
		Updated_at:  s.Updated_at,
	}
}
