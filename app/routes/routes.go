package routes

import (
	"CalFit/controllers/classes"
	"CalFit/controllers/gyms"
	"CalFit/controllers/schedules"
	"CalFit/controllers/sessions"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware      middleware.JWTConfig
	SchedulesHandler   *schedules.Presenter
	GymController      *gyms.GymController
	ClassController    *classes.ClassController
	SessionsController *sessions.Controller
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

	s := e.Group("api/v1/schedules")
	s.POST("", handler.SchedulesHandler.Insert)
	s.GET("", handler.SchedulesHandler.Get)
	s.PUT("", handler.SchedulesHandler.Update)
	s.DELETE("", handler.SchedulesHandler.Delete)

	// unprotected routes
	{
		// gym endpoint
		v1.GET("/gyms", handler.GymController.GetAll)
		v1.GET("/gyms/:gymId", handler.GymController.GetById)

		// class endpoint
		v1.GET("/classes/:classId", handler.ClassController.GetById)
	}

	// superadmin routes
	superadmin := v1.Group("")
	// superadmin.Use(handler.JWTMiddleware.MiddlewareFunc())
	{
		// gym endpoint
		superadmin.POST("/gyms", handler.GymController.Create)
		superadmin.PUT("/gyms/:gymId", handler.GymController.Update)
		superadmin.DELETE("/gyms/:gymId", handler.GymController.Delete)

		// class endpoint
		superadmin.GET("/classes", handler.ClassController.GetAll)
	}
	session := e.Group("/v1/sessions")
	session.POST("", handler.SessionsController.Insert)
}
