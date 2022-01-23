package users

import (
	"CalFit/business/paginations"
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
	if err := repo.DBConn.Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "11111",
			}
			repo.DBConn.Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Create(&data)
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *UsersRepo) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "11111",
			}
			repo.DBConn.Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Create(&data)
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return users.Domain{}, exceptions.ErrUserAlreadyExists
}

func (repo *UsersRepo) GetByUsername(ctx context.Context, email string) (users.Domain, error) {
	data := User{}
	if err := repo.DBConn.Where("email=?", email).Preload("Addresses").First(&data).Error; err != nil {
		return users.Domain{}, err
	}
	domain := data.ToDomain()
	type Relation struct {
		MembershipName string
	}
	relation := Relation{}
	repo.DBConn.Table("users").Select("membership_types.name AS membershipname").Joins("LEFT JOIN membership_types ON users.membership_type_id = membership_types.id").Scan(&relation)
	domain.MembershipName = relation.MembershipName
	return domain, nil
}

func (b *UsersRepo) GetAll(ctx context.Context, pagination paginations.Domain) ([]users.Domain, error) {
	var usersModel []User

	offset := (pagination.Page - 1) * pagination.Limit
	if err := b.DBConn.Preload("Address").Preload("BookingDetails").Limit(pagination.Limit).Offset(offset).Find(&usersModel).Error; err != nil {
		return nil, err
	}
	var result []users.Domain = ToListDomain(usersModel)
	return result, nil
}

func (b *UsersRepo) CountAll(ctx context.Context) (int, error) {
	var count int64
	if err := b.DBConn.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (b *UsersRepo) GetById(ctx context.Context, id string) (users.Domain, error) {
	var user User
	if err := b.DBConn.Preload("Address").Preload("Classes").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, exceptions.ErrNotFound
		}
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}

func (b *UsersRepo) Update(ctx context.Context, id string, user users.Domain) (users.Domain, error) {
	var userModel User
	if err := b.DBConn.Where("id = ?", id).Preload("Address").First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, exceptions.ErrUserNotFound
		}
		return users.Domain{}, err
	}

	userModel.Email = user.Email
	userModel.Password = user.Password
	userModel.AddressID = user.AddressID
	userModel.Photo = user.Photo
	userModel.MembershipTypeID = user.MembershipTypeID
	userModel.FullName = user.FullName
	userModel.UpdatedAt = time.Now()

	updateErr := b.DBConn.Save(&userModel).Error
	if updateErr != nil {
		return users.Domain{}, updateErr
	}

	return userModel.ToDomain(), nil
}
