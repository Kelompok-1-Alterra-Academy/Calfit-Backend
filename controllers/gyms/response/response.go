package response

import (
	"time"
)

type GymResponse struct {
	ID             uint      `json:"id"`
	Name      		string `json:"name"`
	Telephone 		string `json:"telephone"`
	Picture        string   `json:"picture"`
	Operational_admin_ID         uint      `json:"operational_admin_id"`
	// Operational_admin         Admin      `json:"operational_admin_id"`
	Address_ID         uint      `json:"address_id"`
	// Address         Address      `json:"address"`
	// Classes			[]Class		`json:"classes"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}