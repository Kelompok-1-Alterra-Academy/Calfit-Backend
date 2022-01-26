package request

import (
	"CalFit/business/newsletters"
	"time"
)

type NewsResponse struct {
	Id                  int        `json:"id"`
	Title               string     `json:"title"`
	Description         string     `json:"description"`
	Content             string     `json:"content"`
	Operational_adminID int        `json:"operationalAdminId"`
	Created_at          *time.Time `json:"createdAt,omitempty"`
	Updated_at          *time.Time `json:"updatedAt,omitempty"`
}

func FromDomain(domain newsletters.Domain) NewsResponse {
	return NewsResponse{
		Id:                  domain.Id,
		Title:               domain.Title,
		Description:         domain.Description,
		Content:             domain.Content,
		Operational_adminID: domain.Operational_adminID,
		Created_at:          &domain.Created_at,
		Updated_at:          &domain.Updated_at,
	}
}
