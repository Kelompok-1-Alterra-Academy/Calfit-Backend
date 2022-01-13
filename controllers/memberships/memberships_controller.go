package memberships

import (
	"CalFit/business/memberships"
	"CalFit/controllers"
	requests "CalFit/controllers/memberships/request"
	responses "CalFit/controllers/memberships/response"
	"CalFit/exceptions"
	"net/http"

	// "strconv"
	// "strings"

	"github.com/labstack/echo/v4"
)

type MembershipController struct {
	Usecase memberships.MembershipsUsecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewHandler(u memberships.MembershipsUsecase) *MembershipController {
	return &MembershipController{
		Usecase: u,
	}
}

func (m *MembershipController) Insert(c echo.Context) error {
	ctx := c.Request().Context()

	createdMembership := requests.Memberships{}
	c.Bind(&createdMembership)

	membershipDomain := memberships.Domain{
		Name:        createdMembership.Name,
		Description: createdMembership.Description,
	}
	membership, err := m.Usecase.Insert(ctx, membershipDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	membershipResponse := responses.FromDomain(membership)

	return controllers.SuccessResponse(c, http.StatusAccepted, membershipResponse)
}

func (m *MembershipController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	membership, err := m.Usecase.Get(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := []responses.Memberships{}
	for _, val := range membership {
		resFromDomain = append(resFromDomain, responses.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (u *MembershipController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	membership, err := u.Usecase.GetById(ctx, id)
	resFromDomain := responses.FromDomain(membership)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrMembershipNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (m *MembershipController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	updatedMembership := requests.Memberships{}
	c.Bind(&updatedMembership)

	membershipDomain := memberships.Domain{
		Name:        updatedMembership.Name,
		Description: updatedMembership.Description,
	}
	membership, err := m.Usecase.Update(ctx, id, membershipDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	membershipResponse := responses.FromDomain(membership)
	return controllers.SuccessResponse(c, http.StatusAccepted, membershipResponse)
}

func (m *MembershipController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")

	err := m.Usecase.Delete(ctx, id)
	for err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.SuccessResponse(c, http.StatusAccepted, nil)
}
