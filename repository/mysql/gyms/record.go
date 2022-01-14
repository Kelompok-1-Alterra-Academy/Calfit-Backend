package gyms

import (
	"CalFit/business/gyms"
	"CalFit/repository/mysql/addresses"
	"CalFit/repository/mysql/classes"
	"CalFit/repository/mysql/operational_admins"
	"time"

	"gorm.io/gorm"
)

type Gym struct {
	gorm.Model
	Id                  uint                                 `gorm:"primaryKey"`
	Name                string                               `gorm:"type:varchar(100);not null"`
	Description         string                               `gorm:"type:varchar(1024);not null"`
	Telephone           string                               `gorm:"type:varchar(20);not null"`
	Picture             string                               `gorm:"type:varchar(500);not null"`
	Operational_adminID uint                                 `gorm:"not null"`
	AddressID           uint                                 `gorm:"not null"`
	Operational_admin   operational_admins.Operational_admin `gorm:"foreignkey:Operational_adminID"`
	Address             addresses.Address                    `gorm:"foreignkey:AddressID"`
	Classes             []classes.Class
	Created_at          time.Time
	Updated_at          time.Time
}

func (g *Gym) BeforeCreate(tx *gorm.DB) error {
	g.Created_at = time.Now()
	g.Updated_at = time.Now()
	return nil
}

func (g *Gym) ToDomain() gyms.Domain {
	return gyms.Domain{
		Id:                   g.Id,
		Name:                 g.Name,
		Description:          g.Description,
		Telephone:            g.Telephone,
		Picture:              g.Picture,
		Operational_admin_ID: g.Operational_adminID,
		Address_ID:           g.AddressID,
		// Operational_admin:   g.Operational_admin.ToDomain(),
		Address: g.Address.ToDomain(),
		Classes: ToListClassDomain(g.Classes),
		// Classes:    g.Classes.T
		Created_at: g.Created_at,
		Updated_at: g.Updated_at,
	}
}

func FromDomain(domain gyms.Domain) Gym {
	return Gym{
		Id:                  domain.Id,
		Name:                domain.Name,
		Description:         domain.Description,
		Telephone:           domain.Telephone,
		Picture:             domain.Picture,
		Operational_adminID: domain.Operational_admin_ID,
		AddressID:           domain.Address_ID,
		// Operational_admin:   domain.Operational_admin.FromDomain(),
		// Address:             domain.Address.FromDomain(),
		// Classes:             domain.Classes.FromDomain(),
		Created_at: domain.Created_at,
		Updated_at: domain.Updated_at,
	}
}

func ToListDomain(data []Gym) []gyms.Domain {
	var listDomain []gyms.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}

func ToClassDomain(data classes.Class) gyms.ClassDomain {
	return gyms.ClassDomain{
		Id:                 data.Id,
		Name:               data.Name,
		Description:        data.Description,
		Banner_picture_url: data.Banner_picture_url,
		Card_picture_url:   data.Card_picture_url,
		Online:             data.Online,
		Link:               data.Link,
		Category:           data.Category,
		Status:             data.Status,
		Created_at:         data.Created_at,
		Updated_at:         data.Updated_at,
	}
}

func ToListClassDomain(data []classes.Class) []gyms.ClassDomain {
	var listDomain []gyms.ClassDomain
	for _, item := range data {
		listDomain = append(listDomain, ToClassDomain(item))
	}
	return listDomain
}
