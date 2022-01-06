package response

import "CalFit/business/users"

type Auth struct {
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}

func FromDomain(u users.Domain) Auth {
	return Auth{
		Email: u.Email,
		Token: u.Token,
	}
}
