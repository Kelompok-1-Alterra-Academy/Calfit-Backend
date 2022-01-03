package gyms

import (
	"CalFit/business/gyms"
	"CalFit/exceptions"
	"CalFit/repository/mysql/addresses"
	"context"
	"time"

	"gorm.io/gorm"
)

type GymRepository struct {
	Conn *gorm.DB
}

func NewGymRepository(conn *gorm.DB) gyms.DomainRepository {
	return &GymRepository{Conn: conn}
}

func (b *GymRepository) GetAll(ctx context.Context) ([]gyms.Domain, error) {
	var gymsModel []Gym
	if err := b.Conn.Preload("Address").Preload("Classes").Find(&gymsModel).Error; err != nil {
		return nil, err
	}
	var result []gyms.Domain = ToListDomain(gymsModel)
	return result, nil
}

func (b *GymRepository) GetById(ctx context.Context, id string) (gyms.Domain, error) {
	var gym Gym
	if err := b.Conn.Preload("Address").Preload("Classes").Where("id = ?", id).First(&gym).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gyms.Domain{}, exceptions.ErrNotFound
		}
		return gyms.Domain{}, err
	}
	return gym.ToDomain(), nil
}

func (b *GymRepository) Create(ctx context.Context, gym gyms.Domain) (gyms.Domain, error) {
	var gymModel Gym

	// insert address
	createdAddress := addresses.Address{
		Address:   gym.Address.Address,
		District: gym.Address.District,
		City:      gym.Address.City,
		Postal_code: gym.Address.Postal_code,
	}
	insertAddressErr := b.Conn.Create(&createdAddress).Error
	if insertAddressErr != nil {	
		return gyms.Domain{}, insertAddressErr
	}

	// insert gym
	createdGym := Gym{
		Name:      			 gym.Name,
		Telephone: 			 gym.Telephone,
		Picture:   			 gym.Picture,
		Operational_adminID: gym.Operational_admin_ID,
		AddressID:         	 createdAddress.Id,
	}
	insertGymErr := b.Conn.Create(&createdGym).Error
	if insertGymErr != nil {
		return gyms.Domain{}, insertGymErr
	}

	// get gym data
	if  getErr := b.Conn.Preload("Address").Where("id = ?", createdGym.Id).First(&gymModel).Error; getErr != nil {
		return gyms.Domain{}, getErr
	}

	return gymModel.ToDomain(), nil
}

func (b *GymRepository) Update(ctx context.Context, id string, gym gyms.Domain) (gyms.Domain, error) {
	var gymModel Gym
	if err := b.Conn.Where("id = ?", id).Preload("Address").First(&gymModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gyms.Domain{}, exceptions.ErrGymNotFound
		}
		return gyms.Domain{}, err
	}
	
	gymModel.Name = gym.Name
	gymModel.Telephone = gym.Telephone
	gymModel.Picture = gym.Picture
	gymModel.Operational_adminID = gym.Operational_admin_ID
	gymModel.Updated_at = time.Now()

	updateErr := b.Conn.Save(&gymModel).Error
	if updateErr != nil {
		return gyms.Domain{}, updateErr
	}

	return gymModel.ToDomain(), nil
}

func (b *GymRepository) Delete(ctx context.Context, id string) error {
	var gymModel Gym
	if err := b.Conn.Where("id = ?", id).First(&gymModel).Error; err != nil {
		return err
	}
	deleteErr := b.Conn.Delete(&gymModel).Error
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}