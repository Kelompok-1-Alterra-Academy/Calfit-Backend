package operational_admins

import (
	"CalFit/repository/mysql/newsletters"
	"CalFit/repository/mysql/video_contents"
	"time"
)

type Operational_admin struct {
	Id             int `gorm:"primaryKey"`
	Username       string
	Password       string
	Super_adminID  int
	Newsletters    []newsletters.Newsletter
	Video_contents []video_contents.Video_content
	Created_at     time.Time
	Updated_at     time.Time
}
