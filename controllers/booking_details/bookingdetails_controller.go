package bookingdetails

import (
	bookingdetails "CalFit/business/booking_details"
	"CalFit/controllers"
	"CalFit/controllers/booking_details/request"
	"CalFit/controllers/booking_details/response"
	"CalFit/exceptions"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	BookingDetailsUC bookingdetails.Usecase
}

func NewControllers(usecase bookingdetails.Usecase) *Controllers {
	return &Controllers{
		BookingDetailsUC: usecase,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Booking_details{}
	if err := c.Bind(&req); err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.BookingDetailsUC.Insert(ctx, domain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := response.FromDomain(res)
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) GetByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	res, err := controller.BookingDetailsUC.GetByUserID(ctx, userID)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := []response.Booking_details{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (controller *Controllers) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	res, err := controller.BookingDetailsUC.GetByID(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(res))
}

func (controller *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	total, err := strconv.Atoi(c.QueryParam("total"))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}
	res, err := controller.BookingDetailsUC.GetAll(ctx, total)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	resFromDomain := []response.Booking_details{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, http.StatusOK, resFromDomain)
}
