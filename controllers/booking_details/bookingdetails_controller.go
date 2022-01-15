package bookingdetails

import (
	bookingdetails "CalFit/business/booking_details"
	"CalFit/controllers"
	"CalFit/controllers/booking_details/request"
	"CalFit/controllers/booking_details/response"
	"CalFit/exceptions"
	"net/http"

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
