package gyms

import (
	"CalFit/business/gyms"
	"context"
	"log"

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
	log.Println(createdGym)
	insertErr := b.Conn.Create(&createdGym).Error
	if insertErr != nil {
		return gyms.Domain{}, insertErr
	}
	return createdGym.ToDomain(), nil
}

// func (b *GymRepository) UpdateStatus(ctx context.Context, id string, status bool) (gyms.Domain, error) {
// 	var book gyms
// 	if err := b.Conn.Where("book_id = ?", id).First(&book).Error; err != nil {
// 		return gyms.Domain{}, err
// 	}
// 	book.Status = status
// 	book.UpdatedAt = time.Now()
// 	if err := b.Conn.Save(&book).Error; err != nil {
// 		return gyms.Domain{}, err
// 	}
// 	return book.ToDomain(), nil
// }

// // func (b *GymRepository) Update(user *User) error {
// // 	return b.Conn.Save(user).Error
// // }

// // func (b *GymRepository) Delete(user *User) error {
// // 	return b.Conn.Delete(user).Error
// // }
