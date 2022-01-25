package response

import (
	"CalFit/business/admins"
	"CalFit/business/superadmins"
	"CalFit/business/users"
)

type Auth struct {
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}

type OpAdmin struct {
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

type Superadmin struct {
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func FromDomain(u users.Domain) Auth {
	return Auth{
		Email: u.Email,
		Token: u.Token,
	}
}

func FromDomainSuperadmin(s superadmins.Domain) Superadmin {
	return Superadmin{
		Username: s.Username,
		Token:    s.Token,
	}
}

func FromDomainOpAdmin(o admins.Domain) OpAdmin {
	return OpAdmin{
		Username: o.Username,
		Token:    o.Token,
	}
}
