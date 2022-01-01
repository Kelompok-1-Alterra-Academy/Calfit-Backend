package sessions

import (
	"CalFit/business/sessions"
	"CalFit/controllers"
	"CalFit/controllers/sessions/request"
	"CalFit/controllers/sessions/response"
	"CalFit/exceptions"
	"net/http"
	"strconv"

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
	req := request.Sessions{}
	c.Bind(&req)
	domain := req.ToDomain()
	res, err := controller.SessionUC.Insert(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (controller *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := controller.SessionUC.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := []response.Sessions{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) GetById(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := controller.SessionUC.GetById(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrSessionNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := response.FromDomain(res)
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Update(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Sessions{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrEmptyInput)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	domain := req.ToDomain()
	domain.Id = id
	res, err := controller.SessionUC.Update(ctx, domain)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrSessionNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(res))
}
