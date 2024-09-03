package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/model"
	"rest-api/usecase"
	"strconv"
)

type tripController struct {
	tripUseCase usecase.TripUsecase
}

func NewTripController(usecase usecase.TripUsecase) tripController {
	return tripController{
		tripUseCase: usecase,
	}
}

func (p *tripController) GetTrips(ctx *gin.Context) {

	trips, err := p.tripUseCase.GetTrips()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, trips)
}

func (t *tripController) CreateTrip(ctx *gin.Context) {
	var trip model.Trip

	err := ctx.BindJSON(&trip)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := t.tripUseCase.CreateTrip(trip)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (tc *tripController) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			Message: "id must be informed",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tripId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	trip, err := tc.tripUseCase.GetById(tripId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if trip == nil {
		response := model.Response{
			Message: "trip not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, trip)
}
