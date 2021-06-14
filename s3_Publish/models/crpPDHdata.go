package models

import (
	"database/sql"
	"fmt"
)

// GetCrpPublishedDataByID ...
func (cp *CorpPushedDataModel) GetCrpPublishedDataByID() ([]CorpPushedDataModel, error) {
	// CRP_PDH_GET_ALL
	var sa []CorpPushedDataModel
	getAllHCSP, _ := RetriveSP("CRP_PDH_GET_ALL")
	hcRows, err := Db.Query(getAllHCSP, cp.StakeholderID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		return sa, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return sa, nil
	}
	defer hcRows.Close()
	for hcRows.Next() {
		var newSA CorpPushedDataModel
		err = hcRows.Scan(&newSA.PublishID, &newSA.DateOfPublish, &newSA.HiringCriteriaPublished, &newSA.JobsPublished, &newSA.ProfilePublished, &newSA.OtherPublished, &newSA.GeneralNote, &newSA.CreationDate, &newSA.LastUpdatedDate, &newSA.PublishedData)
		if newSA.PublishID == "PDH200000049" {
			fmt.Println("Published data for PDH200000049 : ", newSA.PublishedData)
		}
		if err != nil {
			return sa, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		sa = append(sa, newSA)
	}
	return sa, nil
}

// GetCrpPublishedData ...
func GetCrpPublishedData(publishID string, isOwner bool, subscriber string, subType string) (DbModelError, map[string]interface{}, string) {
	var resp map[string]interface{}
	var customError DbModelError
	customError.ErrTyp = "000"
	isSubscribed := false
	if !isOwner {
		customError, isSubscribed = validateSubscription(publishID, subscriber, subType)
		if customError.ErrTyp != "000" {
			return customError, resp, ""
		}
		if isSubscribed == false {
			customError.ErrTyp = "500"
			customError.ErrCode = "S3PJ003"
			customError.Err = fmt.Errorf("Invalid Subscription, Subscribe to view details")
			return customError, resp, ""
		}
		fmt.Println("isSubscribed -> ", isSubscribed)
	}
	// Get Published Data
	var hc, jb, prf, oi bool
	var pd string

	getByIDSP, _ := RetriveSP("CRP_PDH_GET_PID")
	err := Db.QueryRow(getByIDSP, publishID).Scan(&hc, &jb, &prf, &oi, &pd)
	fmt.Println(getByIDSP)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "S3PJ003"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByIDSP)
		return customError, resp, ""
	} else if err == sql.ErrNoRows {
		customError.ErrTyp = "S3PJ003"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("No data found for specified Publish ID")
		return customError, resp, ""
	}
	fmt.Println("====================>>>>>>>>>>>>>>>>>", hc, jb, prf, oi)

	if jb {
		// var jpdh JobPdhModel
		// getByPID, _ := RetriveSP("CRP_GET_JOB_BY_PID")
		// getByPID = strings.ReplaceAll(getByPID, "34", fmt.Sprint('"'))
		// err := Db.QueryRow(getByPID, publishID).Scan(&jpdh.JobID, &jpdh.JobName, &jpdh.CorporateName, &jpdh.ProgramID, &jpdh.BranchID, &jpdh.MinimumCutoffCategory, &jpdh.MinimumCutoff, &jpdh.ActiveBacklogsAllowed, &jpdh.TotalNumberOfBacklogsAllowed, &jpdh.EduGaps11N12Allowed, &jpdh.EduGapsGradAllowed, &jpdh.EduGapsSchoolAllowed, &jpdh.EduGapsPGAllowed, &jpdh.YearOfPassing, &jpdh.Remarks, &jpdh.Skills)
		// if err != nil {
		// 	customError.ErrTyp = "S3PJ003"
		// 	customError.ErrCode = "500"
		// 	customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByPID)
		// 	return customError, resp, ""
		// }
		// jpdh.Skills = strings.ReplaceAll(jpdh.Skills, "34", "\"")

		// dataAsByte, _ := json.Marshal(jpdh)
		resp = map[string]interface{}{"Jobs": pd}
		fmt.Printf("\n============== %s =================== \n", resp["Jobs"])
		return customError, resp, "Jobs"
	} else if hc {
		// var jpdh HcPdhModel
		// getByPID, _ := RetriveSP("CRP_GET_HC_BY_PID")
		// getByPID = strings.ReplaceAll(getByPID, "34", fmt.Sprint('"'))
		// err := Db.QueryRow(getByPID, publishID).Scan(&jpdh.HcID, &jpdh.HcName, &jpdh.CorporateName, &jpdh.ProgramID, &jpdh.BranchID, &jpdh.MinimumCutoffCategory, &jpdh.MinimumCutoff, &jpdh.ActiveBacklogsAllowed, &jpdh.TotalNumberOfBacklogsAllowed, &jpdh.EduGaps11N12Allowed, &jpdh.EduGapsGradAllowed, &jpdh.EduGapsSchoolAllowed, &jpdh.EduGapsPGAllowed, &jpdh.YearOfPassing, &jpdh.Remarks)
		// if err != nil {
		// 	customError.ErrTyp = "S3PJ003"
		// 	customError.ErrCode = "500"
		// 	customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByPID)
		// 	return customError, resp, ""
		// }
		resp = map[string]interface{}{"HC": pd}
		return customError, resp, "HC"
	} else if prf {
		resp = map[string]interface{}{"Profile": pd}
		return customError, resp, "Profile"
	} else if oi {
		// var jpdh OtherInformationSubModel
		// getByPID, _ := RetriveSP("CRP_GET_OI_BY_PID")
		// getByPID = strings.ReplaceAll(getByPID, "34", fmt.Sprint('"'))
		// err := Db.QueryRow(getByPID, publishID).Scan(&jpdh.Title, &jpdh.Information, &jpdh.Attachment)
		// if err != nil {
		// 	customError.ErrTyp = "S3PJ003"
		// 	customError.ErrCode = "500"
		// 	customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByPID)
		// 	return customError, resp, ""
		// }
		//jpdh.Attachment = []byte(fmt.Sprintf("%s", jpdh.tempAttach))
		resp = map[string]interface{}{"OI": pd}
		return customError, resp, "OI"
	}

	return customError, resp, ""
}

func validateSubscription(publishID string, subscriber string, subType string) (DbModelError, bool) {
	rsp := ""
	successResp := map[string]string{}
	var customError DbModelError
	switch subType {
	case "Corporate":
		rsp = "CRP"
		break
	case "University":
		rsp = "UNV"
		break
	case "Student":
		rsp = "STU"
		break
	default:
		customError = DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid Stakeholder type", Err: fmt.Errorf("" + subType + " is invalid,  Expecting Corporate,University or Student"), SuccessResp: successResp}
		return customError, false
	}
	valid := false
	shSP, _ := RetriveSP(rsp + "_VRF_SUB")
	err := Db.QueryRow(shSP, subscriber, publishID).Scan(&valid)
	if err != nil {
		fmt.Println(err)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Failed to get Publisher details of ID %s", publishID)
		return customError, false
	}
	customError.ErrTyp = "000"
	return customError, valid

}
