package routes

import (
	"CalFit/controllers/schedules"
	"CalFit/controllers/gyms"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware    middleware.JWTConfig
	SchedulesHandler schedules.Presenter
	GymController *gyms.GymController
}

func (handler HandlerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	v1 := e.Group("/api/v1")
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

<<<<<<< HEAD
	s := e.Group("/v1/schedules")
	s.POST("", handler.SchedulesHandler.Insert)
	s.GET("", handler.SchedulesHandler.Get)
	s.PUT("", handler.SchedulesHandler.Update)
	s.DELETE("", handler.SchedulesHandler.Delete)
=======
	// unprotected routes
	{
		v1.GET("/gyms", handler.GymController.GetAll)
	}
>>>>>>> e0b96f2... feat: add request & response for gyms endpoints
}
