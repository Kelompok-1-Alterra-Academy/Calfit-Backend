package request

import (
	"CalFit/business/admins"
	"CalFit/business/superadmins"
	"CalFit/business/users"
)

type Auth struct {
	Username string `json:"username,omitempty" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SuperadminAuth struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type OperationalAdminAuth struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (s SuperadminAuth) ToDomain() superadmins.Domain {
	return superadmins.Domain{
		Username: s.Username,
		Password: s.Password,
	}
}

func (a Auth) ToDomain() users.Domain {
	return users.Domain{
		Email:    a.Email,
		Password: a.Password,
	}
}

func (o OperationalAdminAuth) ToDomain() admins.Domain {
	return admins.Domain{
		Username: o.Username,
		Password: o.Password,
	}
}
