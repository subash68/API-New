// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// Publish ...
func (sp *StuPublishDBModel) Publish() (string, error) {
	// Creating HCID
	var customError DbModelError
	sp.PublishID, customError = CreateUnvPublishID(sp.StakeholderID, "SPDH", "STU_PDH_Get_Last_ID")
	if customError.ErrTyp != "000" || sp.PublishID == "" {
		fmt.Printf("\nFailed to Generate PublishID :%+v\n", customError)
		return "", customError.Err
	}
	if sp.StudentName == "" {
		unvNameCmd, _ := RetriveSP("STU_GET_Name")
		_ = Db.QueryRow(unvNameCmd, sp.StakeholderID).Scan(&sp.StudentName)
	}
	stuPublishCmd, _ := RetriveSP("STU_PDH_INS_NEW")
	stmt, err := Db.Prepare(stuPublishCmd)
	if err != nil {
		return "", fmt.Errorf("Cannot prepare -- %v , -- insert due to %v", stuPublishCmd, err.Error())
	}
	currentTime := time.Now()
	fmt.Printf("\n==============>  %+v <=====================\n", sp)
	_, err = stmt.Exec(sp.StakeholderID, sp.PublishID, sp.StudentName, currentTime, sp.ContactInfoPublished, sp.EducationPublished, sp.LanguagesPublished, sp.CertificationsPublished, sp.AssessmentsPublished, sp.InternshipPublished, sp.OtherInformationPublished, sp.GeneralNote, currentTime, currentTime, sp.PublishedData)
	if err != nil {

		return "", fmt.Errorf("Failed to insert in database -- %v , -- insert due to %v", stuPublishCmd, err.Error())
	}

	return sp.PublishID, nil
}

// PublishOtherInfo ...
func (soi *StuOtherInformationModel) PublishOtherInfo() (string, error) {
	sp := StuPublishDBModel{}
	sp.StakeholderID = soi.StakeholderID
	sp.OtherInformationPublished = true
	sp.PublishedData = "[]"
	sp.GeneralNote = "Other information has been published"
	publishID, err := sp.Publish()
	if err != nil {
		return "", err
	}
	oiInsertCmd, _ := RetriveSP("STU_OI_INS_NEW")
	currentTime := time.Now()

	oiStmt, err := Db.Prepare(oiInsertCmd)
	if err != nil {

		return "", fmt.Errorf("Cannot prepare Other information insert due to %v", err.Error())
	}
	_, err = oiStmt.Exec(soi.StakeholderID, soi.Title, soi.Information, soi.Attachment, true, publishID, currentTime, currentTime)
	if err != nil {

		return "", fmt.Errorf("Failed to insert  Other Information in database due to : %v", err.Error())
	}
	return publishID, nil

}

// GetAllPublishHistory ...
func (sp *StuPublishDBModel) GetAllPublishHistory() ([]StuPublishDBModel, error) {
	var sa []StuPublishDBModel
	getAllHCSP, _ := RetriveSP("STU_PDH_GET_ALL")
	hcRows, err := Db.Query(getAllHCSP, sp.StakeholderID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		return sa, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return sa, nil
	}
	defer hcRows.Close()
	for hcRows.Next() {
		var newSA StuPublishDBModel
		err = hcRows.Scan(&newSA.PublishID, &newSA.StudentName, &newSA.DateOfPublish, &newSA.ContactInfoPublished, &newSA.EducationPublished, &newSA.LanguagesPublished, &newSA.CertificationsPublished, &newSA.AssessmentsPublished, &newSA.InternshipPublished, &newSA.OtherInformationPublished, &newSA.GeneralNote, &newSA.CreationDate, &newSA.LastUpdatedDate, &newSA.PublishedData)
		if err != nil {
			return sa, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		sa = append(sa, newSA)
	}
	return sa, nil
}

// GetAllOI ...
func (soi *StuOtherInformationModel) GetAllOI() (oiArray []StuOtherInformationModel, err error) {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return oiArray, customError.Err
	}
	getAllHCSP, _ := RetriveSP("STU_OI_GET_ALL")
	hcRows, err := Db.Query(getAllHCSP, soi.StakeholderID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		return oiArray, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return oiArray, nil
	}
	defer hcRows.Close()
	for hcRows.Next() {
		var newOI StuOtherInformationModel
		err = hcRows.Scan(&newOI.ID, &newOI.Title, &newOI.Information, &newOI.PublishID, &newOI.Attachment, &newOI.CreationDate, &newOI.LastUpdatedDate, &newOI.PublishedFlag)
		if err != nil {
			return oiArray, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		oiArray = append(oiArray, newOI)
	}
	return oiArray, nil
}

// GetStuPublishedDataByID ...
func GetStuPublishedDataByID(publishID string, isOwner bool, subscriber string, subType string) (DbModelError, map[string]interface{}, string) {
	var resp map[string]interface{}
	var customError DbModelError
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

	var cp, ep, lp, ctp, ap, ip, oip bool
	var pd string

	getByIDSP, _ := RetriveSP("STU_PDH_GET_PID")
	err := Db.QueryRow(getByIDSP, publishID).Scan(&cp, &ep, &lp, &ctp, &ap, &ip, &oip, &pd)
	fmt.Println(getByIDSP)
	if err != nil {
		customError.ErrTyp = "S3PJ003"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByIDSP)
		return customError, resp, ""
	}
	fmt.Println("====================>>>>>>>>>>>>>>>>>", cp, ep, lp, ctp, ap, ip, oip)

	if cp || ep || lp || ctp || ap || ip {
		resp = map[string]interface{}{"Profile": pd}
		fmt.Printf("\n============== %s =================== \n", resp["Jobs"])
		return customError, resp, "Profile"
	} else if oip {
		var jpdh OtherInformationSubModel
		getByPID, _ := RetriveSP("STU_GET_OI_BY_PID")
		getByPID = strings.ReplaceAll(getByPID, "34", fmt.Sprint('"'))
		err := Db.QueryRow(getByPID, publishID).Scan(&jpdh.Title, &jpdh.Information, &jpdh.Attachment)
		if err != nil {
			customError.ErrTyp = "S3PJ003"
			customError.ErrCode = "500"
			customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByPID)
			return customError, resp, ""
		}
		//jpdh.Attachment = []byte(fmt.Sprintf("%s", jpdh.tempAttach))
		resp = map[string]interface{}{"OI": jpdh}
		return customError, resp, "OI"
	}
	return customError, resp, ""
}
