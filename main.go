package main

import (
	"CalFit/app/middlewares"
	"CalFit/app/routes"
	_addressUsecase "CalFit/business/addresses"
	_gymUsecase "CalFit/business/gyms"
	_schedulesUsecase "CalFit/business/schedules"
	_gymHandler "CalFit/controllers/gyms"
	_schedulesHandler "CalFit/controllers/schedules"
	"CalFit/repository/mysql"
	_addressDb "CalFit/repository/mysql/addresses"
	_gymDb "CalFit/repository/mysql/gyms"
	_schedulesRepo "CalFit/repository/mysql/schedules"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	db := mysql.InitDB()
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("jwt.secret"),
		ExpiresDuration: viper.GetInt("jwt.expired"),
	}

	timeoutContext := time.Duration(viper.GetInt("server.timeout")) * time.Second

	// Schedules initialize
	schedulesRepo := _schedulesRepo.NewSchedulesRepo(db)
	schedulesUsecase := _schedulesUsecase.NewSchedulesUsecase(schedulesRepo)
	schedulesHandler := _schedulesHandler.NewHandler(schedulesUsecase)
	addressUsecase := _addressUsecase.NewUsecase(_addressDb.NewAddressRepository(db), timeoutContext)
	gymUsecase := _gymUsecase.NewUsecase(_gymDb.NewGymRepository(db), timeoutContext)
	gymHandler := _gymHandler.NewHandler(*gymUsecase,*addressUsecase)
	
	routesInit := routes.HandlerList{
		JWTMiddleware: configJWT.Init(),
		SchedulesHandler: *schedulesHandler,
		GymController: gymHandler,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
