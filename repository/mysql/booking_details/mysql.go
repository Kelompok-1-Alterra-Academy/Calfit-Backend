package bookingdetails

import (
	bookingdetails "CalFit/business/booking_details"
	"context"

	"gorm.io/gorm"
)

type BookingDetailsRepo struct {
	DBConn *gorm.DB
}

func NewBookingDetailsRepo(db *gorm.DB) bookingdetails.Repository {
	return &BookingDetailsRepo{
		DBConn: db,
	}
}

func (repo *BookingDetailsRepo) Insert(ctx context.Context, domain bookingdetails.Domain) (bookingdetails.Domain, error) {
	data := FromDomain(domain)
	data.Status = "waiting"
	if err := repo.DBConn.Create(&data).Error; err != nil {
		return bookingdetails.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *BookingDetailsRepo) CountAll(ctx context.Context) (int, error) {
	var count int64
	if err := repo.DBConn.Model(&Booking_detail{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
	// data := []Booking_detail{}
	// if err := repo.DBConn.Find(&data).Error; err != nil {
	// 	return 0, err
	// }
	// return len(data), nil
}

func (repo *BookingDetailsRepo) GetByUserID(ctx context.Context, userID int) ([]bookingdetails.Domain, error) {
	data := []Booking_detail{}
	if err := repo.DBConn.Where("user_id=?", userID).Limit(5).Find(&data).Error; err != nil {
		return []bookingdetails.Domain{}, err
	}
	var domain []bookingdetails.Domain
	type Class struct {
		Name           string
		TimeSchedule   string
		GymName        string
		CardPictureUrl string
	}
	class := Class{}
	for _, val := range data {
		domain = append(domain, val.ToDomain())
	}
	for i := range domain {
		repo.DBConn.Table("booking_details").Select("classes.name as name,classes.card_picture_url as card_picture_url, schedules.time_schedule, gyms.name as gym_name").Joins("left join classes on booking_details.class_id=class_id left join schedules on booking_details.schedule_id=schedules.id left join gyms on classes.gym_id=gyms.id").Scan(&class)
		domain[i].ClassName = class.Name
		domain[i].TimeSchedule = class.TimeSchedule
		domain[i].GymName = class.GymName
		domain[i].CardPictureUrl = class.CardPictureUrl
	}
	return domain, nil
}

func (repo *BookingDetailsRepo) GetByID(ctx context.Context, id int) (bookingdetails.Domain, error) {
	data := Booking_detail{}
	if err := repo.DBConn.Where("id=?", id).First(&data).Error; err != nil {
		return bookingdetails.Domain{}, err
	}
	domain := data.ToDomain()
	type Class struct {
		Name         string
		TimeSchedule string
		GymName      string
		Online       bool
		Link         string
	}
	class := Class{}
	repo.DBConn.Table("booking_details").Select("classes.name as name,classes.card_picture_url as card_picture_url,classes.online as online, classes.link as link, schedules.time_schedule, gyms.name as gym_name").Joins("left join classes on booking_details.class_id=class_id left join schedules on booking_details.schedule_id=schedules.id left join gyms on classes.gym_id=gyms.id").Scan(&class)
	domain.ClassName = class.Name
	domain.TimeSchedule = class.TimeSchedule
	domain.GymName = class.GymName
	domain.Online = class.Online
	domain.Link = class.Link
	return domain, nil
}

func (repo *BookingDetailsRepo) GetByGymID(ctx context.Context, total int, gymID int) ([]bookingdetails.Domain, error) {
	data := []Booking_detail{}
	repo.DBConn.Select("booking_details.*").Joins("LEFT JOIN classes ON booking_details.class_id=classes.id LEFT JOIN gyms ON classes.gym_id=gyms.id").Where("gyms.id=?", gymID).Order("created_at desc").Limit(total).Find(&data)
	domain := []bookingdetails.Domain{}
	for _, val := range data {
		domain = append(domain, val.ToDomain())
	}
	return domain, nil
}
