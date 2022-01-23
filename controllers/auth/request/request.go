package request

import "CalFit/business/users"

type Auth struct {
	Username string `json:"username,omitempty" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (a Auth) ToDomain() users.Domain {
	return users.Domain{
		Username: a.Username,
		Email:    a.Email,
		Password: a.Password,
	}
}
