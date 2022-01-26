package auth

import (
	"CalFit/business/superadmins"
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
	UsersUC       users.Usecase
	SuperadminsUC superadmins.Usecase
}

func NewControllers(usersUC users.Usecase, superadminsUC superadmins.Usecase) *Controllers {
	return &Controllers{
		UsersUC:       usersUC,
		SuperadminsUC: superadminsUC,
	}
}

func (controller *Controllers) LoginOAuth(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Auth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUC.LoginOAuth(ctx, domain)
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

func (controller *Controllers) SuperadminRegister(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.SuperadminAuth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.SuperadminsUC.Register(ctx, domain)
	if err != nil {
		if errors.Is(err, exceptions.ErrSuperadminExists) {
			return controllers.ErrorResponse(c, http.StatusConflict, err)
		} else if errors.Is(err, exceptions.ErrInvalidCredentials) {
			return controllers.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomainSuperadmin(res))
}

func (controller *Controllers) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Auth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUC.Login(ctx, domain)
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
	cookie := helpers.CreateCookie(resFromDomain.Token)
	c.SetCookie(cookie)
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) SuperadminLogin(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.SuperadminAuth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.SuperadminsUC.Login(ctx, domain)
	resFromDomain := response.FromDomainSuperadmin(res)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			return controllers.ErrorResponse(c, http.StatusConflict, exceptions.ErrInvalidCredentials)
		}
		if errors.Is(err, exceptions.ErrValidationFailed) {
			return controllers.ErrorResponse(c, http.StatusConflict, exceptions.ErrValidationFailed)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	cookie := helpers.CreateCookie(resFromDomain.Token)
	c.SetCookie(cookie)
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) SuperadminLogout(c echo.Context) error {
	cookie := helpers.DeleteCookie()
	c.SetCookie(cookie)
	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
