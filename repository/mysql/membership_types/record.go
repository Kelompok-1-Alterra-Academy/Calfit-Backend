package membership_types

import (
	"CalFit/business/memberships"
	"CalFit/repository/mysql/classes"
	"CalFit/repository/mysql/users"
	"time"
)

type Membership_type struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Users       []users.User    `gorm:"foreignkey:UserID"`
	Classes     []classes.Class `gorm:"foreignkey:GymID"`
	Created_at  time.Time
	Updated_at  time.Time
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
func (s Membership_type) toDomain() memberships.Domain {
	return memberships.Domain{
		Id:          s.Id,
		Name:        s.Name,
		Description: s.Description,
		Created_at:  s.Created_at,
		Updated_at:  s.Updated_at,
	}
}
