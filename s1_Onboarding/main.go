package main

import (
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/configuration"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
	routers "github.com/jaswanth-gorripati/PGK/s1_Onboarding/routes"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/services"
)

func main() {

	// Loading Environment
	configuration.Config()

	// Initializing Database models
	models.InitDataModel()

	// Loading stored procedures of the database
	models.CreateSP()

	// Configuring Gmail service
	services.ConfigureOAuthMailService()

	// Configuring Payment service
	services.ConfigPaymentClient()

	// Starting routes and serving API
	router := routers.InitialzeRoutes()

	router.Run("0.0.0.0:8080")
}
