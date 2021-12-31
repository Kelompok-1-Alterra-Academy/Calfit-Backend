package response

import (
	"CalFit/business/schedules"
	"CalFit/business/sessions"
)

type Sessions struct {
	Name        string             `json:"name" form:"name"`
	Description string             `json:"description" form:"description"`
	Schedules   []schedules.Domain `json:",omitempty"`
}

func FromDomain(s sessions.Domain) Sessions {
	return Sessions{
		Name:        s.Name,
		Description: s.Description,
		Schedules:   s.Schedules,
	}
}
