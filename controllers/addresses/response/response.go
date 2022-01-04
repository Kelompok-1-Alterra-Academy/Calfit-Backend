package response

import (
	"CalFit/business/addresses"
	"time"
)

type AddressResponse struct {
	ID          uint      `json:"id"`
	Address 	string `json:"address"`
	District    string `json:"district"`
	City        string `json:"city"`
	Postal_code string `json:"postal_code"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}

func FromDomain(domain addresses.Domain) AddressResponse {
	return AddressResponse{
		ID:          domain.Id,
		Address: 	 domain.Address,
		District:    domain.District,
		City:        domain.City,
		Postal_code: domain.Postal_code,
		Created_at:  domain.Created_at,
		Updated_at:  domain.Updated_at,
	}
}