package response

import (
	"CalFit/business/admins"
	"CalFit/business/superadmins"
	"CalFit/business/users"
	"time"
)

type Auth struct {
	Email          string     `json:"email"`
	FullName       string     `json:"fullname"`
	Token          string     `json:"token,omitempty"`
	MembershipID   int        `json:"membership_type_id"`
	MembershipName string     `json:"membership_name,omitempty"`
	Photo          string     `json:"photo,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type OpAdmin struct {
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
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
		MembershipID:   u.MembershipTypeID,
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

func FromDomainOpAdmin(o admins.Domain) OpAdmin {
	return OpAdmin{
		Username: o.Username,
		Token:    o.Token,
	}
}
