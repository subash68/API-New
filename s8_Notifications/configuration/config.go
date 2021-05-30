// Package configuration will export all the environment details for the project to consume
package configuration

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// DbEnv ...
type DbEnv struct {
	DbUserName          string `json:"DB_USER_NAME"`
	DbPassword          string `json:"DB_PASSWORD"`
	DbDatabaseName      string `json:"DB_DATABASE_NAME"`
	DbHost              string `json:"DB_HOST"`
	DbPort              int    `json:"DB_PORT"`
	NotificationsDbName string `json:"NOTIFICATIONS"`
}

// AuthService ...
type AuthService struct {
	Host string `json:"AUTH_SER_HOST"`
	Port int    `json:"AUTH_SER_PORT"`
}

// JwtEnv ...
type JwtEnv struct {
	JwtAccessSecret string `json:"JWT_ACCESS_SECRET"`
}

var dbConfig *DbEnv
var authService *AuthService
var jwtConfig *JwtEnv

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

// JwtConfig provides configuration for jwt authorization and validation
func JwtConfig() JwtEnv {
	return *jwtConfig
}

// Config ...
func Config() {
	file, err := os.Open("./configuration/config.json")

	defer file.Close()

	dat, err := ioutil.ReadAll(file)

	if err != nil {

		fmt.Printf("Error in reading Configration file : %v \n", err.Error())
	}
	err = json.Unmarshal(dat, &dbConfig)
	//fmt.Printf(" db Config %+v \n ", dbConfig)

	err = json.Unmarshal(dat, &jwtConfig)
	//fmt.Printf(" jwt Config %+v \n ", jwtConfig)

	err = json.Unmarshal(dat, &authService)

	if err != nil {

		panic("Error marshalling from Configration file : " + err.Error())
	}

	//dbConfig.DbPassword, err = DecodeBase64ToString(dbConfig.DbPassword)
	jwtConfig.JwtAccessSecret, err = DecodeBase64ToString(jwtConfig.JwtAccessSecret)

	if err != nil {

		panic("Error decoding base64 encoded text :" + err.Error())
	}

	return
}
