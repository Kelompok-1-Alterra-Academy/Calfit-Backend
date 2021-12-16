package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_USERNAME": viper.GetString("database.username"),
		"DB_PASSWORD": viper.GetString("database.password"),
		"DB_HOST":     viper.GetString("database.host"),
		"DB_PORT":     viper.GetString("database.port"),
		"DB_NAME":     viper.GetString("database.name"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_PASSWORD"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(DB)
	return DB
}
