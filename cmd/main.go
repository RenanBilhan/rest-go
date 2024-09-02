package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/controller"
	"rest-api/db"
	"rest-api/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	TripUseCase := usecase.NewTripUsecase()

	TripController := controller.NewTripController(TripUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/trips", TripController.GetTrips)

	server.Run(":8000")

}
