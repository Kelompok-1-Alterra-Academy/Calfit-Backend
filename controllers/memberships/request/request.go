package request

import (
	"CalFit/business/memberships"
	"time"
)

type KeyStruct struct {
	Key string `json:"key"`
}

type Memberships struct {
	Id          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func ToDomain(m Memberships) memberships.Domain {
	if m.Id == 0 {
		return memberships.Domain{
			Name:        m.Name,
			Description: m.Description,
			Created_at:  time.Now(),
		}
	} else if m.Id != 0 {
		return memberships.Domain{
			Id:          m.Id,
			Name:        m.Name,
			Description: m.Description,
			Updated_at:  time.Now(),
		}
	}
	return memberships.Domain{
		Id: m.Id,
	}
}
