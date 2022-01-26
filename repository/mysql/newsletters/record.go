package newsletters

import (
	"CalFit/business/newsletters"
	"time"

	"gorm.io/gorm"
)

type Newsletter struct {
	Id                  int `gorm:"primaryKey"`
	Title               string
	Description         string
	Content             string
	Url_Picture         string
	Operational_adminID int
	Created_at          time.Time
	Updated_at          time.Time
}

func (n *Newsletter) BeforeCreate(tx *gorm.DB) error {
	n.Created_at = time.Now()
	n.Updated_at = time.Now()
	return nil
}

func (n *Newsletter) ToDomain() newsletters.Domain {
	return newsletters.Domain{
		Id:                  n.Id,
		Title:               n.Title,
		Description:         n.Description,
		Content:             n.Content,
		Url_Picture:         n.Url_Picture,
		Operational_adminID: n.Operational_adminID,
		Created_at:          n.Created_at,
		Updated_at:          n.Updated_at,
	}
}

func FromDomain(domain newsletters.Domain) Newsletter {
	return Newsletter{
		Id:                  domain.Id,
		Title:               domain.Title,
		Description:         domain.Description,
		Content:             domain.Content,
		Operational_adminID: domain.Operational_adminID,
		Url_Picture:         domain.Url_Picture,
		Created_at:          domain.Created_at,
		Updated_at:          domain.Updated_at,
	}
}

func ToListDomain(data []Newsletter) []newsletters.Domain {
	var listDomain []newsletters.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}
