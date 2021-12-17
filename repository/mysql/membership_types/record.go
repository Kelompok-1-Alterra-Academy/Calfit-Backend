package membership_types

import (
	"CalFit/repository/mysql/classes"
	"CalFit/repository/mysql/users"
	"time"
)

type Membership_type struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Users       []users.User
	Classes     []classes.Class
	Created_at  time.Time
	Updated_at  time.Time
}
