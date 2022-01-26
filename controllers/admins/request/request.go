package request

import (
	"CalFit/business/admins"
)

type OpAdminAuth struct {
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	ChangedPassword string `json:"changedPassword" form:"changed_password"`
}

func (o OpAdminAuth) ToDomain() admins.Domain {
	return admins.Domain{
		Username:        o.Username,
		Password:        o.Password,
		ChangedPassword: o.ChangedPassword,
	}
}
