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

func (repo *BookingDetailsRepo) Insert(ctx context.Context, bookingDetails bookingdetails.Domain) (bookingdetails.Domain, error)
