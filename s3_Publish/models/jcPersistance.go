// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Insert ...
func (jc *FullJobDb) Insert() <-chan DbModelError {
	job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		job <- customError
		return job
	}
	var jcID string
	jcID, customError = CreateJCID(jc.StakeholderID)
	if customError.ErrTyp != "000" {
		fmt.Printf("\nFailed to create Job ID :%+v\n", customError)
		job <- customError
		return job
	}
	jcInsertCmd, _ := RetriveSP("JOB_HC_MAP_INS")

	jsInsCmd, _ := RetriveSP("JOB_SKill_MAP_INS")

	vals := []interface{}{}

	for _, jobSkill := range jc.Jobs {
		jsInsCmd += "(?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, jcID, jc.JobHcMappingDB.JobName, jc.StakeholderID, jobSkill.SkillID, jobSkill.Skill, jobSkill.NoOfPositions, jobSkill.Location, jobSkill.SalaryRange, jobSkill.DateOfHiring, jobSkill.Status, jobSkill.Remarks, jobSkill.Attachment)
	}
	jsInsCmd = jsInsCmd[0 : len(jsInsCmd)-1]

	jcStmt, err := Db.Prepare(jcInsertCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare  job insert due to %v", err.Error())
		job <- customError
		return job
	}
	fmt.Println(jsInsCmd)
	jsStmt, err := Db.Prepare(jsInsCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare  skill insert due to %v %v", jsInsCmd, err.Error())
		job <- customError
		return job
	}
	if jc.HcID != "" && jc.HcName != "" {
		_, err = jcStmt.Exec(jcID, jc.StakeholderID, jc.HcID, jc.HcName, jc.JobName)
	} else {
		_, err = jcStmt.Exec(jcID, jc.StakeholderID, jc.HiringCriteriaID, jc.HiringCriteriaName, jc.JobName)
	}

	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert Job hiring Mapping in database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		job <- customError
		return job
	}

	_, err = jsStmt.Exec(vals...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert into Skills database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		job <- customError
		return job
	}

	customError.ErrTyp = "000"
	successResp["jcID"] = jcID
	customError.SuccessResp = successResp

	job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return job
}

// GetByID ...
func (jc *FullJobDb) GetByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	getByIDSP, _ := RetriveSP("JOB_HC_GET_BY_ID")
	err := Db.QueryRow(getByIDSP, jc.JobID).Scan(&jc.JobID, &jc.StakeholderID, &jc.HiringCriteriaID, &jc.HiringCriteriaName, &jc.JobName, &jc.CreationDate, &jc.PublishedFlag, &jc.PublishID)
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
	jcRows, err := Db.Query(getAllJCSP, jc.JobID)
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
		jc.Jobs = append(jc.Jobs, newJC)
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// DeleteByID ...
func (jc *JobHcMappingDB) DeleteByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	delByIDSP, _ := RetriveSP("JS_DELETE_BY_ID")
	_, err := Db.Exec(delByIDSP, jc.JobID, jc.StakeholderID) //.Scan(&hc.HcID, &hc.HcName, &hc.CorporateID, &hc.Program, &hc.Course, &hc.MinCutoffCategory, &hc.MinCutoff, &hc.ActiveBacklogsAllowed, &hc.EducationalGapsAllowed, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate)
	if err != nil {
		customError.ErrTyp = "S3PJ005"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Job_HiringCriteria : %v", err.Error())
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// DeleteSkillsByID ...
func (jc *JobSkillsMapping) DeleteSkillsByID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	delByIDSP, _ := RetriveSP("JS_SM_DELETE_BY_ID")
	_, err := Db.Exec(delByIDSP, jc.ID, jc.JobID, jc.StakeholderID) //.Scan(&hc.HcID, &hc.HcName, &hc.CorporateID, &hc.Program, &hc.Course, &hc.MinCutoffCategory, &hc.MinCutoff, &hc.ActiveBacklogsAllowed, &hc.EducationalGapsAllowed, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate)
	if err != nil {
		customError.ErrTyp = "S3PJ005"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Job_HiringCriteria : %v", err.Error())
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// MapHC ...
func (jc *JobHcMappingDB) MapHC(hcid string, hcName string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	updByIDSP, _ := RetriveSP("JOB_UPD_HC_MAP")
	_, err := Db.Exec(updByIDSP, hcid, hcName, jc.JobID, jc.StakeholderID) //.Scan(&hc.HcID, &hc.HcName, &hc.CorporateID, &hc.Program, &hc.Course, &hc.MinCutoffCategory, &hc.MinCutoff, &hc.ActiveBacklogsAllowed, &hc.EducationalGapsAllowed, &hc.YearOfPassing, &hc.Remarks, &hc.CreationDate)
	if err != nil {
		customError.ErrTyp = "S3PJ005"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Job_HiringCriteria : %v", err.Error())
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	return Job
}

// GetAllJC ...
func (jc *JobHcMappingDB) GetAllJC() (jcArray []JobHcMappingDB, err error) {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return jcArray, customError.Err
	}
	getAllJCSP, _ := RetriveSP("JOB_HC_GETALL_BY_ID")
	jcRows, err := Db.Query(getAllJCSP, jc.StakeholderID) //.Scan()
	if err != nil {
		return jcArray, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer jcRows.Close()
	for jcRows.Next() {
		var newJC JobHcMappingDB
		err = jcRows.Scan(&newJC.JobID, &newJC.HiringCriteriaID, &newJC.HiringCriteriaName, &newJC.JobName, &newJC.CreationDate, &newJC.PublishedFlag, &newJC.PublishID)

		if err != nil {
			return jcArray, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		if newJC.HiringCriteriaID.Valid {
			newJC.HcID = newJC.HiringCriteriaID.String
		}
		if newJC.HiringCriteriaName.Valid {
			newJC.HcName = newJC.HiringCriteriaName.String
		}
		jcArray = append(jcArray, newJC)
	}
	return jcArray, nil
}

// CreateJCID ...
func CreateJCID(crpID string) (string, DbModelError) {
	rowSP, _ := RetriveSP("JOB_HC_Last_ID")
	lastID := ""
	//recordsExists := true
	err := Db.QueryRow(rowSP, crpID).Scan(&lastID)
	var idCreationError DbModelError
	if err != nil && err != sql.ErrNoRows {
		idCreationError.ErrTyp = "500"
		idCreationError.Err = fmt.Errorf("Failed to create Job Creation ID ", err)
		idCreationError.ErrCode = "S3PJ001"
		return "", idCreationError
	}
	if err == sql.ErrNoRows {
		lastID = "000000000000000"
	}

	corporateNum, _ := strconv.Atoi(crpID[8:])
	countNum, _ := strconv.Atoi(lastID[len(lastID)-8:])
	fmt.Println("------>", countNum)
	idCreationError.ErrTyp = "000"
	return "JD" + strconv.Itoa(corporateNum) + (fmt.Sprintf("%08d", (countNum + 1))), idCreationError

}

// AddSkillsToJC ....
func (jc *SkillsUpdateJobDb) AddSkillsToJC() <-chan DbModelError {
	job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		job <- customError
		return job
	}

	jsInsCmd, _ := RetriveSP("JOB_SKill_MAP_INS")
	jsDelCmd, _ := RetriveSP("JS_SM_DELETE_All")

	vals := []interface{}{}
	for _, jobSkill := range jc.Jobs {
		jsInsCmd += "(?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, jc.JobID, jc.JobName, jc.StakeholderID, jobSkill.SkillID, jobSkill.Skill, jobSkill.NoOfPositions, jobSkill.Location, jobSkill.SalaryRange, jobSkill.DateOfHiring, jobSkill.Status, jobSkill.Remarks, jobSkill.Attachment)
	}
	jsInsCmd = jsInsCmd[0 : len(jsInsCmd)-1]
	fmt.Println(jsInsCmd)
	fmt.Printf("\n%+v\n", vals)
	jsStmt, err := Db.Prepare(jsInsCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare  skill insert due to %v %v", jsInsCmd, err.Error())
		job <- customError
		return job
	}
	jsDelStmt, err := Db.Prepare(jsDelCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare  skill insert due to %v %v", jsInsCmd, err.Error())
		job <- customError
		return job
	}
	_, err = jsDelStmt.Exec(jc.JobID, jc.StakeholderID)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert into Skills database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		job <- customError
		return job
	}

	_, err = jsStmt.Exec(vals...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert into Skills database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		job <- customError
		return job
	}
	customError.ErrTyp = "000"
	successResp["jcID"] = jc.JobID
	customError.SuccessResp = successResp
	job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return job
}

// // UpdateJobSkills ...
// func UpdateJobSkills(skills []string, jcID string, stakeholder string) DbModelError {
// 	var customError DbModelError
// 	delByIDSP, _ := RetriveSP("JS_DELETE_BY_ID")
// 	_, err := Db.Exec(delByIDSP, jcID)
// 	if err != nil {
// 		customError.ErrTyp = "S3PJ005"
// 		customError.ErrCode = "500"
// 		customError.Err = fmt.Errorf("Failed to Delete Job skills : %v", err.Error())
// 		return customError
// 	}
// 	jsInsCmd, _ := RetriveSP("JS_INS_NEW")

// 	vals := []interface{}{}

// 	for _, skill := range skills {
// 		jsInsCmd += "(?, ?),"
// 		vals = append(vals, jcID, skill)
// 	}
// 	jsInsCmd = jsInsCmd[0 : len(jsInsCmd)-1]
// 	// vals = append(vals, jcID)
// 	// vals = append(vals, stakeholder)
// 	jsStmt, err := Db.Prepare(jsInsCmd)
// 	if err != nil {
// 		customError.ErrTyp = "500"
// 		customError.ErrCode = "S3PJ002"
// 		customError.Err = fmt.Errorf("Cannot prepare  skill insert due to %v %v", jsInsCmd, err.Error())
// 		return customError
// 	}
// 	_, err = jsStmt.Exec(vals...)
// 	if err != nil {
// 		customError.ErrTyp = "500"
// 		customError.Err = fmt.Errorf("Failed to insert into Skills database due to : %v", err.Error())
// 		customError.ErrCode = "S3PJ002"
// 		return customError
// 	}
// 	customError.ErrTyp = "000"
// 	return customError
// }
