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

func (repo *BookingDetailsRepo) GetByUserID(ctx context.Context, userID int) ([]bookingdetails.Domain, error) {
	data := []Booking_detail{}
	if err := repo.DBConn.Where("user_id=?", userID).Find(&data).Error; err != nil {
		return []bookingdetails.Domain{}, err
	}
	var domain []bookingdetails.Domain
	type Class struct {
		Name         string
		TimeSchedule string
		GymName      string
	}
	class := Class{}
	for _, val := range data {
		domain = append(domain, val.ToDomain())
	}
	for i := range domain {
		repo.DBConn.Table("booking_details").Select("classes.name as name, schedules.time_schedule, gyms.name as gym_name").Joins("left join classes on booking_details.class_id=class_id left join schedules on booking_details.schedule_id=schedules.id left join gyms on classes.gym_id=gyms.id").Scan(&class)
		domain[i].ClassName = class.Name
		domain[i].TimeSchedule = class.TimeSchedule
		domain[i].GymName = class.GymName
	}
	return domain, nil
}
