package routes

import (
	"CalFit/controllers/classes"
	"CalFit/controllers/gyms"
	"CalFit/controllers/schedules"
	"CalFit/controllers/sessions"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllersList struct {
	JWTMiddleware       middleware.JWTConfig
	SchedulesController *schedules.Controllers
	GymController       *gyms.GymController
	ClassController     *classes.ClassController
	SessionsController  *sessions.Controllers
}

func (controllers ControllersList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	v1 := e.Group("/api/v1")
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	// unprotected routes
	{
		// gym endpoint
		v1.GET("/gyms", controllers.GymController.GetAll)
		v1.GET("/gyms/:gymId", controllers.GymController.GetById)

		// class endpoint
		v1.GET("/classes/:classId", controllers.ClassController.GetById)

		// schedules endpoint
		v1.POST("/schedules", controllers.SchedulesController.Insert)
		v1.GET("/schedules", controllers.SchedulesController.Get)

		// session endpoint
		v1.POST("/sessions", controllers.SessionsController.Insert)
		v1.GET("/sessions", controllers.SessionsController.GetAll)
		v1.GET("/sessions/:id", controllers.SessionsController.GetById)

	}

	// superadmin routes
	superadmin := v1.Group("")
	// superadmin.Use(controllers.JWTMiddleware.MiddlewareFunc())
	{
		// gym endpoint
		superadmin.POST("/gyms", controllers.GymController.Create)
		superadmin.PUT("/gyms/:gymId", controllers.GymController.Update)
		superadmin.DELETE("/gyms/:gymId", controllers.GymController.Delete)

		// class endpoint
		superadmin.GET("/classes", controllers.ClassController.GetAll)

		// session endpoint
		superadmin.PUT("/sessions/:id", controllers.SessionsController.Update)
		superadmin.DELETE("/sessions/:id", controllers.SessionsController.Delete)
		superadmin.PUT("/schedules/:id", controllers.SchedulesController.Update)
		superadmin.DELETE("/schedules:/:id", controllers.SchedulesController.Delete)
	}
}
