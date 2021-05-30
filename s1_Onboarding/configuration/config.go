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
	DbUserName           string `json:"DB_USER_NAME"`
	DbPassword           string `json:"DB_PASSWORD"`
	DbDatabaseName       string `json:"DB_DATABASE_NAME"`
	DbHost               string `json:"DB_HOST"`
	DbPort               int    `json:"DB_PORT"`
	CorpMasterDbName     string `json:"CRP_MASTER_DB_NAME"`
	UnvMasterDbName      string `json:"UNV_MASTER_DB_NAME"`
	StuMasterDbName      string `json:"STU_MASTER_DB_NAME"`
	ReferralMasterDbName string `json:"REFERRAL_MASTER"`
	DbRedisAddr          string `json:"DB_REDIS_ADDRESS"`
	DbRedisPort          string `json:"DB_REDIS_PORT"`
	DbRedisPass          string `json:"DB_REDIS_PASS"`
	DbRedisDb            string `json:"DB_REDIS_DB"`
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

// PGService ...
type PGService struct {
	Host string `json:"PAY_SER_HOST"`
	Port int    `json:"PAY_SER_PORT"`
}

// EmailEnv ...
type EmailEnv struct {
	EmailClientID     string `json:"EMAIL_CLIENT_ID"`
	EmailClientSecret string `json:"EMAIL_CLIENT_SECRET"`
	EmailRedirectURL  string `json:"EMAIL_REDIRECT_URL"`
	EmailAccessToken  string `json:"EMAIL_ACCESS_TOKEN"`
	EmailRefreshToken string `json:"EMAIL_REFRESH_TOKEN"`
	EmailTokenType    string `json:"EMAIL_TOKEN_TYPE"`
}

// JwtEnv ...
type JwtEnv struct {
	JwtAccessSecret string `json:"JWT_ACCESS_SECRET"`
}

// TwilioEnv ...
type TwilioEnv struct {
	AccSID    string `json:"TWILIO_ACCOUNT_SID"`
	AccSecret string `json:"TWILIO_ACCOUNT_SECRET"`
	VrfSID    string `json:"TWILIO_VRF_TOKEN"`
}

// PaymentEnv ...
type PaymentEnv struct {
	MerchentID string `json:"RAZORPAY_MERCHENT_ID"`
	KeyID      string `json:"RAZORPAY_KEY_ID"`
	KeySecret  string `json:"RAZORPAY_KEY_SECRET"`
}

var emailConfig *EmailEnv
var dbConfig *DbEnv
var authService *AuthService
var tokenService *TokenService
var pgService *PGService
var jwtConfig *JwtEnv
var twilioConfig *TwilioEnv
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

// PGConfig provides the URL for authenticate micro-service
func PGConfig() PGService {
	return *pgService
}

// TokenConfig provides the URL for authenticate micro-service
func TokenConfig() TokenService {
	return *tokenService
}

// EmailConfig provide configuration details of Email service
func EmailConfig() EmailEnv {
	return *emailConfig
}

// JwtConfig provides configuration for jwt authorization and validation
func JwtConfig() JwtEnv {
	return *jwtConfig
}

// TwilioConfig provides configuration for jwt authorization and validation
func TwilioConfig() TwilioEnv {
	return *twilioConfig
}

// PaymentConfig provides configuration for jwt authorization and validation
func PaymentConfig() PaymentEnv {
	return *paymentConfig
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

	err = json.Unmarshal(dat, &emailConfig)
	//fmt.Printf(" email Config %+v \n ", emailConfig)

	err = json.Unmarshal(dat, &jwtConfig)
	//fmt.Printf(" jwt Config %+v \n ", jwtConfig)

	err = json.Unmarshal(dat, &twilioConfig)

	err = json.Unmarshal(dat, &paymentConfig)

	err = json.Unmarshal(dat, &authService)

	err = json.Unmarshal(dat, &tokenService)

	err = json.Unmarshal(dat, &pgService)

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
