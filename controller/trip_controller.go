package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/model"
	"rest-api/usecase"
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

	trips := []model.Trip{
		{
			TripId:      1,
			TripName:    "Malta",
			Description: "Turismo historico",
		},
	}

	ctx.JSON(http.StatusOK, trips)
}
