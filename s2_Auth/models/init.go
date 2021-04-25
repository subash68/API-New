// Package models implements the database modelling for the API
package models

import (

	// Blank initializer
	"log"

	"github.com/go-redis/redis"
	// Blank initializer
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jaswanth-gorripati/PGK/s2_Auth/configuration"
)

// Db declaration
//var Db *sql.DB

// RedisClient ...
var RedisClient *redis.Client

// InitDataModel : Initializing the database models
func InitDataModel() {
	// Getting Configuration details
	dbConfig := configuration.DbConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: dbConfig.DbRedisAddr + ":" + dbConfig.DbRedisPort,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Println("Redis client not connected")
		//panic(err)
	} else {
		log.Println("Redis connected !! ")
	}
}
