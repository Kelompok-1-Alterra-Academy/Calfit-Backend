package gyms

import (
	"CalFit/repository/mysql/addresses"
	"CalFit/repository/mysql/classes"
	"CalFit/repository/mysql/operational_admins"
)

type Gym struct {
	Id                  int `gorm:"primaryKey"`
	Name                string
	Telephone           string
	Picture             string
	Operational_adminID int
	AddressID           int
	Operational_admin   operational_admins.Operational_admin
	Address             addresses.Address
	Classes             []classes.Class
	Created_at          string
	Updated_at          string
}
