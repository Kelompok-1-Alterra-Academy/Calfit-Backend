package addresses

import "time"

type Address struct {
	Id          int `gorm:"primaryKey"`
	Address     string
	District    string
	City        string
	Postal_code string
	Created_at  time.Time
	Updated_at  time.Time
}
