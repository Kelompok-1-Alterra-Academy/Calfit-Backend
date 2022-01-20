package membership_types

import (
	"CalFit/business/memberships"
	classesRepo "CalFit/repository/mysql/classes"
	usersRepo "CalFit/repository/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Membership_type struct {
	gorm.Model
	Id          int                 `gorm:"primaryKey"`
	Name        string              `gorm:"type:varchar(50);not null"`
	Description string              `gorm:"type:varchar(500);not null"`
	Price       int                 `gorm:"type:int;not null"`
	Users       []usersRepo.User    `gorm:"foreignkey:UserID"`
	Classes     []classesRepo.Class `gorm:"foreignkey:ClassID"`
	Created_at  time.Time
	Updated_at  time.Time
}

func FromDomain(domain memberships.Domain) Membership_type {
	return Membership_type{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
}
func (m *Membership_type) ToDomain() memberships.Domain {
	return memberships.Domain{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Created_at:  m.Created_at,
		Updated_at:  m.Updated_at,
	}
}

func ToListDomain(data []Membership_type) []memberships.Domain {
	var listDomain []memberships.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}

func ToClassDomain(data classesRepo.Class) memberships.ClassDomain {
	return memberships.ClassDomain{
		Id:                 data.Id,
		Name:               data.Name,
		Description:        data.Description,
		Banner_picture_url: data.Banner_picture_url,
		Card_picture_url:   data.Card_picture_url,
		Online:             data.Online,
		Link:               data.Link,
		Category:           data.Category,
		Membership_typeID:  data.Membership_typeID,
		Status:             data.Status,
		Price:              data.Price,
		Created_at:         data.Created_at,
		Updated_at:         data.Updated_at,
	}
}

func ToListClassDomain(data []classesRepo.Class) []memberships.ClassDomain {
	var listDomain []memberships.ClassDomain
	for _, item := range data {
		listDomain = append(listDomain, ToClassDomain(item))
	}
	return listDomain
}

func ToUserDomain(u usersRepo.User) memberships.UserDomain {
	return memberships.UserDomain{
		Id:        u.Id,
		Email:     u.Email,
		Photo:     u.Photo,
		Password:  u.Password,
		AddressID: u.AddressID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToListUserDomain(u []usersRepo.User) []memberships.UserDomain {
	var listDomain []memberships.UserDomain
	for _, item := range u {
		listDomain = append(listDomain, ToUserDomain(item))
	}
	return listDomain
}
