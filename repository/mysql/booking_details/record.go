package bookingdetails

import (
	bookingdetails "CalFit/business/booking_details"
	"time"

	"gorm.io/gorm"
)

type Booking_detail struct {
	Id                 int `gorm:"primaryKey"`
	Amount             int
	Status             string
	UserID             int
	OperationalAdminID int
	PaymentID          int
	ClassID            int
	ScheduleID         int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (repo *Booking_detail) BeforeCreate(tx *gorm.DB) error {
	repo.CreatedAt = time.Now()
	repo.UpdatedAt = time.Now()
	return nil
}

func FromDomain(domain bookingdetails.Domain) Booking_detail {
	return Booking_detail{
		Id:                 domain.Id,
		Amount:             domain.Amount,
		Status:             domain.Status,
		UserID:             domain.UserID,
		OperationalAdminID: domain.OperationalAdminID,
		PaymentID:          domain.PaymentID,
		ClassID:            domain.ClassID,
		ScheduleID:         domain.ScheduleID,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}

func (repo Booking_detail) ToDomain() bookingdetails.Domain {
	return bookingdetails.Domain{
		Id:                 repo.Id,
		Amount:             repo.Amount,
		Status:             repo.Status,
		UserID:             repo.UserID,
		OperationalAdminID: repo.OperationalAdminID,
		PaymentID:          repo.PaymentID,
		ClassID:            repo.ClassID,
		ScheduleID:         repo.ScheduleID,
		CreatedAt:          repo.CreatedAt,
		UpdatedAt:          repo.UpdatedAt,
	}
}
