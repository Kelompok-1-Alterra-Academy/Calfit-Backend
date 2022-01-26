package operational_admins

import (
	"CalFit/business/admins"
	"CalFit/business/paginations"
	"CalFit/exceptions"
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type OperationalAdminsRepo struct {
	Conn *gorm.DB
}

func NewOperationalAdminsRepo(db *gorm.DB) admins.Repository {
	return &OperationalAdminsRepo{
		Conn: db,
	}
}

func (repo *OperationalAdminsRepo) Login(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	data := FromDomain(domain)
	if err := repo.Conn.Where("username=?", data.Username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return admins.Domain{}, exceptions.ErrOperationalAdminNotFound
		}
		return admins.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *OperationalAdminsRepo) Register(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	data := FromDomain(domain)
	if err := repo.Conn.Where("username=?", data.Username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			repo.Conn.Create(&data)
			return data.ToDomain(), nil
		}
		return admins.Domain{}, err
	}
	return admins.Domain{}, exceptions.ErrOperationalAdminExists
}

func (repo *OperationalAdminsRepo) UpdatePassword(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	admin := Operational_admin{}
	err := repo.Conn.Where("username=?", domain.Username).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return admins.Domain{}, exceptions.ErrOperationalAdminNotFound
		}
		return admins.Domain{}, err
	}
	admin.Password = domain.Password
	admin.Updated_at = time.Now()
	repo.Conn.Save(&admin)
	res := admin.ToDomain()
	log.Println(res)
	return res, nil
}

func (repo *OperationalAdminsRepo) GetAll(ctx context.Context) ([]admins.Domain, error) {
	var data []Operational_admin
	if err := repo.Conn.Find(&data).Error; err != nil {
		return nil, err
	}
	return ToListDomain(data), nil
}

func (repo *OperationalAdminsRepo) GetByUsername(ctx context.Context, username string) (admins.Domain, error) {
	data := Operational_admin{}
	if err := repo.Conn.Where("username=?", username).First(&data).Error; err != nil {
		return admins.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *OperationalAdminsRepo) Get(ctx context.Context, pagination paginations.Domain) ([]admins.Domain, error) {
	var data []Operational_admin

	offset := (pagination.Page - 1) * pagination.Limit
	if err := repo.Conn.Limit(pagination.Limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}
	var result []admins.Domain = ToListDomain(data)
	return result, nil
}

func (repo *OperationalAdminsRepo) CountAll(ctx context.Context) (int, error) {
	var count int64
	if err := repo.Conn.Model(&Operational_admin{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
