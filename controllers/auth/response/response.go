package response

import "CalFit/business/users"

type Auth struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromDomain(u users.Domain) Auth {
	return Auth{
		Email: u.Email,
		Token: u.Token,
	}
}
