package request

import (
	"CalFit/business/memberships"
	"time"
)

type Memberships struct {
	Id          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func ToDomain(s Memberships) memberships.Domain {
	if s.Id == 0 {
		return memberships.Domain{
			Name:        s.Name,
			Description: s.Description,
			Created_at:  time.Now(),
		}
	} else if s.Id != 0 {
		return memberships.Domain{
			Id:          s.Id,
			Name:        s.Name,
			Description: s.Description,
			Updated_at:  time.Now(),
		}
	}
	return memberships.Domain{
		Id: s.Id,
	}
}
