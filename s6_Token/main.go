package main

import (
	"github.com/jaswanth-gorripati/PGK/s6_Token/configuration"
	"github.com/jaswanth-gorripati/PGK/s6_Token/models"
	routers "github.com/jaswanth-gorripati/PGK/s6_Token/routes"
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
