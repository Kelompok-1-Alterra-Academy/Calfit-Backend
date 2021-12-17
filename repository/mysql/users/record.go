package users

import (
	"CalFit/repository/mysql/addresses"
	"CalFit/repository/mysql/booking_details"
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
	Booking_details   []booking_details.Booking_detail
	Address           addresses.Address
	Created_at        time.Time
	Updated_at        time.Time
}
