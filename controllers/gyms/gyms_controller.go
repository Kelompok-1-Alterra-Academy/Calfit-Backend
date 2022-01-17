package gyms

import (
	"CalFit/business/addresses"
	"CalFit/business/gyms"
	"CalFit/business/paginations"
	"CalFit/controllers"
	requests "CalFit/controllers/gyms/request"
	responses "CalFit/controllers/gyms/response"
	"CalFit/exceptions"
	"strconv"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"

	"net/http"

	// "strings"
	"github.com/labstack/echo/v4"
)

type GymController struct {
	Usecase gyms.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewHandler(u gyms.Usecase) *GymController {
	return &GymController{
		Usecase: u,
	}
}

func (g *GymController) GetAll(c echo.Context) error {
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

	gyms, err := g.Usecase.GetAll(ctx, paginationDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.GymResponse, len(gyms))
	for i, gym := range gyms {
		response[i] = responses.FromDomain(gym)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (g *GymController) CountAll(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := g.Usecase.CountAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, count)
}

func (u *GymController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("gymId")
	gym, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrGymNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.FromDomain(gym)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (g *GymController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	createdGym := requests.CreateGym{}
	c.Bind(&createdGym)

	addressDomain := addresses.Domain{
		Address:     createdGym.Address,
		District:    createdGym.District,
		City:        createdGym.City,
		Postal_code: createdGym.Postal_code,
	}

	gymDomain := gyms.Domain{
		Name:                 createdGym.Name,
		Description:          createdGym.Description,
		Telephone:            createdGym.Telephone,
		Picture:              createdGym.Picture,
		Address:              addressDomain,
		Operational_admin_ID: createdGym.Operational_admin_ID,
	}

	gym, err := g.Usecase.Create(ctx, gymDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	gymResponse := responses.FromDomain(gym)

	return controllers.SuccessResponse(c, http.StatusCreated, gymResponse)
}

func (g *GymController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("gymId")
	updatedGym := requests.CreateGym{}
	c.Bind(&updatedGym)

	addressDomain := addresses.Domain{
		Address:     updatedGym.Address,
		District:    updatedGym.District,
		City:        updatedGym.City,
		Postal_code: updatedGym.Postal_code,
	}

	gymDomain := gyms.Domain{
		Name:                 updatedGym.Name,
		Description:          updatedGym.Description,
		Telephone:            updatedGym.Telephone,
		Picture:              updatedGym.Picture,
		Address:              addressDomain,
		Operational_admin_ID: updatedGym.Operational_admin_ID,
	}

	gym, err := g.Usecase.Update(ctx, id, gymDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	gymResponse := responses.FromDomain(gym)

	return controllers.SuccessResponse(c, http.StatusAccepted, gymResponse)
}

func (g *GymController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("gymId")

	err := g.Usecase.Delete(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusAccepted, nil)
}
