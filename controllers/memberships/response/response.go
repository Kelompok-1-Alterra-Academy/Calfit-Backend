package response

import (
	"CalFit/business/memberships"
	"time"
)

type Memberships struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Classes     []ClassResponse `json:"classes"`
	Users       []UserResponse  `json:"users"`
	Created_at  time.Time       `json:"created_at"`
	Updated_at  time.Time       `json:"updated_at"`
}

type ClassResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Banner_picture_url string `json:"banner_picture_url"`
	Card_picture_url   string `json:"card_picture_url"`
	Online             bool   `json:"online"`
	Link               string `json:"link"`
	Category           string `json:"category"`
	Status             string `json:"status"`
	Membership_typeID  uint   `json:"membership_typeID"`
	Price              int    `json:"price" form:"price"`
	// Booking_details    []booking_details.Domain
	// Schedules          []schedules.Domain
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}

type UserResponse struct {
	Id               int    `json:"id"`
	Email            string `json:"email"`
	Photo            string `json:"photo"`
	Password         string `json:"password"`
	MembershipTypeID int    `json:"membership_typeID"`
	AddressID        uint   `json:"AddressID"`
	// BookingDetails   []bookingDetailsRepo.Booking_detail
	// Address          addresses.Address `gorm:"foreignkey:AddressID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromUserDomain(udomain memberships.UserDomain) UserResponse {
	return UserResponse{
		Id:               udomain.Id,
		Email:            udomain.Email,
		Photo:            udomain.Photo,
		Password:         udomain.Password,
		MembershipTypeID: udomain.MembershipTypeID,
		AddressID:        udomain.AddressID,
		CreatedAt:        udomain.CreatedAt,
		UpdatedAt:        udomain.UpdatedAt,
	}
}

func FromUserDomainList(udomain []memberships.UserDomain) []UserResponse {
	var response []UserResponse
	for _, item := range udomain {
		response = append(response, FromUserDomain(item))
	}
	return response
}

func FromClassDomain(domain memberships.ClassDomain) ClassResponse {
	return ClassResponse{
		ID:                 domain.Id,
		Name:               domain.Name,
		Description:        domain.Description,
		Banner_picture_url: domain.Banner_picture_url,
		Card_picture_url:   domain.Card_picture_url,
		Online:             domain.Online,
		Link:               domain.Link,
		Category:           domain.Category,
		Status:             domain.Status,
		Membership_typeID:  domain.Membership_typeID,
		// Booking_details:    domain.Booking_details,
		// Schedules:          domain.Schedules,
		Created_at: domain.Created_at,
		Updated_at: domain.Updated_at,
	}
}

func FromClassDomainList(domain []memberships.ClassDomain) []ClassResponse {
	var response []ClassResponse
	for _, item := range domain {
		response = append(response, FromClassDomain(item))
	}
	return response
}

func FromDomain(m memberships.Domain) Memberships {
	return Memberships{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Classes:     FromClassDomainList(m.Classes),
		Created_at:  m.Created_at,
		Updated_at:  m.Updated_at,
	}
}
