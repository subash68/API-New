// Package models ...
package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Insert ...
func (oi *OtherInformationModel) Insert() <-chan DbModelError {

	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	oiInsertCmd, _ := RetriveSP("OI_INS_NEW")

	oiStmt, err := Db.Prepare(oiInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare  published history insert due to %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	currentTime := time.Now()
	row, err := oiStmt.Exec(oi.StakeholderID, oi.Title, oi.Information, oi.Attachment, currentTime, currentTime)
	insID, _ := row.LastInsertId()
	fmt.Printf("\n ============= Last inserted ID : %v ==\n", insID)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert  Published History in database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	successResp["insID"] = strconv.Itoa(int(insID))
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// PublishOI ...
func (oi *OtherInformationModel) PublishOI() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating PublishHistory
	pdhIDs, customError := CreatePJID(oi.StakeholderID, 1, "PDH", "PDH_Get_Last_ID")
	if customError.ErrTyp != "000" {
		fmt.Printf("\nDb connection error :%+v\n", customError)
		Job <- customError
		return Job
	}
	// Preparing Database insert
	pdhInsertCmd, _ := RetriveSP("PDH_INS_NEW")
	pdhVals := []interface{}{}

	currentTime := time.Now()

	getAllHCSP, _ := RetriveSP("OI_GET_BY_ID")
	err := Db.QueryRow(getAllHCSP, oi.StakeholderID, oi.ID).Scan(&oi.Title, &oi.Information, &oi.CreationDate)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to retrive Other information due to %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	oiDataAsBytes, _ := json.Marshal(&oi)
	pdhInsertCmd += "(?,?,?,?,?,?,?,?,?,?,?)"

	pdhVals = append(pdhVals, oi.StakeholderID, pdhIDs[0], currentTime, false, false, false, true, "Other Information has been published", currentTime, currentTime, string(oiDataAsBytes))

	pdhStmt, err := Db.Prepare(pdhInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare  published history insert due to %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	_, err = pdhStmt.Exec(pdhVals...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert  Published History in database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}

	updateSP, _ := RetriveSP("OI_UPDATE_BY_TITLE")
	stmtWhere, _ := RetriveSP("OI_UPDATE_WHERE")

	updateSP = updateSP + "  Publish_ID='" + pdhIDs[0] + "', PublishFlag=1 " + stmtWhere
	updateStm, err := Db.Prepare(updateSP)
	if err != nil {
		fmt.Println(updateSP)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare database update due to %v", err.Error())
		Job <- customError
		return Job
	}
	_, err = updateStm.Exec(oi.ID, oi.StakeholderID)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to update the database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	successResp["publishID"] = pdhIDs[0]
	customError.SuccessResp = successResp

	Job <- customError
	return Job

}

// GetAllOI ...
func (oi *OtherInformationModel) GetAllOI(query string) (oiArray []OtherInformationModel, err error) {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return oiArray, customError.Err
	}
	getAllHCSP, _ := RetriveSP(query)
	hcRows, err := Db.Query(getAllHCSP, oi.StakeholderID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		return oiArray, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return oiArray, nil
	}
	defer hcRows.Close()
	for hcRows.Next() {
		var newOI OtherInformationModel
		err = hcRows.Scan(&newOI.ID, &newOI.Title, &newOI.Information, &newOI.Attachment, &newOI.CreationDate, &newOI.LastUpdatedDate, &newOI.PublishedFlag, &newOI.PublishID)
		if err != nil {
			return oiArray, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		oiArray = append(oiArray, newOI)
	}
	return oiArray, nil
}
