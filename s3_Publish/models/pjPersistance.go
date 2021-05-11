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
func (pj *PublishJobs) Insert(ID string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating PublishHistory
	pdhIDs, customError := CreatePJID(ID, len(pj.PublishedJobs), "PDH", "PDH_Get_Last_ID")
	if customError.ErrTyp != "000" {
		fmt.Printf("\nDb connection error :%+v\n", customError)
		Job <- customError
		return Job
	}

	// Preparing Database insert
	pjInsertCmd, _ := RetriveSP("JOB_PUBLISH_BY_ID")

	// Preparing Database insert
	pdhInsertCmd, _ := RetriveSP("PDH_INS_NEW")
	pdhVals := []interface{}{}

	currentTime := time.Now()

	for index := range pj.PublishedJobs {
		pdhInsertCmd += "(?,?,?,?,?,?,?,?,?,?,?),"

		var jc FullJobDb
		getByIDSP, _ := RetriveSP("JOB_HC_GET_BY_ID")
		err := Db.QueryRow(getByIDSP, pj.PublishedJobs[index].JobID).Scan(&jc.JobID, &jc.StakeholderID, &jc.HiringCriteriaID, &jc.HiringCriteriaName, &jc.JobName, &jc.CreationDate, &jc.PublishedFlag, &jc.PublishID)
		if err != nil {
			customError.ErrTyp = "S3PJ003"
			customError.ErrCode = "500"
			customError.Err = fmt.Errorf("Failed to retrieve Created Jobs : %v", err.Error())
			Job <- customError
			return Job
		}
		if jc.HiringCriteriaID.Valid {
			jc.HcID = jc.HiringCriteriaID.String
		}
		if jc.HiringCriteriaName.Valid {
			jc.HcName = jc.HiringCriteriaName.String
		}
		getAllJCSP, _ := RetriveSP("JOB_SKill_GET_BY_ID")
		jcRows, err := Db.Query(getAllJCSP, pj.PublishedJobs[index].JobID)
		if err != nil {
			customError.ErrTyp = "S3PJ003"
			customError.ErrCode = "500"
			customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
			Job <- customError
			return Job
		}
		defer jcRows.Close()
		for jcRows.Next() {
			var newJC JobSkillsMapping
			err = jcRows.Scan(&newJC.ID, &newJC.JobID, &newJC.JobName, &newJC.SkillID, &newJC.Skill, &newJC.NoOfPositions, &newJC.Location, &newJC.SalaryRange, &newJC.DateOfHiring, &newJC.Status, &newJC.Remarks, &newJC.Attachment, &newJC.CreationDate)
			if err != nil {
				customError.ErrTyp = "S3PJ003"
				customError.ErrCode = "500"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				Job <- customError
				return Job
			}
			newJC.Attachment = []byte("")
			jc.Jobs = append(jc.Jobs, newJC)
		}
		jcPubDataAsByte, _ := json.Marshal(&jc)
		jcPubDataAsByteUnescaped, _ := json.RawMessage(jcPubDataAsByte).MarshalJSON()

		pdhVals = append(pdhVals, ID, pdhIDs[index], currentTime, false, true, false, false, "New Job has been published", currentTime, currentTime, string(jcPubDataAsByteUnescaped))
	}
	//pjInsertCmd = pjInsertCmd[0 : len(pjInsertCmd)-1]
	stmt, err := Db.Prepare(pjInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare Published job insert due to %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}

	pdhInsertCmd = pdhInsertCmd[0 : len(pdhInsertCmd)-1]
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
	for i := range pdhIDs {
		_, err = stmt.Exec(pdhIDs[i], currentTime, pj.PublishedJobs[i].JobID, ID)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Failed to Update Published job in database due to : %v", err.Error())
			customError.ErrCode = "S3PJ002"
			Job <- customError
			return Job
		}
	}

	customError.ErrTyp = "000"
	successResp["pjIDs"] = fmt.Sprintf("%v", pdhIDs)
	customError.SuccessResp = successResp

	Job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return Job
}

// GetByID ...
func (pj *PublishedJobsDB) GetByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	getByIDSP, _ := RetriveSP("PJ_GET_BY_ID")
	err := Db.QueryRow(getByIDSP, pj.PublishID).Scan(&pj.PublishID, &pj.JobID, &pj.JobName, &pj.StakeholderID, &pj.CreationDate)
	if err != nil {
		customError.ErrTyp = "S3PJ003"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Published job : %v", err.Error())
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// DeleteByID ...
func (pj *PublishedJobsDB) DeleteByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	delByIDSP, _ := RetriveSP("PJ_DELETE_BY_ID")
	_, err := Db.Exec(delByIDSP, pj.PublishID, pj.StakeholderID) //.Scan(&hc.HcID, &hc.HcName, &hc.StakeholderID, &hc.Program, &hc.Course, &hc.MinCutoffCategory, &hc.MinCutoff, &hc.ActiveBacklogsAllowed, &hc.EducationalGapsAllowed, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate)
	if err != nil {
		customError.ErrTyp = "S3PJ005"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to Delete Published Job : %v", err.Error())
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// GetAllPJ ...
func (pj *PublishedJobsDB) GetAllPJ() (pjArray []PublishedJobsDB, err error) {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return pjArray, customError.Err
	}
	getAllHCSP, _ := RetriveSP("PJ_GET_ALL")
	pjRows, err := Db.Query(getAllHCSP, pj.StakeholderID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		return pjArray, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	if err == sql.ErrNoRows {
		return pjArray, nil
	}
	defer pjRows.Close()
	for pjRows.Next() {
		var newPJ PublishedJobsDB
		err = pjRows.Scan(&newPJ.PublishID, &newPJ.JobID, &newPJ.JobName, &newPJ.StakeholderID, &newPJ.CreationDate)
		if err != nil {
			return pjArray, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		pjArray = append(pjArray, newPJ)
	}
	return pjArray, nil
}

// CreatePJID ...
func CreatePJID(crpID string, count int, code string, queryStr string) ([]string, DbModelError) {
	rowSP, _ := RetriveSP(queryStr)
	lastID := ""
	err := Db.QueryRow(rowSP, crpID).Scan(&lastID)
	var idCreationError DbModelError
	if err != nil && err != sql.ErrNoRows {
		idCreationError.ErrTyp = "500"
		idCreationError.Err = fmt.Errorf("Failed to create Published Job ID ", err)
		idCreationError.ErrCode = "S3PJ001"
		return []string{}, idCreationError
	}
	if err == sql.ErrNoRows {
		lastID = "0000000000000"
	}
	corporateNum, _ := strconv.Atoi(crpID[8:])
	countNum, _ := strconv.Atoi(lastID[len(lastID)-8:])
	idCreationError.ErrTyp = "000"
	var ids []string
	for i := 0; i < count; i++ {
		ids = append(ids, code+strconv.Itoa(corporateNum)+(fmt.Sprintf("%08d", (countNum+(i+1)))))
	}
	fmt.Println(ids)
	return ids, idCreationError
}
