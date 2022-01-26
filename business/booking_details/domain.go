package bookingdetails

import (
	"context"
	"time"
)

type Domain struct {
	Id                 int
	Amount             int
	Status             string
	UserID             int
	OperationalAdminID int
	PaymentID          int
	ClassID            int
	ClassName          string
	CardPictureUrl     string
	Online             bool
	Link               string
	ScheduleID         int
	TimeSchedule       string
	GymName            string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Usecase interface {
	Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByGymID(ctx context.Context, total int, gymID int) ([]Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, bookingDetails Domain) (Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByGymID(ctx context.Context, total int, gymID int) ([]Domain, error)
}
