package membership_types

import (
	"CalFit/business/memberships"
	"CalFit/exceptions"
	"context"
	"time"

	"gorm.io/gorm"
)

type MembershipsRepo struct {
	DBConn *gorm.DB
}

func NewMembershipsRepo(db *gorm.DB) memberships.Repository {
	return &MembershipsRepo{DBConn: db}
}

func (repo *MembershipsRepo) Insert(ctx context.Context, domain memberships.Domain) (memberships.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Create(&data).Error; err != nil {
		return memberships.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *MembershipsRepo) Get(ctx context.Context) ([]memberships.Domain, error) {
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

func (repo *MembershipsRepo) GetById(ctx context.Context, id string) (memberships.Domain, error) {
	var membershipModel Membership_type
	if err := repo.DBConn.Preload("Classes").Where("id = ?", id).First(&membershipModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return memberships.Domain{}, exceptions.ErrNotFound
		}
		return memberships.Domain{}, err
	}
	return membershipModel.ToDomain(), nil
}

func (repo *MembershipsRepo) Update(ctx context.Context, id string, membership memberships.Domain) (memberships.Domain, error) {
	var membershipModel Membership_type
	if err := repo.DBConn.Debug().Where("id = ?", id).First(&membershipModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return memberships.Domain{}, exceptions.ErrMembershipNotFound
		}
		return memberships.Domain{}, exceptions.ErrNotFound

	}

	membershipModel.Name = membership.Name
	membershipModel.Description = membership.Description
	membershipModel.Price = membership.Price
	membershipModel.Updated_at = time.Now()

	updateErr := repo.DBConn.Save(&membershipModel).Error
	if updateErr != nil {
		return memberships.Domain{}, updateErr
	}
	return membershipModel.ToDomain(), nil
}

func (repo *MembershipsRepo) Delete(ctx context.Context, id string) error {
	var membershipModel Membership_type
	if err := repo.DBConn.Debug().Where("id=?", id).First(&membershipModel).Error; err != nil {
		return err
	}
	deleteErr := repo.DBConn.Delete(&membershipModel).Error
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
