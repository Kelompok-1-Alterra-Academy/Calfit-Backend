package memberships

import (
	"CalFit/business/memberships"
	presenter "CalFit/controllers"
	"CalFit/controllers/memberships/request"
	"CalFit/controllers/memberships/response"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	MembershipsUC memberships.Usecase
}

func NewHandler(memberships memberships.Usecase) *Presenter {
	return &Presenter{
		MembershipsUC: memberships,
	}
}

func (p *Presenter) Insert(c echo.Context) error {
	reqMembership_type := request.Memberships{}
	c.Bind(&reqMembership_type)
	domain := request.ToDomain(reqMembership_type)
	res, err := p.MembershipsUC.Insert(domain)
	resFromDomain := response.FromDomain(res)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (p *Presenter) Get(c echo.Context) error {
	reqMembership_type := request.Memberships{}
	domain := request.ToDomain(reqMembership_type)
	res, err := p.MembershipsUC.Get(domain)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	resFromDomain := []response.Memberships{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (p *Presenter) Update(c echo.Context) error {
	id := c.FormValue("id")
	if strings.TrimSpace(id) == "" {
		return presenter.ErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}
	reqMembership_type := request.Memberships{}
	domain := request.ToDomain(reqMembership_type)
	idInt, _ := strconv.Atoi(id)
	domain.Id = idInt
	res, err := p.MembershipsUC.Update(domain)
	resFromDomain := response.FromDomain(res)
	for err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (p *Presenter) Delete(c echo.Context) error {
	reqMembership_type := request.Memberships{}
	c.Bind(&reqMembership_type)
	domain := request.ToDomain(reqMembership_type)
	res, err := p.MembershipsUC.Update(domain)
	resFromDomain := response.FromDomain(res)
	for err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
