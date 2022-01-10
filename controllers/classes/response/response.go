package response

import (
	"CalFit/business/classes"
	"time"
)

type ClassResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Banner_picture_url string `json:"banner_picture_url"`
	Card_picture_url   string `json:"card_picture_url"`
	Category           string `json:"category"`
	Status             string `json:"status"`
	Membership_typeID  uint   `json:"membership_typeID"`
	GymID              uint   `json:"gymID"`
	// Booking_details    []booking_details.Domain
	// Schedules          []schedules.Domain `gorm:"many2many:class_schedules"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}

func FromDomain(domain classes.Domain) ClassResponse {
	return ClassResponse{
		ID:                 domain.Id,
		Name:               domain.Name,
		Description:        domain.Description,
		Banner_picture_url: domain.Banner_picture_url,
		Card_picture_url:   domain.Card_picture_url,
		Category:           domain.Category,
		Status:             domain.Status,
		// Membership_typeID:  domain.Membership_typeID,
		GymID: domain.GymID,
		// Booking_details:    domain.Booking_details,
		// Schedules:          domain.Schedules,
		Created_at: domain.Created_at,
		Updated_at: domain.Updated_at,
	}
}
