package main

import (
	"CalFit/app/middlewares"
	"CalFit/app/routes"
	_membershipsUsecase "CalFit/business/memberships"
	schedulesUsecase "CalFit/business/schedules"
	_membershipsHandler "CalFit/controllers/memberships"
	schedulesHandler "CalFit/controllers/schedules"
	"CalFit/repository/mysql"
	_membershipsRepo "CalFit/repository/mysql/membership_types"
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

	membershipsRepo := _membershipsRepo.NewMembershipsRepo(db)
	membershipsUsecase := _membershipsUsecase.NewMembershipsUsecase(membershipsRepo)
	membershipsHandler := _membershipsHandler.NewHandler(membershipsUsecase)

	routesInit := routes.HandlerList{
		JWTMiddleware:      configJWT.Init(),
		SchedulesHandler:   *schedulesHandler,
		MembershipsHandler: *membershipsHandler,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
