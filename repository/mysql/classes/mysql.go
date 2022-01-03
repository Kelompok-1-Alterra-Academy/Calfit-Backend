package classes

import (
	"CalFit/business/classes"
	"CalFit/exceptions"
	"context"

	"gorm.io/gorm"
)

type ClassRepository struct {
	Conn *gorm.DB
}

func NewClassRepository(conn *gorm.DB) classes.DomainRepository {
	return &ClassRepository{Conn: conn}
}

func (c *ClassRepository) GetAll(ctx context.Context) ([]classes.Domain, error) {
	var classesModel []Class
	// if err := c.Conn.Preload("Address").Find(&classesModel).Error; err != nil {
	if err := c.Conn.Find(&classesModel).Error; err != nil {
		return nil, err
	}
	var result []classes.Domain = ToListDomain(classesModel)
	return result, nil
}

func (b *ClassRepository) GetById(ctx context.Context, id string) (classes.Domain, error) {
	var class Class
	if err := b.Conn.Preload("Address").Where("id = ?", id).First(&class).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return classes.Domain{}, exceptions.ErrNotFound
		}
		return classes.Domain{}, err
	}
	return class.ToDomain(), nil
}

func (b *ClassRepository) Create(ctx context.Context, class classes.Domain, gymId string) (classes.Domain, error) {
	var classModel Class

	createdClass := Class{
		Name:               class.Name,
		Description:        class.Description,
		Banner_picture_url: class.Banner_picture_url,
		Card_picture_url:   class.Card_picture_url,
		Category:           class.Category,
		Status:             class.Status,
		GymID:              class.GymID,
		// Membership_typeID:  class.Membership_typeID,
	}
	err := b.Conn.Create(&createdClass).Error
	if err != nil {
		return classes.Domain{}, err
	}

	// get gym data
	if getErr := b.Conn.Preload("Gym").Where("id = ?", createdClass.Id).First(&classModel).Error; getErr != nil {
		return classes.Domain{}, getErr
	}

	return classModel.ToDomain(), nil
}
