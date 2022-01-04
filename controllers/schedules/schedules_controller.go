package schedules

import (
	"CalFit/business/schedules"
	"CalFit/controllers"
	"CalFit/controllers/schedules/request"
	"CalFit/controllers/schedules/response"
	"CalFit/exceptions"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	SchedulesUC schedules.Usecase
}

func NewControllers(schedules schedules.Usecase) *Controllers {
	return &Controllers{
		SchedulesUC: schedules,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	res, err := controller.SchedulesUC.Insert(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (controller *Controllers) Get(c echo.Context) error {
	reqSchedule := request.Schedules{}
	domain := request.ToDomain(reqSchedule)
	res, err := controller.SchedulesUC.Get(domain)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrScheduleNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := []response.Schedules{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Update(c echo.Context) error {
	id := c.FormValue("id")
	if strings.TrimSpace(id) == "" {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrMissingId)
	}
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	idInt, _ := strconv.Atoi(id)
	domain.Id = idInt
	res, err := controller.SchedulesUC.Update(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrScheduleNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Delete(c echo.Context) error {
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	res, err := controller.SchedulesUC.Delete(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrScheduleNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
