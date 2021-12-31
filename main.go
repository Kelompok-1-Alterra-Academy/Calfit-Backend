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
<<<<<<< Updated upstream
	schedulesRepo := schedulesRepo.NewSchedulesRepo(db)
	schedulesUsecase := schedulesUsecase.NewSchedulesUsecase(schedulesRepo)
	schedulesHandler := schedulesHandler.NewHandler(schedulesUsecase)
=======
	schedulesRepo := _schedulesRepo.NewSchedulesRepo(db)
	schedulesUsecase := _schedulesUsecase.NewSchedulesUsecase(schedulesRepo)
	schedulesHandler := _schedulesHandler.NewHandler(schedulesUsecase)
	gymUsecase := _gymUsecase.NewUsecase(_gymDb.NewGymRepository(db), timeoutContext)
	gymHandler := _gymHandler.NewHandler(*gymUsecase)
	classUsecase := _classUsecase.NewUsecase(_classDb.NewClassRepository(db), timeoutContext)
	classHandler := _classHandler.NewHandler(*classUsecase)
>>>>>>> Stashed changes

	routesInit := routes.HandlerList{
		JWTMiddleware:    configJWT.Init(),
		SchedulesHandler: *schedulesHandler,
<<<<<<< Updated upstream
=======
		GymController:    gymHandler,
		ClassController:  classHandler,
>>>>>>> Stashed changes
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
