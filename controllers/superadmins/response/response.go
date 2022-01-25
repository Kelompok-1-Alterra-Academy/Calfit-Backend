package response

import (
	"CalFit/business/superadmins"
)

type Superadmin struct {
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func FromDomainSuperadmin(s superadmins.Domain) Superadmin {
	return Superadmin{
		Username: s.Username,
		Token:    s.Token,
	}
}
