package response

import (
	"CalFit/business/gyms"
	addresses "CalFit/controllers/addresses/response"
	"time"
)

type GymResponse struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	Telephone            string `json:"telephone"`
	Picture              string `json:"picture"`
	Operational_admin_ID uint   `json:"operationalAdminId"`
	// Operational_admin         Admin      `json:"operationalAdmin"`
	// Address_ID         uint      `json:"addressId"`
	Address    addresses.AddressResponse `json:"address"`
	Classes    []ClassResponse           `json:"classes"`
	Created_at time.Time                 `json:"createdAt"`
	Updated_at time.Time                 `json:"updatedAt"`
}

type ClassResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Banner_picture_url string `json:"banner_picture_url"`
	Card_picture_url   string `json:"card_picture_url"`
	Category           string `json:"category"`
	Status             string `json:"status"`
	Membership_typeID  uint   `json:"membership_typeID"`
	// Booking_details    []booking_details.Domain
	// Schedules          []schedules.Domain
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}

func FromClassDomain(domain gyms.ClassDomain) ClassResponse {
	return ClassResponse{
		ID:                 domain.Id,
		Name:               domain.Name,
		Description:        domain.Description,
		Banner_picture_url: domain.Banner_picture_url,
		Card_picture_url:   domain.Card_picture_url,
		Category:           domain.Category,
		Status:             domain.Status,
		// Membership_typeID:  domain.Membership_typeID,
		// Booking_details:    domain.Booking_details,
		// Schedules:          domain.Schedules,
		Created_at: domain.Created_at,
		Updated_at: domain.Updated_at,
	}
}

func FromClassDomainList(domain []gyms.ClassDomain) []ClassResponse {
	var response []ClassResponse
	for _, item := range domain {
		response = append(response, FromClassDomain(item))
	}
	return response
}

func FromDomain(domain gyms.Domain) GymResponse {
	return GymResponse{
		ID:                   domain.Id,
		Name:                 domain.Name,
		Telephone:            domain.Telephone,
		Picture:              domain.Picture,
		Operational_admin_ID: domain.Operational_admin_ID,
		Address:              addresses.FromDomain(domain.Address),
		Classes:              FromClassDomainList(domain.Classes),
		Created_at:           domain.Created_at,
		Updated_at:           domain.Updated_at,
	}
}
