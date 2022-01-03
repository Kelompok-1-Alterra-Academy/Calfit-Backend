package users

import (
	"CalFit/repository/mysql/addresses"
	bookingdetails "CalFit/repository/mysql/booking_details"
	"time"
)

type User struct {
	Id                int `gorm:"primaryKey"`
	Username          string
	Email             string
	Photo             string
	Password          string
	Membership_typeID int
	AddressID         int
	Booking_details   []bookingdetails.Booking_detail
	Address           addresses.Address
	Created_at        time.Time
	Updated_at        time.Time
}
