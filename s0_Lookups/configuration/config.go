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
	DbUserName                          string `json:"DB_USER_NAME"`
	DbPassword                          string `json:"DB_PASSWORD"`
	DbDatabaseName                      string `json:"DB_DATABASE_NAME"`
	DbHost                              string `json:"DB_HOST"`
	DbPort                              int    `json:"DB_PORT"`
	LUT10Boards                         string `json:"LUT_10thBoardCatalog"`
	LUT12Boards                         string `json:"LUT_12thBoardsCatalog"`
	LUTAccountStatus                    string `json:"LUT_AccountStatus"`
	LUTBranchCatalog                    string `json:"LUT_BranchCatalog"`
	LUTCorporateCategory                string `json:"LUT_CorporateCategory"`
	LUTCorporateIndustry                string `json:"LUT_CorporateIndustry"`
	LUTCorporateType                    string `json:"LUT_CorporateType"`
	LUTJob                              string `json:"LUT_Job"`
	LUTLanguageProficiency              string `json:"LUT_LanguageProficiency"`
	LUTModeOfIssueOfToken               string `json:"LUT_ModeOfIssueOfToken"`
	LUTNotificationType                 string `json:"LUT_NotificationType"`
	LUTPaymentMode                      string `json:"LUT_PaymentMode"`
	LUTProgramCatalog                   string `json:"LUT_Programcatalog"`
	LUTProgramType                      string `json:"LUT_ProgramType"`
	LUTSkillProficiency                 string `json:"LUT_SKillProficiency"`
	LUTSkillsMaster                     string `json:"LUT_Skills_Master"`
	LUTSortBy                           string `json:"LUT_SortBy"`
	LUTStakeholderType                  string `json:"LUT_StakeholderType"`
	LUTStudentEventResult               string `json:"LUT_StudentEventResult"`
	LUTStudentEventType                 string `json:"LUT_StudentEventType"`
	LUTStudentProfileVerificationStatus string `json:"LUT_StudentProfileVerificationStatus"`
	LUTStudentVerificationType          string `json:"LUT_StudentVerificationType"`
	LUTSubscriptionType                 string `json:"LUT_SubscriptionType"`
	LUTTokenEventsDefinition            string `json:"LUT_TokenEventsDefinition"`
	LUTUniversityAccreditationType      string `json:"LUT_UniversityAccreditationType"`
	LUTUniversityCOEType                string `json:"LUT_UniversityCOEType"`
	LUTUniversityCatalog                string `json:"LUT_University_Catalog"`
	LUTUniversitySplOffersType          string `json:"LUT_UniversitySplOffersType"`
	LUTUniversityTieUpType              string `json:"LUT_UniversityTieUpType"`
	LUTUniversityType                   string `json:"LUT_UniversityType"`
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
	if err != nil {
		log.Errorf("\nFailed to open the Configuration file %v\n", err)
	}

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
