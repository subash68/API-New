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
	DbUserName     string `json:"DB_USER_NAME"`
	DbPassword     string `json:"DB_PASSWORD"`
	DbDatabaseName string `json:"DB_DATABASE_NAME"`
	DbHost         string `json:"DB_HOST"`
	DbPort         int    `json:"DB_PORT"`
	CrpPayDbName   string `json:"CRP_PAYMENT"`
	UnvPayDbName   string `json:"UNV_PAYMENT"`
	StuPayDbName   string `json:"STU_PAYMENT"`
}

// AuthService ...
type AuthService struct {
	Host string `json:"AUTH_SER_HOST"`
	Port int    `json:"AUTH_SER_PORT"`
}

// TokenService ...
type TokenService struct {
	Host string `json:"TOKEN_SER_HOST"`
	Port int    `json:"TOKEN_SER_PORT"`
}

// OnboardService ...
type OnboardService struct {
	Host string `json:"ONBOARD_SER_HOST"`
	Port int    `json:"ONBOARD_SER_PORT"`
}

// PaymentEnv ...
type PaymentEnv struct {
	MerchentID string `json:"RAZORPAY_MERCHENT_ID"`
	KeyID      string `json:"RAZORPAY_KEY_ID"`
	KeySecret  string `json:"RAZORPAY_KEY_SECRET"`
}

var dbConfig *DbEnv
var authService *AuthService
var tokenService *TokenService
var onboardService *OnboardService
var paymentConfig *PaymentEnv

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

// TokenConfig provides the URL for authenticate micro-service
func TokenConfig() TokenService {
	return *tokenService
}

// OnboardConfig provides the URL for authenticate micro-service
func OnboardConfig() OnboardService {
	return *onboardService
}

// PaymentConfig provides configuration for jwt authorization and validation
func PaymentConfig() PaymentEnv {
	return *paymentConfig
}

// Config ...
func Config() {
	file, err := os.Open("./configuration/env")

	defer file.Close()

	dat, err := ioutil.ReadAll(file)

	if err != nil {

		fmt.Printf("Error in reading Configration file : %v \n", err.Error())
	}
	err = json.Unmarshal(dat, &dbConfig)
	//fmt.Printf(" db Config %+v \n ", dbConfig)

	err = json.Unmarshal(dat, &paymentConfig)

	err = json.Unmarshal(dat, &authService)

	err = json.Unmarshal(dat, &tokenService)

	err = json.Unmarshal(dat, &onboardService)

	if err != nil {

		panic("Error marshalling from Configration file : " + err.Error())
	}

	//dbConfig.DbPassword, err = DecodeBase64ToString(dbConfig.DbPassword)

	if err != nil {

		panic("Error decoding base64 encoded text :" + err.Error())
	}

	return
}
