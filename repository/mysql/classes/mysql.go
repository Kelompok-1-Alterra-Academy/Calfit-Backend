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
