package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/controller"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	TripRepository := repository.NewTripRepository(dbConnection)
	TripUseCase := usecase.NewTripUsecase(TripRepository)

	TripController := controller.NewTripController(TripUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/trips", TripController.GetTrips)
	server.POST("/trips", TripController.CreateTrip)
	server.GET("/trip/:id", TripController.GetById)

	server.Run(":8000")

}
