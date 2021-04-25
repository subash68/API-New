// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Insert ...
func (oi *UnvOtherInformationModel) Insert() <-chan DbModelError {

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
	_, err = oiStmt.Exec(oi.StakeholderID, oi.Title, oi.Information, oi.Attachment, currentTime, currentTime)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert  Published History in database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// PublishOI ...
func (oi *UnvOtherInformationModel) PublishOI() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating PublishHistory
	pdhIDs, customError := CreateUnvPublishID(oi.StakeholderID, "UPDH", "UNV_PDH_Get_Last_ID")
	if customError.ErrTyp != "000" {
		fmt.Printf("\nDb connection error :%+v\n", customError)
		Job <- customError
		return Job
	}
	// Preparing Database insert
	pdhInsertCmd, _ := RetriveSP("UNV_PDH_INS_NEW")
	pdhVals := []interface{}{}

	currentTime := time.Now()

	pdhInsertCmd += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	pdhVals = append(pdhVals, oi.StakeholderID, pdhIDs, currentTime, false, false, false, false, false, false, true, false, false, "Other Information has been published", currentTime, currentTime, "[]")

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

	oiInsertCmd, _ := RetriveSP("UNV_OI_INS_NEW")

	oiStmt, err := Db.Prepare(oiInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare  published history insert due to %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	_, err = oiStmt.Exec(oi.StakeholderID, oi.Title, oi.Information, oi.Attachment, true, pdhIDs, currentTime, currentTime)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert  Published History in database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}

	customError.ErrTyp = "000"
	successResp["publishID"] = pdhIDs
	customError.SuccessResp = successResp

	Job <- customError
	return Job

}

// GetAllOI ...
func (oi *UnvOtherInformationModel) GetAllOI(query string) (oiArray []OtherInformationModel, err error) {
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
		err = hcRows.Scan(&newOI.ID, &newOI.Title, &newOI.Information, &newOI.PublishID, &newOI.Attachment, &newOI.CreationDate, &newOI.LastUpdatedDate, &newOI.PublishedFlag)
		if err != nil {
			return oiArray, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		oiArray = append(oiArray, newOI)
	}
	return oiArray, nil
}
