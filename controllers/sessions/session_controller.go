package sessions

import (
	"CalFit/business/sessions"
	"CalFit/controllers"
	"CalFit/controllers/sessions/request"
	"CalFit/controllers/sessions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	SessionUC sessions.Usecase
}

func NewControllers(sessionUC sessions.Usecase) *Controllers {
	return &Controllers{
		SessionUC: sessionUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	reqSession := request.Sessions{}
	c.Bind(&reqSession)
	domain := request.ToDomain(reqSession)
	res, err := controller.SessionUC.Insert(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
