package main

import (
	"CalFit/app/middlewares"
	"CalFit/app/routes"
	_gymUsecase "CalFit/business/gyms"
	_gymController "CalFit/controllers/gyms"
	"CalFit/repository/mysql"
	_gymDb "CalFit/repository/mysql/gyms"
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
	Conn:=mysql.InitDB()
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("jwt.secret"),
		ExpiresDuration: viper.GetInt("jwt.expired"),
	}

	timeoutContext := time.Duration(viper.GetInt("server.timeout")) * time.Second

	gymUsecase := _gymUsecase.NewUsecase(_gymDb.NewGymRepository(Conn), timeoutContext)
	gymController := _gymController.NewGymController(*gymUsecase)

	routesInit := routes.HandlerList{
		JWTMiddleware: configJWT.Init(),
		GymController: gymController,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
