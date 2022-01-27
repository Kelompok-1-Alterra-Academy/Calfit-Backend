package operational_admins

import (
	"CalFit/repository/mysql/newsletters"
	"CalFit/repository/mysql/superadmins"

	// "CalFit/repository/mysql/video_contents"
	"CalFit/business/admins"
	"time"

	"gorm.io/gorm"
)

type Operational_admin struct {
	Id           int `gorm:"primaryKey"`
	Username     string
	Password     string
	SuperadminID int
	Newsletters  []newsletters.Newsletter
	// Video_contents []video_contents.Video_content
	Superadmin superadmins.Superadmin `gorm:"foreignKey:SuperadminID"`
	Created_at time.Time
	Updated_at time.Time
}

func (o *Operational_admin) BeforeCreate(tx *gorm.DB) error {
	o.Created_at = time.Now()
	o.Updated_at = time.Now()
	return nil
}

func FromDomain(domain admins.Domain) Operational_admin {
	return Operational_admin{
		Id:         domain.Id,
		Username:   domain.Username,
		Password:   domain.Password,
		Created_at: domain.Created_at,
		Updated_at: domain.Updated_at,
	}
}
func (o *Operational_admin) ToDomain() admins.Domain {
	return admins.Domain{
		Id:         o.Id,
		Username:   o.Username,
		Password:   o.Password,
		Created_at: o.Created_at,
		Updated_at: o.Updated_at,
	}
}

func ToListDomain(admin []Operational_admin) []admins.Domain {
	var listDomain []admins.Domain
	for _, item := range admin {
		listDomain = append(listDomain, item.ToDomain())
	}

	return listDomain
}
