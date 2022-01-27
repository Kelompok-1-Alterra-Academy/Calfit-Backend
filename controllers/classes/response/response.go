package response

import (
	"CalFit/business/classes"
	"CalFit/business/schedules"
	"CalFit/controllers/schedules/response"
	"time"
)

type ClassResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Banner_picture_url string `json:"banner_picture_url,omitempty"`
	Card_picture_url   string `json:"card_picture_url,omitempty"`
	Online             bool   `json:"online"`
	Link               string `json:"link,omitempty"`
	Category           string `json:"category,omitempty"`
	Status             string `json:"status,omitempty"`
	Membership_typeID  uint   `json:"membership_typeID,omitempty"`
	Price              int    `json:"price,omitempty"`
	// Booking_details    []booking_details.Domain
	GymName    string               `json:"gym_name,omitempty"`
	Schedules  []response.Schedules `json:"schedules,omitempty"`
	Created_at *time.Time           `json:"createdAt,omitempty"`
	Updated_at *time.Time           `json:"updatedAt,omitempty"`
}

func FromDomain(domain classes.Domain) ClassResponse {
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
		// Membership_typeID:  domain.Membership_typeID,
		GymName:           domain.GymName,
		Price:             domain.Price,
		Membership_typeID: domain.Membership_typeID,
		// Booking_details:    domain.Booking_details,
		Schedules:  toListSchedules(domain.Schedules),
		Created_at: &domain.Created_at,
		Updated_at: &domain.Updated_at,
	}
}

func toListSchedules(domain []schedules.Domain) []response.Schedules {
	req := []response.Schedules{}
	for _, val := range domain {
		req = append(req, toReqSchedule(val))
	}
	return req
}

func toReqSchedule(domain schedules.Domain) response.Schedules {
	return response.Schedules{
		ID:           domain.Id,
		TimeSchedule: domain.TimeSchedule,
		Duration:     domain.Duration,
		SessionID:    domain.Duration,
	}
}
