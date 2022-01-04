package membership_types

import (
	"CalFit/business/memberships"
	"CalFit/repository/mysql/classes"

	// "CalFit/repository/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Membership_type struct {
	gorm.Model
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(50);not null"`
	Description string `gorm:"type:varchar(500);not null"`
	// Users       []users.User    `gorm:"foreignkey:UserID"`
	Classes    []classes.Class `gorm:"foreignkey:GymID"`
	Created_at time.Time
	Updated_at time.Time
}

func FromDomain(domain memberships.Domain) Membership_type {
	return Membership_type{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
}
func (m *Membership_type) ToDomain() memberships.Domain {
	return memberships.Domain{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Created_at:  m.Created_at,
		Updated_at:  m.Updated_at,
	}
}
