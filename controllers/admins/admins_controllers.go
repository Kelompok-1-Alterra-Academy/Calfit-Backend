package admins

import (
	"CalFit/business/admins"
	"CalFit/controllers"
	requests "CalFit/controllers/admins/request"
	responses "CalFit/controllers/admins/response"
	"CalFit/exceptions"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OpAdminControllers struct {
	OpAdminUC admins.Usecase
}

func NewControllers(opAdminsUC admins.Usecase) *OpAdminControllers {
	return &OpAdminControllers{
		OpAdminUC: opAdminsUC,
	}
}

func (controller *OpAdminControllers) UpdatePassword(c echo.Context) error {
	ctx := c.Request().Context()
	req := requests.OpAdminAuth{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.OpAdminUC.UpdatePassword(ctx, domain)
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
	superadminResponse := responses.FromDomainOpAdmin(res)
	return controllers.SuccessResponse(c, http.StatusOK, superadminResponse)
}
