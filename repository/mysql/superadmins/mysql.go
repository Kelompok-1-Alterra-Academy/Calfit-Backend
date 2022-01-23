package superadmins

import (
	"CalFit/business/superadmins"
	"CalFit/exceptions"
	"context"
	"errors"

	"gorm.io/gorm"
)

type SuperadminsRepo struct {
	DBConn *gorm.DB
}

func NewSuperadminsRepo(db *gorm.DB) superadmins.Repository {
	return &SuperadminsRepo{
		DBConn: db,
	}
}

func (repo *SuperadminsRepo) Login(ctx context.Context, domain superadmins.Domain) (superadmins.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("username=?", data.Username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return superadmins.Domain{}, exceptions.ErrSuperadminNotFound
		}
		return superadmins.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *SuperadminsRepo) Register(ctx context.Context, domain superadmins.Domain) (superadmins.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("username=?", data.Username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			repo.DBConn.Create(&data)
			return data.ToDomain(), nil
		}
		return superadmins.Domain{}, err
	}
	return superadmins.Domain{}, exceptions.ErrSuperadminExists
}

func (repo *SuperadminsRepo) GetAll(ctx context.Context) ([]superadmins.Domain, error) {
	var data []Superadmin
	if err := repo.DBConn.Find(&data).Error; err != nil {
		return nil, err
	}
	return ToListDomain(data), nil
}

func (repo *SuperadminsRepo) GetByUsername(ctx context.Context, username string) (superadmins.Domain, error) {
	data := Superadmin{}
	if err := repo.DBConn.Where("username=?", username).First(&data).Error; err != nil {
		return superadmins.Domain{}, err
	}
	return data.ToDomain(), nil
}
