package models

import (
	"database/sql"
	"fmt"
	"net/url"
	"reflect"
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
func UpdateProfileData(updateQuery url.Values, spName string, spExt string, stakeholder string, attachmentUpdate bool, attachment []byte, attachmentName string, ppUpdate bool, ppFile []byte) DbModelError {
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
		updateString = updateString + " Attachment= ?, AttachFile_Name= ?,"
		values = append(values, attachment, attachmentName)
	}
	if ppUpdate {
		updateString = updateString + " ProfilePicture= ?,"
		values = append(values, ppFile)
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
func UpdateProfilePic(pic []byte, userID string, userType string, sp string) DbModelError {
	var customError DbModelError

	switch userType {
	case "Corporate":
		sp = "CORP_" + sp
		break
	case "University":
		sp = "UNV_" + sp
		break
	case "Student":
		sp = "STU_" + sp
		break
	default:
		customError.ErrTyp = "S3PRF001"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Invalid Usertype")
		return customError
	}
	updatePP, _ := RetriveSP(sp)
	updateStm, err := Db.Prepare(updatePP)
	if err != nil {
		fmt.Println(updatePP)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare database Prepare due to %v", err.Error())
		return customError
	}
	_, err = updateStm.Exec(pic, userID)
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
func GetProfilePic(stakeholderID string, userType string, sp string) ([]byte, DbModelError) {
	var customError DbModelError
	var ppic []byte
	switch userType {
	case "Corporate":
		sp = "CORP_" + sp
		break
	case "University":
		sp = "UNV_" + sp
		break
	case "Student":
		sp = "STU_" + sp
		break
	default:
		customError.ErrTyp = "S3PRF001"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Invalid Usertype")
		return ppic, customError
	}
	queryPP, _ := RetriveSP(sp)

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
