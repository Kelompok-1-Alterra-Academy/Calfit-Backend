package users

import (
	"CalFit/business/users"
	"CalFit/exceptions"
	"CalFit/repository/mysql/addresses"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DBConn *gorm.DB
}

func NewUsersRepo(db *gorm.DB) users.Repository {
	return &UsersRepo{
		DBConn: db,
	}
}

func (repo *UsersRepo) LoginOAuth(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "11111",
			}
			repo.DBConn.Debug().Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Debug().Create(&data)
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *UsersRepo) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "11111",
			}
			repo.DBConn.Debug().Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Debug().Create(&data)
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return users.Domain{}, exceptions.ErrUserAlreadyExists
}

func (repo *UsersRepo) GetByUsername(ctx context.Context, email string) (users.Domain, error) {
	data := User{}
	if err := repo.DBConn.Debug().Where("email=?", email).First(&data).Error; err != nil {
		return users.Domain{}, err
	}
	return data.ToDomain(), nil
}
