package response

import "CalFit/business/users"

type Auth struct {
	Email          string `json:"email"`
	Token          string `json:"token,omitempty"`
	MembershipName string `json:"membership_name"`
}

func FromDomain(u users.Domain) Auth {
	return Auth{
		Email:          u.Email,
		Token:          u.Token,
		MembershipName: u.MembershipName,
	}
}
