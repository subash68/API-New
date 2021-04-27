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
	DbUserName       string `json:"DB_USER_NAME"`
	DbPassword       string `json:"DB_PASSWORD"`
	DbDatabaseName   string `json:"DB_DATABASE_NAME"`
	DbHost           string `json:"DB_HOST"`
	DbPort           int    `json:"DB_PORT"`
	CorpMasterDbName string `json:"CORP_MASTER_DB_NAME"`
	UnvMasterDbName  string `json:"UNV_MASTER_DB_NAME"`
	CorpHCDbName     string `json:"CRP_HC_DB"`
	JobHcDbName      string `json:"CRP_JC_DB"`
	JobSkillDbName   string `json:"CRP_JSM_DB"`
	CorpPJDbName     string `json:"CRP_PJ_DB"`
	CorpPDHDbName    string `json:"CRP_PDH_DB"`
	CorpOIDbName     string `json:"CRP_OI_DB"`
	UnvPDHDbName     string `json:"UNV_PDH_DB"`
	UnvOIDbName      string `json:"UNV_OI_DB"`
	LutSkillDbName   string `json:"CRP_SKILL_DB"`
	LutProgramDbName string `json:"CRP_PRG_DB"`
	LutDptDbName     string `json:"CRP_DPT_DB"`
	UnvAccrDbName    string `json:"UNV_ACCR_DB"`
	UnvCoesDbName    string `json:"UNV_COES_DB"`
	UnvProgDbName    string `json:"UNV_PROG_DB"`
	UnvBranchDbName  string `json:"UNV_BRANCH_DB"`
	UnvRankDbName    string `json:"UNV_RANK_DB"`
	UnvSplOfrDbName  string `json:"UNV_SPLOFR_DB"`
	UnvTieupDbName   string `json:"UNV_TIEUP_DB"`
	UnvSubDBName     string `json:"UNV_SUB_DB"`
	CrpSubDBName     string `json:"CRP_SUB_DB"`
	StuSubDBName     string `json:"STU_SUB_DB"`
	StuPublishDBName string `json:"STU_PUBLISH_DB"`
	StuOiDBName      string `json:"STU_OI_DB"`
	DbRedisAddr      string `json:"DB_REDIS_ADDRESS"`
	DbRedisPort      string `json:"DB_REDIS_PORT"`
	DbRedisPass      string `json:"DB_REDIS_PASS"`
	DbRedisDb        string `json:"DB_REDIS_DB"`
}

// AuthService ...
type AuthService struct {
	Host string `json:"AUTH_SER_HOST"`
	Port int    `json:"AUTH_SER_PORT"`
}

// NftService ...
type NftService struct {
	Host string `json:"NFT_SER_HOST"`
	Port int    `json:"NFT_SER_PORT"`
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
var nftService *NftService
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

// NftConfig provides the URL for authenticate micro-service
func NftConfig() NftService {
	return *nftService
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
	file, err := os.Open("./configuration/env")

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

	err = json.Unmarshal(dat, &nftService)

	if err != nil {

		panic("Error marshalling from Configration file : " + err.Error())
	}

	// dbConfig.DbPassword, err = DecodeBase64ToString(dbConfig.DbPassword)
	jwtConfig.JwtAccessSecret, err = DecodeBase64ToString(jwtConfig.JwtAccessSecret)

	if err != nil {

		panic("Error decoding base64 encoded text :" + err.Error())
	}

	return
}
