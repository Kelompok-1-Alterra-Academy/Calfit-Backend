package sessions

import (
	"CalFit/business/sessions"
	presenter "CalFit/controllers"
	"CalFit/controllers/sessions/request"
	"CalFit/controllers/sessions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	SessionUC sessions.Usecase
}

func NewController(sessionUC sessions.Usecase) *Controller {
	return &Controller{
		SessionUC: sessionUC,
	}
}

func (controller *Controller) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	reqSession := request.Sessions{}
	c.Bind(&reqSession)
	domain := request.ToDomain(reqSession)
	res, err := controller.SessionUC.Insert(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
