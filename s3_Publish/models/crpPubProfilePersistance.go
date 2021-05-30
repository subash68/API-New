// Package models ...
package models

import (
	"fmt"
	"time"
)

// Insert ...
func (profilePub *PublishDataModel) Insert() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating PublishHistory
	pdhIDs, customError := CreatePJID(profilePub.StakeholderID, 1, "PDH", "PDH_Get_Last_ID")
	if customError.ErrTyp != "000" {
		fmt.Printf("\nDb connection error :%+v\n", customError)
		Job <- customError
		return Job
	}
	// Preparing Database insert
	pdhInsertCmd, _ := RetriveSP("PDH_INS_NEW")
	pdhVals := []interface{}{}

	currentTime := time.Now()

	pdhInsertCmd += "(?,?,?,?,?,?,?,?,?,?,?)"

	pdhVals = append(pdhVals, profilePub.StakeholderID, pdhIDs[0], currentTime, false, false, true, false, "Profile details has been published", currentTime, currentTime, profilePub.PublishData)

	//pdhInsertCmd = pdhInsertCmd[0 : len(pdhInsertCmd)-1]
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
	customError.ErrTyp = "000"
	successResp["publishID"] = pdhIDs[0]
	customError.SuccessResp = successResp

	pdhInsertCmd, _ = RetriveSP("CRP_PRF_PUB_UPD")
	pdhStmt, err = Db.Prepare(pdhInsertCmd)
	fmt.Println("Query ==> ", pdhInsertCmd, profilePub.StakeholderID)
	if err != nil {
		fmt.Printf("Cannot prepare Profile Publish update due to %v", err.Error())
	}

	_, err = pdhStmt.Exec(profilePub.StakeholderID)
	if err != nil {
		fmt.Printf("Cannot prepare Profile Publish update due to %v", err.Error())
	}

	Job <- customError
	return Job
}
