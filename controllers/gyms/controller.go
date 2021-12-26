package gyms

import (
	"CalFit/app/presenter"
	requests "CalFit/app/presenter/gyms/request"
	responses "CalFit/app/presenter/gyms/response"
	"CalFit/business/addresses"
	"CalFit/business/gyms"
	"CalFit/exceptions"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"

	"net/http"

	// "strings"
	"github.com/labstack/echo/v4"
)

type GymController struct {
	Usecase gyms.Usecase
	AddressUsecase addresses.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewGymController(u gyms.Usecase, a addresses.Usecase) *GymController {
	return &GymController{
		Usecase: u,
		AddressUsecase: a,
	}
}

func (b *GymController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	gyms, err := b.Usecase.GetAll(ctx)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.GymResponse, len(gyms))
	for i, gym := range gyms {
		response[i] = responses.FromDomain(gym)
	}
	return presenter.SuccessResponse(c, response)
}

func (u *GymController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("gymId")
	gym, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	
	response := responses.FromDomain(gym)

	return presenter.SuccessResponse(c, response)
}
		
func (b *GymController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	
	createdGym := requests.CreateGym{}
	c.Bind(&createdGym)
	
	gymDomain := gyms.Domain{
		Name: createdGym.Name,
		Telephone: createdGym.Telephone,
		Picture: createdGym.Picture,
		Address: createdGym.Address,
		Operational_admin_ID: createdGym.Operational_admin_ID,
	}
	
	gym, err := b.Usecase.Create(ctx, gymDomain)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	gymResponse := responses.FromDomain(gym)
	
	return presenter.SuccessResponse(c, gymResponse)
	// return presenter.SuccessResponse(c, http.StatusOK)
}

func (b *GymController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("gymId")
	updatedGym := requests.CreateGym{}
	c.Bind(&updatedGym)


	gymDomain := gyms.Domain{
		Name: updatedGym.Name,
		Telephone: updatedGym.Telephone,
		Picture: updatedGym.Picture,
		Address: updatedGym.Address,
		Operational_admin_ID: updatedGym.Operational_admin_ID,
	}

	gym, err := b.Usecase.Update(ctx, id, gymDomain)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	gymResponse := responses.FromDomain(gym)

	return presenter.SuccessResponse(c, gymResponse)
}