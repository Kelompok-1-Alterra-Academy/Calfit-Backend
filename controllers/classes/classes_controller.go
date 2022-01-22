package classes

import (
	"CalFit/business/classes"
	"CalFit/business/paginations"
	"CalFit/controllers"
	requests "CalFit/controllers/classes/request"
	responses "CalFit/controllers/classes/response"
	"CalFit/exceptions"
	"strconv"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"

	"net/http"

	// "strings"
	"github.com/labstack/echo/v4"
)

type ClassController struct {
	Usecase classes.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewHandler(u classes.Usecase) *ClassController {
	return &ClassController{
		Usecase: u,
	}
}

func (u *ClassController) GetAll(c echo.Context) error {
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

	classes, err := u.Usecase.GetAll(ctx, paginationDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := make([]responses.ClassResponse, len(classes))
	for i, gym := range classes {
		response[i] = responses.FromDomain(gym)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (u *ClassController) CountAll(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := u.Usecase.CountAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, count)
}

func (u *ClassController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("classId")
	class, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrClassNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.FromDomain(class)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (u *ClassController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	gymId := c.Param("gymId")
	createdClass := requests.CreateClass{}
	c.Bind(&createdClass)

	intGymId, err := strconv.Atoi(gymId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	classDomain := classes.Domain{
		Name:               createdClass.Name,
		Description:        createdClass.Description,
		Banner_picture_url: createdClass.Banner_picture_url,
		Card_picture_url:   createdClass.Card_picture_url,
		Online:             createdClass.Online,
		Link:               createdClass.Link,
		Price:              createdClass.Price,
		Category:           createdClass.Category,
		Status:             createdClass.Status,
		GymID:              uint(intGymId),
		Membership_typeID:  uint(createdClass.Membership_typeID),
	}

	class, err := u.Usecase.Create(ctx, classDomain, gymId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.FromDomain(class)
	return controllers.SuccessResponse(c, http.StatusCreated, response)
}

func (u *ClassController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	classId := c.Param("classId")
	gymId := c.Param("gymId")
	updatedClass := requests.CreateClass{}
	c.Bind(&updatedClass)

	intGymId, err := strconv.Atoi(gymId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	updatedClassDomain := classes.Domain{
		Name:               updatedClass.Name,
		Description:        updatedClass.Description,
		Banner_picture_url: updatedClass.Banner_picture_url,
		Card_picture_url:   updatedClass.Card_picture_url,
		Online:             updatedClass.Online,
		Link:               updatedClass.Link,
		Price:              updatedClass.Price,
		Category:           updatedClass.Category,
		Status:             updatedClass.Status,
		GymID:              uint(intGymId),
		Membership_typeID:  uint(updatedClass.Membership_typeID),
	}

	class, err := u.Usecase.Update(ctx, classId, updatedClassDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.FromDomain(class)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (u *ClassController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("classId")
	err := u.Usecase.Delete(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrClassNotFound)
		}
		if err == exceptions.ErrEmptyInput {
			return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrEmptyInput)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
