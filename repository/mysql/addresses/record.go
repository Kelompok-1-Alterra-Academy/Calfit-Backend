package addresses

import (
	"CalFit/business/addresses"
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Id          uint `gorm:"primaryKey"`
	Address     string `gorm:"type:varchar(512);not null"`
	District    string `gorm:"type:varchar(256);not null"`
	City        string `gorm:"type:varchar(256);not null"`
	Postal_code string `gorm:"type:varchar(5);not null"`
	Created_at  time.Time
	Updated_at  time.Time
}

func (a *Address) BeforeCreate(tx *gorm.DB) error {
	a.Created_at = time.Now()
	a.Updated_at = time.Now()
	return nil
}

func (a *Address) ToDomain() addresses.Domain {
	return addresses.Domain{
		Id:          a.Id,
		Address: 	 a.Address,
		District:    a.District,
		City:        a.City,
		Postal_code: a.Postal_code,
		Created_at:  a.Created_at,
		Updated_at:  a.Updated_at,
	}
}

func FromDomain(domain addresses.Domain) Address {
	return Address{
		Id:          domain.Id,
		Address: 	 domain.Address,
		District:    domain.District,
		City:        domain.City,
		Postal_code: domain.Postal_code,
		Created_at:  domain.Created_at,
		Updated_at:  domain.Updated_at,
	}
}

func ToListDomain(data []Address) []addresses.Domain {
	var listDomain []addresses.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}