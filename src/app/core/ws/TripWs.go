package ws

import (
	"PackAndGo/src/app/core/exception"
	"PackAndGo/src/app/core/service"
	"PackAndGo/src/app/core/service/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* TYPES DEFINITION */

type TripWs interface {
	Create(context *gin.Context)
	FindById(context *gin.Context)
	FindAll(context *gin.Context)
}

type DefaultTripWs struct {
	tripService service.TripService
}

/* TYPES CONSTRUCTOR */

func NewDefaultTripWs(tripService service.TripService) *DefaultTripWs {
	return &DefaultTripWs{
		tripService: tripService,
	}
}

/* DefaultTripWs METHODS */

func (ws *DefaultTripWs) Create(context *gin.Context) {

	var tripDto *dto.TripDTO
	if err := context.ShouldBindJSON(&tripDto); err != nil {
		ex := exception.BadRequestException("error unmarshalling request json to object", err)
		context.JSON(ex.Code, ex)
		return
	}

	if ex := ws.tripService.Create(context, tripDto); ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusCreated, tripDto)
}

func (ws *DefaultTripWs) FindById(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		ex := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(ex.Code, ex)
		return
	}

	if context.Request.Body != http.NoBody {
		ex := exception.BadRequestException("body not allowed", nil)
		context.JSON(ex.Code, ex)
		return
	}

	tripDto, ex := ws.tripService.FindById(context, id)
	if ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusOK, tripDto)
}

func (ws *DefaultTripWs) FindAll(context *gin.Context) {

	tripsDto, ex := ws.tripService.FindAll(context)
	if ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusOK, tripsDto)
}
