package schedules

import (
	"CalFit/business/schedules"
	presenter "CalFit/controllers"
	"CalFit/controllers/schedules/request"
	"CalFit/controllers/schedules/response"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	SchedulesUC schedules.Usecase
}

func NewHandler(schedules schedules.Usecase) *Presenter {
	return &Presenter{
		SchedulesUC: schedules,
	}
}

func (p *Presenter) Insert(c echo.Context) error {
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	res, err := p.SchedulesUC.Insert(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (p *Presenter) Get(c echo.Context) error {
	reqSchedule := request.Schedules{}
	domain := request.ToDomain(reqSchedule)
	res, err := p.SchedulesUC.Get(domain)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	resFromDomain := []response.Schedules{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (p *Presenter) Update(c echo.Context) error {
	id := c.FormValue("id")
	if strings.TrimSpace(id) == "" {
		return presenter.ErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	idInt, _ := strconv.Atoi(id)
	domain.Id = idInt
	res, err := p.SchedulesUC.Update(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (p *Presenter) Delete(c echo.Context) error {
	reqSchedule := request.Schedules{}
	c.Bind(&reqSchedule)
	domain := request.ToDomain(reqSchedule)
	res, err := p.SchedulesUC.Delete(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
