package schedules

import (
	"CalFit/app/presenter"
	"CalFit/app/presenter/schedules/request"
	"CalFit/app/presenter/schedules/response"
	"CalFit/bussiness/schedules"
	"net/http"

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
	resFromDomain := response.FromDomainAdd(res)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
