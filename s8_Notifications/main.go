package main

import (
	"github.com/jaswanth-gorripati/PGK/s8_Notifications/configuration"
	"github.com/jaswanth-gorripati/PGK/s8_Notifications/models"
	routers "github.com/jaswanth-gorripati/PGK/s8_Notifications/routes"
)

func main() {

	// Loading Environment
	configuration.Config()

	// Initializing Database models
	models.InitDataModel()

	// Loading stored procedures of the database
	models.CreateSP()

	// Starting routes and serving API
	router := routers.InitialzeRoutes()

	router.Run(":8080")
}
