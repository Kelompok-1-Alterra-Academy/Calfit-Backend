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

func FromDomain(m memberships.Domain) Memberships {
	return Memberships{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Created_at:  m.Created_at,
		Updated_at:  m.Updated_at,
	}
}
