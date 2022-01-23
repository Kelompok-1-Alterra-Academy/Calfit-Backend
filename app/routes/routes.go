package routes

import (
	"CalFit/app/middlewares"
	"CalFit/controllers/auth"
	bookingdetails "CalFit/controllers/booking_details"
	"CalFit/controllers/classes"
	"CalFit/controllers/gyms"
	"CalFit/controllers/schedules"
	"CalFit/controllers/sessions"

	"CalFit/controllers/memberships"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllersList struct {
	JWTMiddleware            middleware.JWTConfig
	MembershipsController    *memberships.MembershipController
	SchedulesController      *schedules.Controllers
	GymController            *gyms.GymController
	ClassController          *classes.ClassController
	SessionsController       *sessions.Controllers
	AuthController           *auth.Controllers
	BookingDetailsController *bookingdetails.Controllers
}

func (controllers ControllersList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	v1 := e.Group("/api/v1")
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	// unprotected routes
	{
		// gym endpoint
		v1.GET("/gyms", controllers.GymController.GetAll)
		// v1.GET("/gyms/count", controllers.GymController.CountAll)
		v1.GET("/gyms/:gymId", controllers.GymController.GetById)

		// class endpoint
		v1.GET("/classes", controllers.ClassController.GetAll)
		// v1.GET("/classes/count", controllers.ClassController.CountAll)
		v1.GET("/classes/:classId", controllers.ClassController.GetById)
		v1.POST("/classes/:class_id/bookings", controllers.BookingDetailsController.Insert)

		// membership endpoint
		v1.GET("/memberships", controllers.MembershipsController.Get)
		v1.GET("/memberships/:Id", controllers.MembershipsController.GetById)

		// schedules endpoint
		v1.POST("/schedules", controllers.SchedulesController.Insert)
		v1.GET("/schedules", controllers.SchedulesController.Get)
		v1.PUT("/schedules", controllers.SchedulesController.Update)
		v1.DELETE("/schedules", controllers.SchedulesController.Delete)

		// session endpoint
		v1.POST("/sessions", controllers.SessionsController.Insert)
		v1.GET("/sessions", controllers.SessionsController.GetAll)
		v1.GET("/sessions/:id", controllers.SessionsController.GetById)

		v1.POST("/auth/loginOAuth", controllers.AuthController.LoginOAuth)
		v1.POST("/auth/login", controllers.AuthController.Login)
		v1.POST("/auth/register", controllers.AuthController.Register)

		v1.POST("/superadmin/login", controllers.AuthController.SuperadminLogin)
		v1.POST("/superadmin/register", controllers.AuthController.SuperadminRegister)
	}

	// Member routes
	member := v1.Group("", middleware.JWTWithConfig(controllers.JWTMiddleware))
	{
		member.GET("/account/:id/mybookings", controllers.BookingDetailsController.GetByUserID, middlewares.Member())
		member.GET("/bookings/:id", controllers.BookingDetailsController.GetByID, middlewares.Member())
	}

	// admin routes
	admin := v1.Group("", middleware.JWTWithConfig(controllers.JWTMiddleware))
	{
		// gym endpoint
		admin.GET("/gyms/count", controllers.GymController.CountAll, middlewares.OperationalAdmin())
		admin.POST("/gyms", controllers.GymController.Create, middlewares.Superadmin())
		admin.PUT("/gyms/:gymId", controllers.GymController.Update, middlewares.OperationalAdmin())
		admin.DELETE("/gyms/:gymId", controllers.GymController.Delete, middlewares.Superadmin())

		// class endpoint
		admin.GET("/classes/count", controllers.ClassController.CountAll, middlewares.OperationalAdmin())
		admin.POST("/gyms/:gymId/classes", controllers.ClassController.Create, middlewares.OperationalAdmin())
		admin.PUT("/gyms/:gymId/classes/:classId", controllers.ClassController.Update, middlewares.OperationalAdmin())
		admin.DELETE("/gyms/:gymId/classes/:classId", controllers.ClassController.Delete, middlewares.OperationalAdmin())

		//membership endpoint
		admin.POST("/memberships", controllers.MembershipsController.Insert, middlewares.Superadmin())
		admin.PUT("/memberships/:Id", controllers.MembershipsController.Update, middlewares.Superadmin())
		admin.DELETE("/memberships/:Id", controllers.MembershipsController.Delete, middlewares.Superadmin())

		// session endpoint
		admin.PUT("/sessions/:id", controllers.SessionsController.Update, middlewares.Superadmin())
		admin.DELETE("/sessions/:id", controllers.SessionsController.Delete, middlewares.Superadmin())
		admin.PUT("/schedules/:id", controllers.SchedulesController.Update, middlewares.Superadmin())
		admin.DELETE("/schedules:/:id", controllers.SchedulesController.Delete, middlewares.Superadmin())
	}
}
