// Package models implements the database modelling for the API
package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	// Blank initializer
	"github.com/go-redis/redis"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"

	// Blank initializer
	_ "github.com/go-sql-driver/mysql"
)

// Db declaration
var Db *sql.DB

// RedisClient ...
var RedisClient *redis.Client

// InitDataModel : Initializing the database models
func InitDataModel() {
	// Getting Configuration details
	dbConfig := configuration.DbConfig()

	// Creating Connection String
	con := dbConfig.DbUserName + ":" + dbConfig.DbPassword + "@tcp(" + dbConfig.DbHost + ":" + strconv.Itoa(dbConfig.DbPort) + ")/" + dbConfig.DbDatabaseName + "?charset=utf8&&parseTime=true"
	fmt.Println(con)
	// Declaring Error so that it would not effect the Db declaration in below statement
	var err error
	Db, err = sql.Open("mysql", con)

	// Catch if error occurs
	if err != nil {
		log.Fatalf("Error in connecting to Database  %v ", err.Error())
	} else {
		log.Println("DB connected")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: dbConfig.DbRedisAddr + ":" + dbConfig.DbRedisPort,
	})

	_, err = RedisClient.Ping().Result()
	if err != nil {
		log.Println("Redis client not connected")
		panic(err)
	} else {
		log.Println("Redis connected !! ")
	}
}
