package classes

import (
	"CalFit/business/classes"
	"CalFit/business/paginations"
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

func (c *ClassRepository) GetAll(ctx context.Context, pagination paginations.Domain) ([]classes.Domain, error) {
	var classesModel []Class

	offset := (pagination.Page - 1) * pagination.Limit
	if err := c.Conn.Limit(pagination.Limit).Offset(offset).Find(&classesModel).Error; err != nil {
		return nil, err
	}
	var result []classes.Domain = ToListDomain(classesModel)
	return result, nil
}

func (c *ClassRepository) GetById(ctx context.Context, id string) (classes.Domain, error) {
	var class Class
	if err := c.Conn.Where("id = ?", id).Preload("Schedules").First(&class).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return classes.Domain{}, exceptions.ErrNotFound
		}
		return classes.Domain{}, err
	}
	return class.ToDomain(), nil
}

func (c *ClassRepository) Create(ctx context.Context, class classes.Domain, gymId string) (classes.Domain, error) {
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
	err := c.Conn.Create(&createdClass).Error
	if err != nil {
		return classes.Domain{}, err
	}

	return createdClass.ToDomain(), nil
}

func (c *ClassRepository) Update(ctx context.Context, id string, class classes.Domain) (classes.Domain, error) {
	var classModel Class
	if err := c.Conn.Where("id = ?", id).First(&classModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return classes.Domain{}, exceptions.ErrClassNotFound
		}
		return classes.Domain{}, err
	}
	classModel.Name = class.Name
	classModel.Description = class.Description
	classModel.Banner_picture_url = class.Banner_picture_url
	classModel.Card_picture_url = class.Card_picture_url
	classModel.Category = class.Category
	classModel.Status = class.Status
	// classModel.GymID = class.GymID
	// classModel.Membership_typeID = class.Membership_typeID
	if err := c.Conn.Save(&classModel).Error; err != nil {
		return classes.Domain{}, err
	}

	return classModel.ToDomain(), nil
}

func (c *ClassRepository) Delete(ctx context.Context, id string) error {
	var class Class
	if err := c.Conn.Where("id = ?", id).Delete(&class).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.ErrNotFound
		}
		return err
	}
	return nil
}
