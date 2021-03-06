package main

import (
	"os"

	"github.com/jaswanth-gorripati/PGK/s0_Lookups/configuration"
	"github.com/jaswanth-gorripati/PGK/s0_Lookups/models"
	routers "github.com/jaswanth-gorripati/PGK/s0_Lookups/routes"
	services "github.com/jaswanth-gorripati/PGK/s0_Lookups/services"
)

func main() {
	var logLevel string

	if len(os.Args) > 1 {
		logLevel = os.Args[1]
	} else {
		logLevel = "debug"
	}
	services.InitLogger(logLevel)

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
