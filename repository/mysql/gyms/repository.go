package gyms

import (
	"CalFit/business/gyms"
	"context"
	"log"
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
	if err := b.Conn.Find(&gymsModel).Error; err != nil {
		return nil, err
	}
	var result []gyms.Domain = ToListDomain(gymsModel)
	return result, nil
}

func (b *GymRepository) GetById(ctx context.Context, id string) (gyms.Domain, error) {
	var gym Gym
	if err := b.Conn.Where("id = ?", id).First(&gym).Error; err != nil {
		return gyms.Domain{}, err
	}
	log.Println(gym.Address)
	return gym.ToDomain(), nil
}

func (b *GymRepository) Create(ctx context.Context, gym gyms.Domain) (gyms.Domain, error) {
	createdGym := Gym{
		Name:      			 gym.Name,
		Telephone: 			 gym.Telephone,
		Picture:   			 gym.Picture,
		Operational_adminID: gym.Operational_admin_ID,
		AddressID:         	 1,
	}
	createdGym.BeforeCreate()
	
	insertErr := b.Conn.Create(&createdGym).Error
	if insertErr != nil {
		return gyms.Domain{}, insertErr
	}
	return createdGym.ToDomain(), nil
}

func (b *GymRepository) Update(ctx context.Context, id string, gym gyms.Domain) (gyms.Domain, error) {
	var gymModel Gym
	if err := b.Conn.Where("id = ?", id).First(&gymModel).Error; err != nil {
		return gyms.Domain{}, err
	}

	gymModel.Name = gym.Name
	gymModel.Telephone = gym.Telephone
	gymModel.Picture = gym.Picture
	gymModel.Operational_adminID = gym.Operational_admin_ID
	gymModel.AddressID = 1
	gymModel.Updated_at = time.Now()

	updateErr := b.Conn.Save(&gymModel).Error
	if updateErr != nil {
		return gyms.Domain{}, updateErr
	}
	return gymModel.ToDomain(), nil
}

// func (b *GymRepository) Delete(user *User) error {
// 	return b.Conn.Delete(user).Error
// }
