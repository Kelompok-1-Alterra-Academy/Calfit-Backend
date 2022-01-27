package request

import (
	"CalFit/business/admins"
)

type OpAdmin struct {
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func FromDomainOpAdmin(o admins.Domain) OpAdmin {
	return OpAdmin{
		Username: o.Username,
		Token:    o.Token,
	}
}
