package models

import (
	"database/sql"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/jaswanth-gorripati/PGK/s5_Subcription/configuration"
)

// CompleteDB ...
type CompleteDB struct {
	CorporateMasterDB
	UniversityMasterDb
	StudentMasterDb
}

// CheckPing Checks if connection exists with Database
func CheckPing(customError *DbModelError) {
	err := Db.Ping()
	if err != nil {
		customError.Err = fmt.Errorf("Error While connecting CORPORATE Table %w ", err)
		customError.ErrCode = "S1AUT912"
		customError.ErrTyp = "500"
		customError.SuccessResp = map[string]string{}
		fmt.Printf(" line 21 %+v ", customError)
	}
	return
}

// UpdateProfileData ...
func UpdateProfileData(updateQuery url.Values, spName string, spExt string, stakeholder string, attachmentUpdate bool, attachment []byte) DbModelError {
	updateString, _ := RetriveSP(spName)
	values := []interface{}{}
	var customError DbModelError
	for key, val := range updateQuery {
		if key != "skill" {
			dbKey, exists := GetDbKey(key)
			if exists {
				// customError.ErrTyp = "500"
				// customError.ErrCode = "S3UPDT001"
				// customError.Err = fmt.Errorf("Invalid key " + key + ", Cannot update")
				// return customError

				updateString = updateString + " " + dbKey + "= ?,"
				fmt.Println(updateString)
				values = append(values, val[0])
			}
		}
	}
	if attachmentUpdate {
		updateString = updateString + " Attachment= ?,"
		values = append(values, attachment)
	}
	valLength := reflect.ValueOf(values).Len()
	if valLength == 0 {
		customError.ErrTyp = "000"
		return customError
	}
	values = append(values, stakeholder)
	updateString = updateString[0 : len(updateString)-1]
	whereCond, _ := RetriveSP(spExt)
	updateString = updateString + "" + whereCond

	fmt.Printf("\n keys : %v \nvalues %v\n", updateString, values)

	updateStm, err := Db.Prepare(updateString)
	if err != nil {
		fmt.Println(updateString)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare database Prepare due to %v", err.Error())
		return customError
	}
	_, err = updateStm.Exec(values...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to update the database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		return customError
	}
	customError.ErrTyp = "000"
	return customError
}

// GetProfile ...
func GetProfile(stakeholderID string, getSP string) (*sql.Row, DbModelError) {
	queryCmd, _ := RetriveSP(getSP)
	var customError DbModelError

	row := Db.QueryRow(queryCmd, stakeholderID)

	customError.ErrTyp = "000"

	return row, customError
}

// UpdateProfilePic ...
func UpdateProfilePic(pic []byte, userID string, sp string) DbModelError {
	updatePP, _ := RetriveSP(sp)
	var customError DbModelError
	updateStm, err := Db.Prepare(updatePP)
	if err != nil {
		fmt.Println(updatePP)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare database Prepare due to %v", err.Error())
		return customError
	}
	_, err = updateStm.Exec(userID, pic, pic)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to update the database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		return customError
	}
	customError.ErrTyp = "000"
	return customError

}

// GetProfilePic ...
func GetProfilePic(stakeholderID string, sp string) ([]byte, DbModelError) {
	queryPP, _ := RetriveSP(sp)
	var customError DbModelError
	var ppic []byte
	err := Db.QueryRow(queryPP, stakeholderID).Scan(&ppic)
	if err != nil {
		customError.ErrTyp = "S3PRF001"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve  Profile information : %v", err.Error())
		return ppic, customError
	}
	customError.ErrTyp = "000"
	return ppic, customError
}

// CreateSudID ... UNV_INSIGHTS_Get_Last_ID, SUBUI
func CreateSudID(ID string, query string, code string) (string, error) {
	rowSP, _ := RetriveSP(query)
	lastID := ""
	err := Db.QueryRow(rowSP, ID).Scan(&lastID)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if err == sql.ErrNoRows {
		lastID = "000000000000000"
	}
	fmt.Println("--------------------> ", lastID, "-----", ID)
	corporateNum, _ := strconv.Atoi(ID[7:])
	countNum, _ := strconv.Atoi(lastID[len(lastID)-7:])

	return (code + strconv.Itoa(corporateNum) + (fmt.Sprintf("%07d", (countNum + 1)))), nil
}

// GetUnvDetailsByID ...
func GetUnvDetailsByID(ID string) UnvCDDataModel {
	sp, _ := RetriveSP("UNV_GET_PROFILE_BY_ID")
	var ud UnvCDDataModel
	err := Db.QueryRow(sp, ID).Scan(&ud.Name, &ud.Location, &ud.YearOfEst, &ud.Programs, &ud.Ranking, &ud.Accredations)
	if err != nil {
		fmt.Printf("\n=========Error getting University details ========= %v\n", err)
	}

	fmt.Printf("\n\n Got UD --> %v", ud)
	return ud
}

// GetEmailsForCH ...
func GetEmailsForCH(CID string, UID string) (string, string) {
	sp, _ := RetriveSP("GET_CI_EMAILS")
	var ce, ue string
	err := Db.QueryRow(sp, CID, UID).Scan(&ce, &ue)
	if err != nil {
		fmt.Printf("\n=========Error getting Emails ========= %v\n", err)
	}

	return ce, ue
}

// GetCorpDetailsByID ...
func GetCorpDetailsByID(ID string) CorpCDDataModel {
	sp, _ := RetriveSP("CORP_GET_PROFILE_BY_ID")
	var ud CorpCDDataModel
	err := Db.QueryRow(sp, ID).Scan(&ud.Name, &ud.CIN, &ud.Location, &ud.Category)
	if err != nil {
		fmt.Printf("\n=========Error getting Corporate details ========= %v\n", err)
	}
	ud.CorporateID = ID
	fmt.Printf("\n\n Got UD --> %v", ud)
	return ud
}

func parseSubscriptionType(ct string) string {
	switch ct {
	case "CProfile", "CProfile details":
		return "CP"
		break
	case "UProfile", "UProfile details":
		return "UP"
		break
	case "UOther Information", "Uother Information":
		return "UO"
		break
	case "COther Information", "Cother Information":
		return "CO"
		break
	case "Student Database":
		return "SD"
		break
	case "University Information":
		return "UI"
		break
	case "Campus Hiring":
		return "CR"
		break
	case "Hiring Insights":
		return "HI"
		break
	case "CNew Job":
		return "CJ"
		break
	case "CHiring Criteria":
		return "CHC"
		break
	default:
		return "O"
	}
	return "O"
}

// GetSubTypeFromPublishID ...
func GetSubTypeFromPublishID(publishID string, userType string) (string, string) {
	dbName := ""
	dbConfig := configuration.DbConfig()
	switch userType {
	case "Corporate":
		dbName = dbConfig.DbDatabaseName + "." + dbConfig.CrpPubDBName
		break
	case "University":
		dbName = dbConfig.DbDatabaseName + "." + dbConfig.UnvPubDBName
		break
	}
	sp, _ := RetriveSP("GET_PUB_SUB_TYPE")
	sp = strings.ReplaceAll(sp, "//REPLACE_DB_NAME", dbName)
	var gn, sh string
	err := Db.QueryRow(sp, publishID).Scan(&sh, &gn)
	if err != nil {
		return "", "O"
	}
	gn = strings.Split(gn, " has been published")[0]
	return sh, parseSubscriptionType(userType[:1] + gn)
}
