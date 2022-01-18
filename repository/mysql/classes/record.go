package classes

import (
	"CalFit/business/classes"
	"CalFit/business/schedules"
	bookingdetails "CalFit/repository/mysql/booking_details"

	// memberships "CalFit/repository/mysql/membership_types"

	// "CalFit/repository/mysql/gyms"
	schedulesRepo "CalFit/repository/mysql/schedules"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	Id                 uint `gorm:"primaryKey"`
	Name               string
	Description        string
	Banner_picture_url string
	Card_picture_url   string
	Online             bool
	Link               string
	Category           string
	Status             string
	Membership_typeID  uint
	GymID              uint
	Price              int
	// Gym				   gyms.Gym
	// Membership_type memberships.Membership_type
	Booking_details []bookingdetails.Booking_detail
	Schedules       []schedulesRepo.Schedule `gorm:"many2many:class_schedules"`
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
		Online:             c.Online,
		Link:               c.Link,
		Category:           c.Category,
		Status:             c.Status,
		GymID:              c.GymID,
		Price:              c.Price,
		Created_at:         c.Created_at,
		Updated_at:         c.Updated_at,
		Schedules:          convertToSchedulesArray(c.Schedules),
	}
}

func FromDomain(domain classes.Domain) Class {
	return Class{
		Id:                 domain.Id,
		Name:               domain.Name,
		Description:        domain.Description,
		Banner_picture_url: domain.Banner_picture_url,
		Card_picture_url:   domain.Card_picture_url,
		Online:             domain.Online,
		Link:               domain.Link,
		Category:           domain.Category,
		Status:             domain.Status,
		Price:              domain.Price,
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

func convertToSchedulesArray(schedule []schedulesRepo.Schedule) []schedules.Domain {
	schedulesDomain := []schedules.Domain{}
	for _, val := range schedule {
		schedulesDomain = append(schedulesDomain, toScheduleDomain(val))
	}
	return schedulesDomain
}

func toScheduleDomain(schedule schedulesRepo.Schedule) schedules.Domain {
	return schedules.Domain{
		Id:           schedule.Id,
		TimeSchedule: schedule.TimeSchedule,
		Duration:     schedule.Duration,
		SessionID:    schedule.SessionID,
		CreatedAt:    schedule.CreatedAt,
		UpdatedAt:    schedule.UpdatedAt,
	}
}
