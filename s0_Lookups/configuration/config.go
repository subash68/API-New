// Package configuration will export all the environment details for the project to consume
package configuration

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	services "github.com/jaswanth-gorripati/PGK/s0_Lookups/services"
)

// DbEnv ...
type DbEnv struct {
	DbUserName              string `json:"DB_USER_NAME"`
	DbPassword              string `json:"DB_PASSWORD"`
	DbDatabaseName          string `json:"DB_DATABASE_NAME"`
	DbHost                  string `json:"DB_HOST"`
	DbPort                  int    `json:"DB_PORT"`
	TokenBalanceDbName      string `json:"TKN_BLNC"`
	TokenAllocatinoDbName   string `json:"TKN_ALLOC"`
	TokenTransactionsDbName string `json:"TKN_TRANSAC"`
}

// AuthService ...
type AuthService struct {
	Host string `json:"AUTH_SER_HOST"`
	Port int    `json:"AUTH_SER_PORT"`
}

var dbConfig *DbEnv
var authService *AuthService

// DecodeBase64ToString Decode the given base64 encoded text to string
func DecodeBase64ToString(str string) (string, error) {
	byteEncodedStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", fmt.Errorf("Error converting base64 encoded text to string  : %w", err)
	}
	return string(byteEncodedStr), nil
}

// DbConfig structure provides configuration details of database
func DbConfig() DbEnv {
	return *dbConfig
}

// AuthConfig provides the URL for authenticate micro-service
func AuthConfig() AuthService {
	return *authService
}

// Config ...
func Config() {
	log := services.Logger
	file, err := os.Open("./configuration/config.json")

	log.Errorf("\nFailed to open the Configuration file %v\n", err)

	defer file.Close()

	dat, err := ioutil.ReadAll(file)

	if err != nil {

		log.Errorf("\nError in reading Configration file : %v \n", err.Error())
	}
	err = json.Unmarshal(dat, &dbConfig)
	//fmt.Printf(" db Config %+v \n ", dbConfig)

	err = json.Unmarshal(dat, &authService)

	if err != nil {
		panic("Error marshalling from Configration file : " + err.Error())
	}

	return
}
