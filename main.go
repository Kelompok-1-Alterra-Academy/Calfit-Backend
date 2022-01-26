package main

import (
	"CalFit/app/middlewares"
	"CalFit/app/routes"
	_opadminsUsecase "CalFit/business/admins"
	_bookingDetailsUsecase "CalFit/business/booking_details"
	_classUsecase "CalFit/business/classes"
	_gymUsecase "CalFit/business/gyms"
	_membershipsUsecase "CalFit/business/memberships"
	_newslettersUsecase "CalFit/business/newsletters"
	_schedulesUsecase "CalFit/business/schedules"
	_sessionsUsecase "CalFit/business/sessions"
	_superadminsUsecase "CalFit/business/superadmins"
	_usersUsecase "CalFit/business/users"
	_opadminsController "CalFit/controllers/admins"
	_authController "CalFit/controllers/auth"
	bookingdetails "CalFit/controllers/booking_details"
	_classController "CalFit/controllers/classes"
	_gymController "CalFit/controllers/gyms"
	_membershipsController "CalFit/controllers/memberships"
	_newslettersController "CalFit/controllers/newsletters"
	_schedulesController "CalFit/controllers/schedules"
	_sessionsController "CalFit/controllers/sessions"
	_superadminsController "CalFit/controllers/superadmins"
	_usersController "CalFit/controllers/users"
	"CalFit/repository/mysql"
	_bookingDetailsRepo "CalFit/repository/mysql/booking_details"
	_classDb "CalFit/repository/mysql/classes"
	_gymDb "CalFit/repository/mysql/gyms"
	_membershipsRepo "CalFit/repository/mysql/membership_types"
	_newslettersRepo "CalFit/repository/mysql/newsletters"
	_opadminsRepo "CalFit/repository/mysql/operational_admins"
	_schedulesRepo "CalFit/repository/mysql/schedules"
	_sessionsRepo "CalFit/repository/mysql/sessions"
	_superadminsRepo "CalFit/repository/mysql/superadmins"
	_usersRepo "CalFit/repository/mysql/users"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
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
		SecretJWT:       viper.GetString("JWT_SECRET"),
		ExpiresDuration: viper.GetInt("JWT_EXPIRED"),
	}

	timeoutContext := time.Duration(viper.GetInt("SERVER_TIMEOUT")) * time.Second

	// Schedules initialize
	schedulesRepo := _schedulesRepo.NewSchedulesRepo(db)
	schedulesUsecase := _schedulesUsecase.NewSchedulesUsecase(schedulesRepo, timeoutContext)
	schedulesController := _schedulesController.NewControllers(schedulesUsecase)

	// Gyms initialize
	gymsRepo := _gymDb.NewGymRepository(db)
	gymsUsecase := _gymUsecase.NewUsecase(gymsRepo, timeoutContext)
	gymsHandler := _gymController.NewHandler(*gymsUsecase)

	// Classes initialize
	classesRepo := _classDb.NewClassRepository(db)
	classesUsecase := _classUsecase.NewUsecase(classesRepo, gymsRepo, timeoutContext)
	classesHandler := _classController.NewHandler(*classesUsecase)

	// Sessions initialize
	sessionsRepo := _sessionsRepo.NewSessionsRepo(db)
	sessionsUsecase := _sessionsUsecase.NewSessionsUsecase(sessionsRepo, timeoutContext)
	sessionsController := _sessionsController.NewControllers(sessionsUsecase)

	// Newsletters initialize
	newslettersRepo := _newslettersRepo.NewNewsletterRepository(db)
	newslettersUsecase := _newslettersUsecase.NewNewsletterUsecase(newslettersRepo, timeoutContext)
	newslettersController := _newslettersController.NewHandler(newslettersUsecase)

	// Memberships initialize
	membershipsRepo := _membershipsRepo.NewMembershipsRepo(db)
	membershipsUsecase := _membershipsUsecase.NewMembershipsUsecase(membershipsRepo, timeoutContext)
	membershipsController := _membershipsController.NewHandler(*membershipsUsecase)

	// Profile initialize
	profileRepo := _usersRepo.NewProfileRepo(db)
	profileUsecase := _usersUsecase.NewProfileUsecase(profileRepo, timeoutContext)
	profileController := _usersController.NewHandler(profileUsecase)

	// Users initialize
	usersRepo := _usersRepo.NewUsersRepo(db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepo, timeoutContext, &configJWT)

	// Operationaladmins initialize

	operationaladminsRepo := _opadminsRepo.NewOperationalAdminsRepo(db)
	operatinaladminsUsecase := _opadminsUsecase.NewOperationaldminsUsecase(operationaladminsRepo, timeoutContext, &configJWT)
	operationaladminsController := _opadminsController.NewControllers(operatinaladminsUsecase)

	// Superadmins initialize
	superadminsRepo := _superadminsRepo.NewSuperadminsRepo(db)
	superadminsUsecase := _superadminsUsecase.NewSuperadminsUsecase(superadminsRepo, timeoutContext, &configJWT)
	superadminsController := _superadminsController.NewControllers(superadminsUsecase)

	// Auth initialize
	authController := _authController.NewControllers(usersUsecase, superadminsUsecase, operatinaladminsUsecase)

	// BookingDetails initialize
	bookingDetailsRepo := _bookingDetailsRepo.NewBookingDetailsRepo(db)
	bookingDetailsUsecase := _bookingDetailsUsecase.NewBookingDetailsUsecase(bookingDetailsRepo, timeoutContext)
	bookingDetailsController := bookingdetails.NewControllers(bookingDetailsUsecase)

	// Users initialize
	usersController := _usersController.NewControllers(usersUsecase)

	routesInit := routes.ControllersList{
		JWTMiddleware:               configJWT.Init(),
		SchedulesController:         schedulesController,
		GymController:               gymsHandler,
		MembershipsController:       membershipsController,
		ClassController:             classesHandler,
		SessionsController:          sessionsController,
		AuthController:              authController,
		BookingDetailsController:    bookingDetailsController,
		UsersController:             usersController,
		SuperadminsController:       superadminsController,
		ProfileController:           profileController,
		NewslettersController:       newslettersController,
		OperationaladminsController: operationaladminsController,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("SERVER_PORT")))
}
