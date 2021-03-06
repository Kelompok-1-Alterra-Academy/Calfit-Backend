package request

import "CalFit/business/users"

type KeyStruct struct {
	Key string `json:"key"`
}

type Users struct {
	Email            string `json:"email" form:"email"`
	Photo            string `json:"photo" form:"photo"`
	Password         string `json:"password" form:"password"`
	MembershipTypeID int    `json:"membershipTypeId" form:"membershipTypeId"`
	// AddressID        uint					`json:"email" form:"email"`
	Token       string `json:"token" form:"token"`
	FullName    string `json:"fullname" form:"fullname"`
	Address     string `json:"address" form:"address"`
	District    string `json:"district" form:"district"`
	City        string `json:"city" form:"city"`
	Postal_code string `json:"postalCode" form:"postalCode"`
}

func (a Users) ToDomain() users.Domain {
	return users.Domain{
		Email:            a.Email,
		FullName:         a.FullName,
		Password:         a.Password,
		MembershipTypeID: a.MembershipTypeID,
		Photo:            a.Photo,
	}
}
