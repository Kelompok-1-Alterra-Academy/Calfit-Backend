package main

import (
	"CalFit/app/middlewares"
	"CalFit/app/routes"
	_classUsecase "CalFit/business/classes"
	_gymUsecase "CalFit/business/gyms"
	_schedulesUsecase "CalFit/business/schedules"
	_sessionsUsecase "CalFit/business/sessions"
	_classController "CalFit/controllers/classes"
	_gymController "CalFit/controllers/gyms"
	_schedulesController "CalFit/controllers/schedules"
	_sessionsController "CalFit/controllers/sessions"
	"CalFit/repository/mysql"
	_classDb "CalFit/repository/mysql/classes"
	_gymDb "CalFit/repository/mysql/gyms"
	_schedulesRepo "CalFit/repository/mysql/schedules"
	_sessionsRepo "CalFit/repository/mysql/sessions"
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
	schedulesController := _schedulesController.NewControllers(schedulesUsecase)
	gymUsecase := _gymUsecase.NewUsecase(_gymDb.NewGymRepository(db), timeoutContext)
	gymHandler := _gymController.NewHandler(*gymUsecase)
	classUsecase := _classUsecase.NewUsecase(_classDb.NewClassRepository(db), _gymDb.NewGymRepository(db), timeoutContext)
	classHandler := _classController.NewHandler(*classUsecase)

	// Sessions initialize
	sessionsRepo := _sessionsRepo.NewSessionsRepo(db)
	sessionsUsecase := _sessionsUsecase.NewSessionsUsecase(sessionsRepo, timeoutContext)
	sessionsController := _sessionsController.NewControllers(sessionsUsecase)

	routesInit := routes.ControllersList{
		JWTMiddleware:       configJWT.Init(),
		SchedulesController: schedulesController,
		GymController:       gymHandler,
		ClassController:     classHandler,
		SessionsController:  sessionsController,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
