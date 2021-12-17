package video_contents

import "time"

type Video_content struct {
	Id                  int `gorm:"primaryKey"`
	Title               string
	Url                 string
	Description         string
	Operational_adminID int
	Created_at          time.Time
	Updated_at          time.Time
}
