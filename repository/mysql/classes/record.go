package classes

import (
	"CalFit/business/classes"
	bookingdetails "CalFit/repository/mysql/booking_details"

	// "CalFit/repository/mysql/gyms"
	"CalFit/repository/mysql/schedules"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	Id                 uint `gorm:"primaryKey"`
	Name               string
	Description        string
	Banner_picture_url string
	Card_picture_url   string
	Category           string
	Status             string
	// Membership_typeID  uint
	GymID uint
	// Gym				   gyms.Gym
	Booking_details []bookingdetails.Booking_detail
	Schedules       []schedules.Schedule `gorm:"many2many:class_schedules"`
	Created_at      time.Time
	Updated_at      time.Time
}

func (c *Class) BeforeCreate(tx *gorm.DB) error {
	c.Created_at = time.Now()
	c.Updated_at = time.Now()
	return nil
}

func (c *Class) ToDomain() classes.Domain {
	return classes.Domain{
		Id:                 c.Id,
		Name:               c.Name,
		Description:        c.Description,
		Banner_picture_url: c.Banner_picture_url,
		Card_picture_url:   c.Card_picture_url,
		Category:           c.Category,
		Status:             c.Status,
		GymID:              c.GymID,
		Created_at:         c.Created_at,
		Updated_at:         c.Updated_at,
	}
}

func FromDomain(domain classes.Domain) Class {
	return Class{
		Id:                 domain.Id,
		Name:               domain.Name,
		Description:        domain.Description,
		Banner_picture_url: domain.Banner_picture_url,
		Card_picture_url:   domain.Card_picture_url,
		Category:           domain.Category,
		Status:             domain.Status,
		Created_at:         domain.Created_at,
		Updated_at:         domain.Updated_at,
	}
}

func ToListDomain(data []Class) []classes.Domain {
	var listDomain []classes.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}
