package main

import (
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/configuration"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/models"
	routers "github.com/jaswanth-gorripati/PGK/s5_Subcription/routes"
)

func main() {

	// Loading Environment
	configuration.Config()

	// Initializing Database models
	models.InitDataModel()

	// Loading stored procedures of the database
	models.CreateSP()

	// // Configuring Gmail service
	// services.ConfigureOAuthMailService()

	// // Configuring Payment service
	// services.ConfigPaymentClient()

	// Starting routes and serving API
	router := routers.InitialzeRoutes()

	router.Run(":8080")
}
