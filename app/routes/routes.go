package routes

import (
	"CalFit/controllers/schedules"

	"CalFit/controllers/memberships"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware      middleware.JWTConfig
	SchedulesHandler   schedules.Presenter
	MembershipsHandler memberships.Presenter
}

func (handler HandlerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	s := e.Group("/v1/schedules")
	s.POST("", handler.SchedulesHandler.Insert)
	s.GET("", handler.SchedulesHandler.Get)
	s.PUT("", handler.SchedulesHandler.Update)
	s.DELETE("", handler.SchedulesHandler.Delete)

	m := e.Group("/v1/memberships")
	m.POST("", handler.MembershipsHandler.Insert)
	m.GET("", handler.MembershipsHandler.Get)
	m.PUT("", handler.MembershipsHandler.Update)
	m.DELETE("", handler.MembershipsHandler.Delete)
}
