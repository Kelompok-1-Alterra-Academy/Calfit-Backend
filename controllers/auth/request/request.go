package request

import "CalFit/business/users"

type Auth struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (a Auth) ToDomain() users.Domain {
	return users.Domain{
		Email:    a.Email,
		Password: a.Password,
	}
}