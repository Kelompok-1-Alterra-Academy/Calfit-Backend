package main

import (
	"CalFit/app/middlewares"
	"CalFit/app/routes"
	schedulesUsecase "CalFit/business/schedules"
	schedulesHandler "CalFit/controllers/schedules"
	"CalFit/repository/mysql"
	schedulesRepo "CalFit/repository/mysql/schedules"
	"log"

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

	// Schedules initialize
	schedulesRepo := schedulesRepo.NewSchedulesRepo(db)
	schedulesUsecase := schedulesUsecase.NewSchedulesUsecase(schedulesRepo)
	schedulesHandler := schedulesHandler.NewHandler(schedulesUsecase)

	routesInit := routes.HandlerList{
		JWTMiddleware:    configJWT.Init(),
		SchedulesHandler: *schedulesHandler,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
