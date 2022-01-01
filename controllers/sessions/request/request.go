package request

import (
	"CalFit/business/sessions"
)

type Sessions struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func (s Sessions) ToDomain() sessions.Domain {
	return sessions.Domain{
		Name:        s.Name,
		Description: s.Description,
	}
}
