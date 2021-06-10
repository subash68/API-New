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

// RegCount ...
type RegCount struct {
	CorpCount                     int `json:"corporatesRegistered,omitempty"`
	StuCount                      int `json:"studentsRegistered,omitempty"`
	UnvCount                      int `json:"universitiesRegistered,omitempty"`
	JobsPublished                 int `json:"jobsPublished,omitempty"`
	JobsPublishedInMonth          int `json:"jobsPublishedInMonth,omitempty"`
	ApplicationsReceived          int `json:"applicationsReceived,omitempty"`
	ApplicationsReceivedInTwoDays int `json:"applicationsReceivedInTwoDays,omitempty"`
	JobOffersMade                 int `json:"jobOffersMade,omitempty"`
	JobOffersMadeInLastMonth      int `json:"jobOffersMadeInLastMonth,omitempty"`
	CurrentlyOnline               int `json:"currentlyOnline,omitempty"`
	JoinedLastWeek                int `json:"joinedLastWeek,omitempty"`
	GotPlaced                     int `json:"gotPlaced,omitempty"`
	JobOpenings                   int `json:"jobOpenings,omitempty"`
	StudentsPlaced                int `json:"studentsPlaced,omitempty"`
	StudentsAwaitingJobs          int `json:"studentsAwaitingJobs,omitempty"`
	StudentsRegInLastWeek         int `json:"studentsRegisteredInLastWeek,omitempty"`
	CandidatesHiredInLastWeek     int `json:"candidatesHiredInLastWeek,omitempty"`
	CandidatesHiredSoFar          int `json:"candidatesHiredSoFar",omitempty"`
	JobsPostedInLastWeek          int `json:"jobsPostedInLastWeek,omitempty"`
	JobsPostedTillDate            int `json:"jobsPostedTillDate,omitempty"`
	JobOpeningsInLastMonth        int `json:"jobOpeningsInLastMonth,omitempty"`
	StudentsPublishedInLastYear   int `json:"studentsPublishedInLastYear,omitempty"`
	StudentsAwaitingInCurrentYear int `json:"studentsAwaitingInCurrentYear,omitempty"`
}

// GetRegisteredCounts ...
func GetRegisteredCounts() RegCount {
	var rc RegCount
	dbQuery, _ := RetriveSP("REG_SH_COUNT")
	_ = Db.QueryRow(dbQuery).Scan(&rc.CorpCount, &rc.StuCount, &rc.UnvCount)
	rc.CurrentlyOnline = 1571
	rc.JoinedLastWeek = 82
	rc.GotPlaced = 25152
	return rc
}

// GetJobsPublishedCount ...
func (rc *RegCount) GetJobsPublishedCount(ID string) {
	dbQuery, _ := RetriveSP("PJ_GET_COUNT")
	_ = Db.QueryRow(dbQuery, ID).Scan(&rc.JobsPublished)
	fmt.Printf("Getting Jobs published %s = %v ", dbQuery, rc.JobsPublished)
	rc.ApplicationsReceived = 647
	rc.JobOffersMade = 392
	rc.JobOffersMadeInLastMonth = 94
	rc.ApplicationsReceivedInTwoDays = 49
	rc.JobOffersMadeInLastMonth = 58
	return
}

// GetUnvStatsCount ...
func (rc *RegCount) GetUnvStatsCount() {
	// dbQuery, _ := RetriveSP("PJ_GET_COUNT")
	// _ = Db.QueryRow(dbQuery, ID).Scan(&rc.JobsPublished)
	// fmt.Printf("Getting Jobs published %s = %v ", dbQuery, rc.JobsPublished)
	rc.JobOpenings = 647
	rc.StudentsPlaced = 592
	rc.StudentsAwaitingJobs = 172
	rc.JobOpeningsInLastMonth = 151
	rc.StudentsPublishedInLastYear = 4877
	rc.StudentsAwaitingInCurrentYear = 247
	return
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
