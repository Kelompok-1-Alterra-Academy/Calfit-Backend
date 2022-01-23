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
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	data.MembershipTypeID = 1
		// 	data.CreatedAt = time.Now()
		// 	address := addresses.Address{
		// 		Address:     "default",
		// 		District:    "default",
		// 		City:        "default",
		// 		Postal_code: "11111",
		// 	}
		// 	repo.DBConn.Create(&address)
		// 	data.AddressID = address.Id
		// 	repo.DBConn.Create(&data)
		// 	return data.ToDomain(), nil
		// }
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
	return superadmins.Domain{}, exceptions.ErrUserAlreadyExists
}

func (repo *SuperadminsRepo) GetByUsername(ctx context.Context, username string) (superadmins.Domain, error) {
	data := Superadmin{}
	if err := repo.DBConn.Where("username=?", username).First(&data).Error; err != nil {
		return superadmins.Domain{}, err
	}
	return data.ToDomain(), nil
}
