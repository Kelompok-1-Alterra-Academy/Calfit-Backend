package newsletters

import "time"

type Newsletter struct {
	Id                  int `gorm:"primaryKey"`
	Title               string
	Content             string
	Operational_adminID int
	Created_at          time.Time
	Updated_at          time.Time
}
