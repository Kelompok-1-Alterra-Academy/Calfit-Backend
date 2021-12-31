package classes

import (
	"CalFit/business/classes"
	"CalFit/controllers"
	responses "CalFit/controllers/classes/response"
	"CalFit/exceptions"

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

func (b *ClassController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	classes, err := b.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.ClassResponse, len(classes))
	for i, gym := range classes {
		response[i] = responses.FromDomain(gym)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (u *ClassController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("classId")
	class, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrClassNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.FromDomain(class)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}
