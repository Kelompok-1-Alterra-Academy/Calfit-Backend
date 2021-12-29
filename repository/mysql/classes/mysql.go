package classes

import (
	"CalFit/business/classes"
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