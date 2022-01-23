package request

import (
	"CalFit/business/superadmins"
)

type SuperadminAuth struct {
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"newPassword" form:"new_password"`
}

func (s SuperadminAuth) ToDomain() superadmins.Domain {
	return superadmins.Domain{
		Username:    s.Username,
		Password:    s.Password,
		NewPassword: s.NewPassword,
	}
}
