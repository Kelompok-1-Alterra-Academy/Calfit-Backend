package schedules

import (
	"CalFit/business/schedules"
	"CalFit/controllers"
	"CalFit/controllers/schedules/request"
	"CalFit/controllers/schedules/response"
	"CalFit/exceptions"
	"net/http"
	"strconv"

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
	ctx := c.Request().Context()
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	res, err := controller.SchedulesUC.Insert(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (controller *Controllers) Get(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := controller.SchedulesUC.Get(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := []response.Schedules{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Update(c echo.Context) error {
	ctx := c.Request().Context()
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	id, _ := strconv.Atoi(c.Param("id"))
	domain := request.ToDomain(reqSchedule)
	domain.Id = id
	res, err := controller.SchedulesUC.Update(ctx, domain)
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
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := controller.SchedulesUC.Delete(ctx, id)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrScheduleNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
