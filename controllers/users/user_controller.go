package users

import (
	"CalFit/business/users"
	"CalFit/controllers"
	"CalFit/controllers/auth/response"
	"CalFit/controllers/users/request"
	"CalFit/exceptions"
	"errors"
	"net/http"
	"strconv"

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

func (controller *Controllers) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	res, err := controller.UsersUC.GetByID(ctx, id)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			return controllers.ErrorResponse(c, http.StatusConflict, exceptions.ErrInvalidCredentials)
		}
		if errors.Is(err, exceptions.ErrValidationFailed) {
			return controllers.ErrorResponse(c, http.StatusConflict, exceptions.ErrValidationFailed)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Update(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUC.Update(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrUserNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}
