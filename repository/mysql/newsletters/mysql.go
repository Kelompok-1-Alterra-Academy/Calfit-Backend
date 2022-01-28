package newsletters

import (
	"CalFit/business/newsletters"
	"CalFit/business/paginations"
	"CalFit/exceptions"
	"context"

	"gorm.io/gorm"
)

type NewsletterRepository struct {
	Conn *gorm.DB
}

func NewNewsletterRepository(conn *gorm.DB) newsletters.Repository {
	return &NewsletterRepository{Conn: conn}
}

func (n *NewsletterRepository) GetAll(ctx context.Context, pagination paginations.Domain) ([]newsletters.Domain, error) {
	var newslettersModel []Newsletter

	offset := (pagination.Page - 1) * pagination.Limit
	if err := n.Conn.Limit(pagination.Limit).Offset(offset).Find(&newslettersModel).Error; err != nil {
		return nil, err
	}
	var result []newsletters.Domain = ToListDomain(newslettersModel)
	return result, nil
}

func (n *NewsletterRepository) CountAll(ctx context.Context) (int, error) {
	var count int64
	if err := n.Conn.Model(&Newsletter{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (n *NewsletterRepository) GetById(ctx context.Context, id string) (newsletters.Domain, error) {
	var news Newsletter
	if err := n.Conn.Where("id = ?", id).First(&news).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return newsletters.Domain{}, exceptions.ErrNotFound
		}
		return newsletters.Domain{}, err
	}
	return news.ToDomain(), nil
}

func (n *NewsletterRepository) Create(ctx context.Context, news newsletters.Domain) (newsletters.Domain, error) {
	createdNewsletter := Newsletter{
		Title:               news.Title,
		Description:         news.Description,
		Content:             news.Content,
		Operational_adminID: news.Operational_adminID,
		Url_Picture:         news.Url_Picture,
		Created_at:          news.Created_at,
		Updated_at:          news.Updated_at,
	}
	err := n.Conn.Create(&createdNewsletter).Error
	if err != nil {
		return newsletters.Domain{}, err
	}

	return createdNewsletter.ToDomain(), nil
}

func (n *NewsletterRepository) Update(ctx context.Context, id string, news newsletters.Domain) (newsletters.Domain, error) {
	var newsModel Newsletter
	if err := n.Conn.Where("id = ?", id).First(&newsModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return newsletters.Domain{}, exceptions.ErrClassNotFound
		}
		return newsletters.Domain{}, err
	}
	newsModel.Title = news.Title
	newsModel.Description = news.Description
	newsModel.Content = news.Content
	newsModel.Operational_adminID = news.Operational_adminID
	newsModel.Url_Picture = news.Url_Picture
	newsModel.Created_at = news.Created_at
	newsModel.Updated_at = news.Updated_at
	if err := n.Conn.Save(&newsModel).Error; err != nil {
		return newsletters.Domain{}, err
	}

	return newsModel.ToDomain(), nil
}

func (n *NewsletterRepository) Delete(ctx context.Context, id string) error {
	var news Newsletter
	if err := n.Conn.Where("id = ?", id).Delete(&news).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.ErrNotFound
		}
		return err
	}
	return nil
}
