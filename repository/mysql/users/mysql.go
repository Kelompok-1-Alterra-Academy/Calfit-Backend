package users

import (
	"CalFit/business/users"
	"CalFit/repository/mysql/addresses"
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var mysqlErr *mysql.MySQLError

type UsersRepo struct {
	DBConn *gorm.DB
}

func NewUsersRepo(db *gorm.DB) users.Repository {
	return &UsersRepo{
		DBConn: db,
	}
}

func (repo *UsersRepo) Login(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Debug().Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "default",
			}
			repo.DBConn.Debug().Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Debug().Create(&data)
			return data.ToDomain(), nil
		} else if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return data.ToDomain(), nil
}
