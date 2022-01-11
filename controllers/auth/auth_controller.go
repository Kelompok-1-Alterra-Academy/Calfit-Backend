package auth

import (
	"CalFit/business/users"
	"CalFit/controllers"
	"CalFit/controllers/auth/request"
	"CalFit/controllers/auth/response"
	"CalFit/exceptions"
	"CalFit/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	UsersUC users.Usecase
}

func NewControllers(usersUC users.Usecase) *Controllers {
	return &Controllers{
		UsersUC: usersUC,
	}
}

func (controller *Controllers) LoginOauth(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Auth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUC.LoginOauth(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			return controllers.ErrorResponse(c, http.StatusConflict, exceptions.ErrInvalidCredentials)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	cookie := helpers.CreateCookie(resFromDomain.Token)
	c.SetCookie(cookie)
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Auth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUC.Register(ctx, domain)
	if err != nil {
		if errors.Is(err, exceptions.ErrUserAlreadyExists) {
			return controllers.ErrorResponse(c, http.StatusConflict, exceptions.ErrUserAlreadyExists)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(res))
}
