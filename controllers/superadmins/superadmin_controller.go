package superadmins

import (
	"CalFit/business/superadmins"
	"CalFit/controllers"
	"CalFit/controllers/auth/response"
	"CalFit/controllers/superadmins/request"
	"CalFit/exceptions"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	SuperadminsUC superadmins.Usecase
}

func NewControllers(superadminsUC superadmins.Usecase) *Controllers {
	return &Controllers{
		SuperadminsUC: superadminsUC,
	}
}

func (controller *Controllers) UpdatePassword(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.SuperadminAuth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.SuperadminsUC.UpdatePassword(ctx, domain)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			return controllers.ErrorResponse(c, http.StatusConflict, err)
		} else if errors.Is(err, exceptions.ErrSuperadminNotFound) {
			return controllers.ErrorResponse(c, http.StatusNotFound, err)
		} else if errors.Is(err, exceptions.ErrValidationFailed) {
			return controllers.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	superadminResponse := response.FromDomainSuperadmin(res)
	return controllers.SuccessResponse(c, http.StatusOK, superadminResponse)
}
