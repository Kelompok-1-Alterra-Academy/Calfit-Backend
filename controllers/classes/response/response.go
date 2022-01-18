package response

import (
	"CalFit/business/classes"
	"CalFit/business/schedules"
	"CalFit/controllers/schedules/request"
	"fmt"
	"time"
)

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
	Price              int    `json:"price"`
	// Booking_details    []booking_details.Domain
	GymID      uint                `json:"gymID"`
	GymName    string              `json:"gym_name"`
	Schedules  []request.Schedules `json:"schedules,omitempty"`
	Created_at time.Time           `json:"createdAt"`
	Updated_at time.Time           `json:"updatedAt"`
}

func FromDomain(domain classes.Domain) ClassResponse {
	fmt.Println("ini domain", domain.Schedules)
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
		GymID:             domain.GymID,
		// Booking_details:    domain.Booking_details,
		Schedules:  toListSchedules(domain.Schedules),
		Created_at: domain.Created_at,
		Updated_at: domain.Updated_at,
	}
}

func toListSchedules(domain []schedules.Domain) []request.Schedules {
	req := []request.Schedules{}
	for _, val := range domain {
		req = append(req, toReqSchedule(val))
	}
	return req
}

func toReqSchedule(domain schedules.Domain) request.Schedules {
	return request.Schedules{
		TimeSchedule: domain.TimeSchedule,
		Duration:     domain.Duration,
		SessionID:    domain.Duration,
	}
}
