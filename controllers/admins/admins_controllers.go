package admins

import (
	"CalFit/business/admins"
	"CalFit/business/paginations"
	"CalFit/controllers"
	requests "CalFit/controllers/admins/request"
	responses "CalFit/controllers/admins/response"
	"CalFit/exceptions"
	"errors"
	"net/http"
	"strconv"

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

func (controller *OpAdminControllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	paginationDomain := paginations.Domain{
		Page:  1,
		Limit: 0,
	}

	// get pagination query
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	sort := c.QueryParam("sort")

	var intPage, intLimit int
	var err error
	if page != "" {
		intPage, err = strconv.Atoi(page)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
		paginationDomain.Page = intPage
	}
	if limit != "" {
		intLimit, err = strconv.Atoi(limit)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
		paginationDomain.Limit = intLimit
	}

	paginationDomain.Sort = sort

	admins, err := controller.OpAdminUC.GetAll(ctx, paginationDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.OpAdmin, len(admins))
	for i, opadmin := range admins {
		response[i] = responses.FromDomainOpAdmin(opadmin)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (controller *OpAdminControllers) CountAll(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := controller.OpAdminUC.CountAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, count)
}
