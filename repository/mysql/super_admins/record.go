package super_admins

import (
	"CalFit/repository/mysql/operational_admins"
	"time"
)

type Super_admin struct {
	Id                 int `gorm:"primaryKey"`
	Username           string
	Password           string
	Operational_admins []operational_admins.Operational_admin
	Created_at         time.Time
	Updated_at         time.Time
}
