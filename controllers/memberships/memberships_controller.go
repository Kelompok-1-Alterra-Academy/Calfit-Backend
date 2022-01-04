package memberships

import (
	"CalFit/business/memberships"
	"CalFit/controllers"
	"CalFit/controllers/memberships/request"
	"CalFit/controllers/memberships/response"
	"CalFit/exceptions"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	MembershipsUC memberships.Usecase
}

func NewControllers(memberships memberships.Usecase) *Controllers {
	return &Controllers{
		MembershipsUC: memberships,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	reqMembership_type := request.Memberships{}
	c.Bind(&reqMembership_type)
	domain := request.ToDomain(reqMembership_type)
	res, err := controller.MembershipsUC.Insert(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (controller *Controllers) Get(c echo.Context) error {
	reqMembership_type := request.Memberships{}
	domain := request.ToDomain(reqMembership_type)
	res, err := controller.MembershipsUC.Get(domain)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrInternalServerError)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	resFromDomain := []response.Memberships{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (controller *Controllers) Update(c echo.Context) error {
	id := c.FormValue("id")
	if strings.TrimSpace(id) == "" {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrMissingId)
	}
	reqMembership_type := request.Memberships{}
	domain := request.ToDomain(reqMembership_type)
	idInt, _ := strconv.Atoi(id)
	domain.Id = idInt
	res, err := controller.MembershipsUC.Update(domain)
	resFromDomain := response.FromDomain(res)
	for err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrMembershipNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) Delete(c echo.Context) error {
	reqMembership_type := request.Memberships{}
	c.Bind(&reqMembership_type)
	domain := request.ToDomain(reqMembership_type)
	res, err := controller.MembershipsUC.Update(domain)
	resFromDomain := response.FromDomain(res)
	for err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrMembershipNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
