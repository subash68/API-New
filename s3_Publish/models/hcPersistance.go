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
func (hc *MultipleHC) Insert(sID string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating HCID
	hcIDs, customError := CreateHCID(sID, len(hc.HiringCriterias), hc.HiringCriterias)
	if customError.ErrTyp != "000" {
		fmt.Printf("\nFailed to Generate Hiring Criteria IDs :%+v\n", customError)
		Job <- customError
		return Job
	}

	// Preparing Databse insert
	hcInsertCmd, _ := RetriveSP("HC_INS_NEW")

	vals := []interface{}{}

	for index, hc := range hc.HiringCriterias {
		hcInsertCmd += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, hcIDs[index], hc.HiringCriteriaName, sID, hc.ProgramID, hc.DepartmentID, hc.CutOffCategory, hc.CutOff, hc.EduGapsSchoolAllowed, hc.EduGaps11N12Allowed, hc.EduGapsGradAllowed, hc.EduGapsPGAllowed, hc.AllowActiveBacklogs, hc.NumberOfAllowedBacklogs, hc.YearOfPassing, hc.Remarks)
	}
	hcInsertCmd = hcInsertCmd[0 : len(hcInsertCmd)-1]

	stmt, err := Db.Prepare(hcInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare -- %v , %v -- insert due to %v", hcInsertCmd, vals, err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert in database -- %v , %v -- insert due to %v", hcInsertCmd, vals, err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	successResp["hcIDs"] = fmt.Sprintf("%v", hcIDs)
	customError.SuccessResp = successResp

	Job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return Job
}

// GetByID ...
func (hc *HiringCriteriaDB) GetByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	getByIDSP, _ := RetriveSP("HC_GET_BY_ID")
	err := Db.QueryRow(getByIDSP, hc.HiringCriteriaID).Scan(&hc.HiringCriteriaID, &hc.HiringCriteriaName, &hc.ProgramID, &hc.DepartmentID, &hc.CutOffCategory, &hc.CutOff, &hc.EduGapsSchoolAllowed, &hc.EduGaps11N12Allowed, &hc.EduGapsGradAllowed, &hc.EduGapsPGAllowed, &hc.AllowActiveBacklogs, &hc.NumberOfAllowedBacklogs, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate, &hc.PublishedFlagNull, &hc.PublishIDNull)

	if err != nil {
		customError.ErrTyp = "S3PJ003"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Hiring criteria : %v", err.Error())
		Job <- customError
		return Job
	}
	if hc.PublishedFlagNull.Valid {
		hc.PublishedFlag = hc.PublishedFlagNull.Bool
	}

	if hc.PublishIDNull.Valid {
		hc.PublishID = hc.PublishIDNull.String
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// DeleteByID ...
func (hc *HiringCriteriaDB) DeleteByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	delByIDSP, _ := RetriveSP("HC_DELETE_BY_ID")
	_, err := Db.Exec(delByIDSP, hc.HiringCriteriaID, hc.StakeholderID) //.Scan(&hc.HcID, &hc.HcName, &hc.CorporateID, &hc.Program, &hc.Course, &hc.MinCutoffCategory, &hc.MinCutoff, &hc.ActiveBacklogsAllowed, &hc.EducationalGapsAllowed, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate)
	if err != nil {
		customError.ErrTyp = "S3PJ005"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Hiring criteria : %v", err.Error())
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// GetAllHC ...
func (hc *HiringCriteriaDB) GetAllHC(query string) (hcArray []HiringCriteriaDB, err error) {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return hcArray, customError.Err
	}
	getAllHCSP, _ := RetriveSP(query)
	hcRows, err := Db.Query(getAllHCSP, hc.StakeholderID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		return hcArray, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return hcArray, nil
	}
	defer hcRows.Close()
	for hcRows.Next() {
		var newHC HiringCriteriaDB
		err = hcRows.Scan(&newHC.HiringCriteriaID, &newHC.HiringCriteriaName, &newHC.ProgramID, &newHC.DepartmentID, &newHC.CutOffCategory, &newHC.CutOff, &newHC.EduGapsSchoolAllowed, &newHC.EduGaps11N12Allowed, &newHC.EduGapsGradAllowed, &newHC.EduGapsPGAllowed, &newHC.AllowActiveBacklogs, &newHC.NumberOfAllowedBacklogs, &newHC.YearOfPassing, &newHC.Remarks, &newHC.CreationDate, &newHC.PublishedFlagNull, &newHC.PublishIDNull)
		if newHC.PublishedFlagNull.Valid {
			newHC.PublishedFlag = newHC.PublishedFlagNull.Bool
		} else {
			newHC.PublishedFlag = false
		}

		if newHC.PublishIDNull.Valid {
			newHC.PublishID = newHC.PublishIDNull.String
		} else {
			newHC.PublishID = ""
		}
		if err != nil {
			return hcArray, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		hcArray = append(hcArray, newHC)
	}
	return hcArray, nil
}

// CreateHCID ...
func CreateHCID(crpID string, count int, hc []HiringCriteriaDB) ([]string, DbModelError) {
	rowSP, _ := RetriveSP("HC_Get_Last_ID")
	lastID := ""
	err := Db.QueryRow(rowSP, crpID).Scan(&lastID)

	var idCreationError DbModelError
	if err != nil && err != sql.ErrNoRows {
		idCreationError.ErrTyp = "500"
		idCreationError.Err = fmt.Errorf("Failed to create Hiring Criteria ID ", err)
		idCreationError.ErrCode = "S3PJ001"
		return []string{}, idCreationError
	}
	if err == sql.ErrNoRows {
		lastID = "0000000000000"
	}
	corporateNum, _ := strconv.Atoi(crpID[7:])
	countNum, _ := strconv.Atoi(lastID[len(lastID)-7:])
	fmt.Println("--------------------> ", lastID, countNum)
	idCreationError.ErrTyp = "000"
	var ids []string
	for i := 0; i < count; i++ {
		ids = append(ids, ("HC" + strconv.Itoa(corporateNum) + (fmt.Sprintf("%07d", (countNum + (i + 1))))))

	}
	return ids, idCreationError
}

// PublishHC ...
func (phc *PublishHiringCriteriasModel) PublishHC() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating PublishHistory
	pdhIDs, customError := CreatePJID(phc.StakeholderID, len(phc.HiringCriteriaIDs), "PDH", "PDH_Get_Last_ID")
	if customError.ErrTyp != "000" {
		fmt.Printf("\nDb connection error :%+v\n", customError)
		Job <- customError
		return Job
	}
	// Preparing Database insert
	pdhInsertCmd, _ := RetriveSP("PDH_INS_NEW")
	pdhVals := []interface{}{}

	currentTime := time.Now()
	for index := range phc.HiringCriteriaIDs {

		pdhInsertCmd += "(?,?,?,?,?,?,?,?,?,?,?),"
		getByIDSP, _ := RetriveSP("HC_GET_BY_ID")
		var hc HiringCriteriaDB
		err := Db.QueryRow(getByIDSP, phc.HiringCriteriaIDs[index]).Scan(&hc.HiringCriteriaID, &hc.HiringCriteriaName, &hc.ProgramID, &hc.DepartmentID, &hc.CutOffCategory, &hc.CutOff, &hc.EduGapsSchoolAllowed, &hc.EduGaps11N12Allowed, &hc.EduGapsGradAllowed, &hc.EduGapsPGAllowed, &hc.AllowActiveBacklogs, &hc.NumberOfAllowedBacklogs, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate, &hc.PublishedFlagNull, &hc.PublishIDNull)
		hcPubDataAsBytes, _ := json.Marshal(&hc)
		if err != nil {
			customError.ErrTyp = "S3PJ003"
			customError.ErrCode = "500"
			customError.Err = fmt.Errorf("Failed to retrieve Hiring criteria : %v", err.Error())
			Job <- customError
			return Job
		}

		pdhVals = append(pdhVals, phc.StakeholderID, pdhIDs[index], currentTime, true, false, false, false, "Hiring Criteria has been published", currentTime, currentTime, string(hcPubDataAsBytes))
	}

	pdhInsertCmd = pdhInsertCmd[0 : len(pdhInsertCmd)-1]
	pdhStmt, err := Db.Prepare(pdhInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare  published history insert due to %v --- %s --- %+v ", err.Error(), pdhInsertCmd, pdhVals)
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	_, err = pdhStmt.Exec(pdhVals...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert  Published History in database due to : %v --- %s --- %+v ", err.Error(), pdhInsertCmd, pdhVals)
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}

	for index := range phc.HiringCriteriaIDs {
		updateSP, _ := RetriveSP("HC_UPDATE_BY_ID")
		stmtWhere, _ := RetriveSP("HC_UPDATE_WHERE")

		updateSP = updateSP + " PublishID='" + pdhIDs[index] + "', PublishFlag=1 " + stmtWhere

		updateStm, err := Db.Prepare(updateSP)
		if err != nil {
			fmt.Println(updateSP)
			customError.ErrTyp = "500"
			customError.ErrCode = "S3PJ002"
			customError.Err = fmt.Errorf("Cannot prepare database update due to %v --- %s", err.Error(), updateSP)
			Job <- customError
			return Job
		}
		_, err = updateStm.Exec(phc.HiringCriteriaIDs[index], phc.StakeholderID)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Failed to update the database due to : %v", err.Error())
			customError.ErrCode = "S3PJ002"
			Job <- customError
			return Job
		}
	}
	customError.ErrTyp = "000"
	successResp["publishID"] = fmt.Sprintf("%v", pdhIDs)
	customError.SuccessResp = successResp

	Job <- customError
	return Job

}
