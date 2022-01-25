package users

import (
	"CalFit/business/addresses"
	"CalFit/business/paginations"
	"CalFit/business/users"
	"CalFit/controllers"
	requests "CalFit/controllers/users/request"
	responses "CalFit/controllers/users/response"
	"CalFit/exceptions"
	"net/http"

	"strconv"
	// "strings"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	Usecase users.PUseCase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewHandler(u users.PUseCase) *ProfileController {
	return &ProfileController{
		Usecase: u,
	}
}

func (g *ProfileController) GetAll(c echo.Context) error {
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

	users, err := g.Usecase.GetAll(ctx, paginationDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.UsersResponse, len(users))
	for i, user := range users {
		response[i] = responses.FromDomain(user)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (g *ProfileController) CountAll(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := g.Usecase.CountAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, count)
}

func (u *ProfileController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	user, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrGymNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.FromDomain(user)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (g *ProfileController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	updatedProfile := requests.Users{}
	c.Bind(&updatedProfile)

	addressDomain := addresses.Domain{
		Address:     updatedProfile.Address,
		District:    updatedProfile.District,
		City:        updatedProfile.City,
		Postal_code: updatedProfile.Postal_code,
	}

	userDomain := users.Domain{
		Email:            updatedProfile.Email,
		Password:         updatedProfile.Password,
		Address:          addressDomain,
		Photo:            updatedProfile.Photo,
		MembershipTypeID: updatedProfile.MembershipTypeID,
		FullName:         updatedProfile.FullName,
	}

	user, err := g.Usecase.Update(ctx, id, userDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	userResponse := responses.FromDomain(user)

	return controllers.SuccessResponse(c, http.StatusAccepted, userResponse)
}
