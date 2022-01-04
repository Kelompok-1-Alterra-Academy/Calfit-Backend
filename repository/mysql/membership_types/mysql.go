package membership_types

import (
	"CalFit/business/memberships"
	"CalFit/exceptions"
	"errors"

	"gorm.io/gorm"
)

type MembershipsRepo struct {
	DBConn *gorm.DB
}

func NewMembershipsRepo(db *gorm.DB) memberships.Repository {
	return &MembershipsRepo{
		DBConn: db,
	}
}

func (repo *MembershipsRepo) Insert(domain memberships.Domain) (memberships.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Create(&data).Error; err != nil {
		return memberships.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *MembershipsRepo) Get(domain memberships.Domain) ([]memberships.Domain, error) {
	data := []Membership_type{}
	if err := repo.DBConn.Debug().Find(&data).Error; err != nil {
		return []memberships.Domain{}, err
	}
	var domainMemberships []memberships.Domain
	for _, val := range data {
		domainMemberships = append(domainMemberships, val.ToDomain())
	}
	return domainMemberships, nil
}

func (repo *MembershipsRepo) Update(domain memberships.Domain) (memberships.Domain, error) {
	data := FromDomain(domain)
	repo.DBConn.Debug().Where("id=?", domain.Id)
	if err := repo.DBConn.Debug().Where("id=?", data.Id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return memberships.Domain{}, errors.New("record not found")
		}
		return memberships.Domain{}, exceptions.ErrNotFound
	}
	return data.ToDomain(), nil
}

func (repo *MembershipsRepo) Delete(domain memberships.Domain) (memberships.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Where("id=?", data.Id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return memberships.Domain{}, errors.New("record not found")
		}
	}
	repo.DBConn.Where("id=?", data.Id).Delete(&data)
	return data.ToDomain(), nil
}
