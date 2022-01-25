package superadmins

import (
	superadmins "CalFit/business/superadmins"
	"time"

	"gorm.io/gorm"
)

type Superadmin struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Password string
	// Operational_admins []operational_admins.Operational_admin
	Created_at time.Time
	Updated_at time.Time
}

func (s *Superadmin) BeforeCreate(tx *gorm.DB) error {
	s.Created_at = time.Now()
	s.Updated_at = time.Now()
	return nil
}

func (s *Superadmin) ToDomain() superadmins.Domain {
	return superadmins.Domain{
		Id:        s.Id,
		Username:  s.Username,
		Password:  s.Password,
		CreatedAt: s.Created_at,
		UpdatedAt: s.Updated_at,
	}
}

func FromDomain(domain superadmins.Domain) Superadmin {
	return Superadmin{
		Id:         domain.Id,
		Username:   domain.Username,
		Password:   domain.Password,
		Created_at: domain.CreatedAt,
		Updated_at: domain.UpdatedAt,
	}
}

func ToListDomain(data []Superadmin) []superadmins.Domain {
	var listDomain []superadmins.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}
