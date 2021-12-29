package response

import (
	"CalFit/business/gyms"
	addresses "CalFit/controllers/addresses/response"
	"time"
)

type GymResponse struct {
	ID             uint      `json:"id"`
	Name      		string `json:"name"`
	Telephone 		string `json:"telephone"`
	Picture        string   `json:"picture"`
	Operational_admin_ID         uint      `json:"operationalAdminId"`
	// Operational_admin         Admin      `json:"operationalAdmin"`
	// Address_ID         uint      `json:"addressId"`
	Address         addresses.AddressResponse      `json:"address"`
	// Classes			[]Class		`json:"classes"`
	Created_at      time.Time `json:"createdAt"`
	Updated_at      time.Time `json:"updatedAt"`
}

func FromDomain(domain gyms.Domain) GymResponse {
	return GymResponse{
		ID:            domain.Id,
		Name: 		domain.Name,
		Telephone: 	domain.Telephone,
		Picture: 	domain.Picture,
		Operational_admin_ID: domain.Operational_admin_ID,
		// Address_ID: 	domain.Address_ID,
		Address: 		addresses.FromDomain(domain.Address),
		Created_at:     domain.Created_at,
		Updated_at:     domain.Updated_at,
	}
}