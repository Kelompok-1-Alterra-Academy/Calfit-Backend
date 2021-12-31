package sessions

import (
	"CalFit/business/sessions"
	"CalFit/controllers/sessions/request"
	"errors"

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
	c.Bind(reqSession)
	domain := request.ToDomain(reqSession)
	controller.SessionUC.Insert(ctx, domain)

	return errors.New("s")
}
