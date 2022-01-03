package auth

import (
	"CalFit/business/users"
	"CalFit/controllers"
	"CalFit/controllers/auth/request"
	"CalFit/controllers/auth/response"
	"CalFit/exceptions"
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

func (controller *Controllers) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Auth{}
	c.Bind(&req)
	domain := req.ToDomain()
	res, err := controller.UsersUC.Login(ctx, domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}
