package response

import (
	"CalFit/business/superadmins"
	"CalFit/business/users"
	"time"
)

type Auth struct {
	Email          string     `json:"email"`
	FullName       string     `json:"fullname"`
	Token          string     `json:"token,omitempty"`
	MembershipName string     `json:"membership_name,omitempty"`
	Photo          string     `json:"photo,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type Superadmin struct {
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func FromDomain(u users.Domain) Auth {
	return Auth{
		Email:          u.Email,
		FullName:       u.FullName,
		Token:          u.Token,
		MembershipName: u.MembershipName,
		Photo:          u.Photo,
		CreatedAt:      &u.CreatedAt,
		UpdatedAt:      &u.UpdatedAt,
	}
}

func FromDomainSuperadmin(s superadmins.Domain) Superadmin {
	return Superadmin{
		Username: s.Username,
		Token:    s.Token,
	}
}
