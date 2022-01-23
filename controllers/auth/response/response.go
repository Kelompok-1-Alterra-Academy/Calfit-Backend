package response

import "CalFit/business/users"

type Auth struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Token          string `json:"token,omitempty"`
	MembershipName string `json:"membership_name,omitempty"`
}

func FromDomain(u users.Domain) Auth {
	return Auth{
		Email:          u.Email,
		Username:       u.Username,
		Token:          u.Token,
		MembershipName: u.MembershipName,
	}
}
