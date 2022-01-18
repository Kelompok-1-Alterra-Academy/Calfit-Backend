package bookingdetails

import (
	bookingdetails "CalFit/business/booking_details"
	"context"

	"gorm.io/gorm"
)

type BookingDetailsRepo struct {
	DBConn *gorm.DB
}

func NewBookingDetailsRepo(db *gorm.DB) bookingdetails.Repository {
	return &BookingDetailsRepo{
		DBConn: db,
	}
}

func (repo *BookingDetailsRepo) Insert(ctx context.Context, domain bookingdetails.Domain) (bookingdetails.Domain, error) {
	data := FromDomain(domain)
	data.Status = "waiting"
	if err := repo.DBConn.Create(&data).Error; err != nil {
		return bookingdetails.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *BookingDetailsRepo) GetByUserID(ctx context.Context, userID int) ([]bookingdetails.Domain, error) {
	data := []Booking_detail{}
	if err := repo.DBConn.Where("user_id=?", userID).Find(&data).Error; err != nil {
		return []bookingdetails.Domain{}, err
	}
	var domain []bookingdetails.Domain
	for _, val := range data {
		domain = append(domain, val.ToDomain())
	}
	return domain, nil
}
