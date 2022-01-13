package mysql

import (
	"CalFit/repository/mysql/addresses"
	bookingdetails "CalFit/repository/mysql/booking_details"
	"CalFit/repository/mysql/classes"
	"CalFit/repository/mysql/gyms"
	"CalFit/repository/mysql/membership_types"
	"CalFit/repository/mysql/newsletters"
	"CalFit/repository/mysql/operational_admins"
	"CalFit/repository/mysql/payments"
	"CalFit/repository/mysql/schedules"
	"CalFit/repository/mysql/sessions"
	"CalFit/repository/mysql/super_admins"
	"CalFit/repository/mysql/users"
	"CalFit/repository/mysql/video_contents"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_USERNAME": viper.GetString("DB_USERNAME"),
		"DB_PASSWORD": viper.GetString("DB_PASSWORD"),
		"DB_HOST":     viper.GetString("DB_HOST"),
		"DB_PORT":     viper.GetString("DB_PORT"),
		"DB_NAME":     viper.GetString("DB_NAME"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_PASSWORD"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&super_admins.Super_admin{},
		&operational_admins.Operational_admin{},
		&addresses.Address{},
		&gyms.Gym{},
		&sessions.Session{},
		&schedules.Schedule{},
		&classes.Class{},
<<<<<<< HEAD
		&video_contents.Video_content{},
		&bookingdetails.Booking_detail{},
=======
		&users.User{},
		&booking_details.Booking_detail{},
		&membership_types.Membership_type{},
		&payments.Payment{},
		&video_contents.Video_content{},
		&newsletters.Newsletter{},
>>>>>>> d5b800f... fix: foreign key error when migrating tables
	)
	return DB
}
