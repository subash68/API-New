package main

import (
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/configuration"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/models"
	routers "github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/routes"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/services"
)

func main() {

	// Loading Environment
	configuration.Config()

	// Initializing Database models
	models.InitDataModel()

	// Loading stored procedures of the database
	models.CreateSP()

	services.ConfigPaymentClient()

	// Starting routes and serving API
	router := routers.InitialzeRoutes()

	router.Run(":8080")
}
