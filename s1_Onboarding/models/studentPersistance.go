// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"time"

	// Blank import for initializing
	_ "github.com/go-sql-driver/mysql"
)

// Insert Corporate Table ....
func (data *StudentMasterDb) Insert(expiryDate string) <-chan DbModelError {

	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Verify as a new User
	var studentExists bool
	dbSP, _ := RetriveSP("STU_EXISTS_WITH_EMAIL")
	err := Db.QueryRow(dbSP, data.CollegeEmail).Scan(&data.StakeholderID, &data.CollegeEmail, &studentExists)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("query operation failed" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT001", fmt.Errorf("Cannot Read Database %v ", err.Error()), successResp,
		}
		return Job
	}
	//fmt.Printf(" 49 %v  %+v\n ", studentExists, err)

	// Return if already exists
	if studentExists {
		Job <- DbModelError{
			"403", "S1AUT002", fmt.Errorf("User already Signed Up with this CollegeEmail"), successResp,
		}
		return Job

	}
	sID, cbError := createStuSID(data.DateOfBirth)
	if cbError.ErrCode != "000" {
		Job <- cbError
		return Job
	}
	data.StakeholderID = sID
	fmt.Println(data.StakeholderID)
	// Prepare Db Insert
	dbSP, _ = RetriveSP("STU_INS_NEW_USR")
	stmt, err := Db.Prepare(dbSP)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT003", fmt.Errorf("Error While registering Student %v ", err.Error()), successResp,
		}
		return Job
	}
	defer stmt.Close()

	results, err := stmt.Exec(&data.StakeholderID, &data.FirstName, &data.MiddleName, &data.LastName, &data.PersonalEmail, &data.CollegeEmail, &data.PhoneNumber, &data.AlternatePhoneNumber, &data.CollegeID, &data.Gender, &data.DateOfBirth, &data.AadharNumber, &data.PermanentAddressLine1, &data.PermanentAddressLine2, &data.PermanentAddressLine3, &data.PermanentAddressCountry, &data.PermanentAddressState, &data.PermanentAddressCity, &data.PermanentAddressDistrict, &data.PermanentAddressZipcode, &data.PermanentAddressPhone, &data.PresentAddressLine1, &data.PresentAddressLine2, &data.PresentAddressLine3, &data.PresentAddressCountry, &data.PresentAddressState, &data.PresentAddressCity, &data.PresentAddressDistrict, &data.PresentAddressZipcode, &data.PresentAddressPhone, &data.FathersGuardianFullName, &data.FathersGuardianOccupation, &data.FathersGuardianCompany, &data.FathersGuardianPhoneNumber, &data.FathersGuardianEmailID, &data.MothersGuardianFullName, &data.MothersGuardianOccupation, &data.MothersGuardianCompany, &data.MothersGuardianDesignation, &data.MothersGuardianPhoneNumber, &data.MothersGuardianEmailID, &data.AccountStatus, &data.Password, expiryDate)
	fmt.Printf("reults: %+v \n %+v", results, err)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT004", fmt.Errorf("Error While registering Student %v ", err.Error()), successResp,
		}
		return Job
	}

	// Print data in Console
	fmt.Printf("line 80 %+v %+v \n ", data, err)

	customError.ErrTyp = "000"
	successResp["Phone"] = data.CollegeEmail
	successResp["StakeholderID"] = data.StakeholderID
	successResp["Email"] = data.PhoneNumber
	customError.SuccessResp = successResp

	Job <- customError

	return Job

}

func createStuSID(dob string) (string, DbModelError) {

	fmt.Printf("\n ---> dob: %v yoe: %v\n", dob)
	rowCount := 0

	lutSP, _ := RetriveSP("STU_ROW_CNT")
	err := Db.QueryRow(lutSP).Scan(&rowCount)
	if err != nil {
		return "", DbModelError{
			"500", "", fmt.Errorf("Failed to Create Platform Unique ID, due to db connection error"), map[string]string{},
		}
	}
	layout := "2006-01-02"
	dobYear, err := time.Parse(layout, dob)
	if err != nil {
		return "", DbModelError{
			"500", "", fmt.Errorf("Invalid Dob string, expected format yyyy-mm-dd"), map[string]string{},
		}
	}

	partialID := fmt.Sprintf("%010d", (rowCount + 1))
	return fmt.Sprint("S", dobYear.Year(), partialID), DbModelError{
		"000", "", nil, map[string]string{},
	}
}

// UpdateVrfStatus ...
func (data *StudentMasterDb) UpdateVrfStatus() <-chan DbModelError {
	Job := make(chan DbModelError, 1)

	customError := UpdateVrfInfoInDB(data.StakeholderID, "", data.PrimaryPhoneVerified, data.PrimaryEmailVerified, "STU_MBL_VRF_QRY", "STU_ACC_STATUS_UPD")

	Job <- customError

	return Job

}

// UpdateAccountStatus ...
func (data *StudentMasterDb) UpdateAccountStatus(expiryDate string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)

	customError := ActivateAccount(data.StakeholderID, expiryDate, "STU_ACC_ACTIVATION")

	Job <- customError

	return Job

}

// GetContactInfo ...
func (data *StudentMasterDb) GetContactInfo(sid string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	customError := FetchContactInfo(sid, "STU_VRF_ME_QRY")
	Job <- customError
	return Job
}

// Login ...
func (data *StudentMasterDb) Login(userID string, password string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	customError := LoginStakehodler(userID, password, "STU_LOGIN")
	Job <- customError
	return Job
}
