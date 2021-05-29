// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"strconv"
	// Blank import for initializing
	//_ "github.com/go-sql-driver/mysql"
)

// CorporateData model for reference

// Insert Corporate Table ....
func (data *CorporateMasterDB) Insert(expiryDate string) <-chan DbModelError {

	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Verify as a new User
	var CorporateExists bool
	dbSP, _ := RetriveSP("CORP_EXISTS_WITH_EMAIL")
	err := Db.QueryRow(dbSP, data.PrimaryContactEmail).Scan(&data.StakeholderID, &data.PrimaryContactEmail, &CorporateExists)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("query operation failed" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT001", fmt.Errorf("Cannot Read Database %v ", err.Error()), successResp,
		}
		return Job
	}
	//fmt.Printf(" 49 %v  %+v\n ", CorporateExists, err)

	// Return if already exists
	if CorporateExists {
		Job <- DbModelError{
			"403", "S1AUT002", fmt.Errorf("Corporate already Signed Up with this Email"), successResp,
		}
		return Job

	}

	sID, cbError := createCrpSID(data.CorporateType, data.CorporateCategory, strconv.FormatInt(data.YearOfEstablishment, 10))
	if cbError.ErrCode != "000" {
		Job <- cbError
		return Job
	}
	data.StakeholderID = sID
	fmt.Println(data.StakeholderID)
	// Prepare Db Insert
	dbSP, _ = RetriveSP("CORP_INS_NEW_USR")
	stmt, err := Db.Prepare(dbSP)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT003", fmt.Errorf("Error While registering Corporate %v ", err.Error()), successResp,
		}
		return Job
	}
	defer stmt.Close()

	stmt.Exec(data.StakeholderID, data.CorporateName, data.CIN, data.CorporateHQAddressLine1, data.CorporateHQAddressLine2, data.CorporateHQAddressLine3, data.CorporateHQAddressCountry, data.CorporateHQAddressState, data.CorporateHQAddressCity, data.CorporateHQAddressDistrict, data.CorporateHQAddressZipCode, data.CorporateHQAddressPhone, data.CorporateHQAddressEmail, data.CorporateLocalBranchAddressLine1, data.CorporateLocalBranchAddressLine2, data.CorporateLocalBranchAddressLine3, data.CorporateLocalBranchAddressCountry, data.CorporateLocalBranchAddressState, data.CorporateLocalBranchAddressCity, data.CorporateLocalBranchAddressDistrict, data.CorporateLocalBranchAddressZipCode, data.CorporateLocalBranchAddressPhone, data.CorporateLocalBranchAddressEmail, data.PrimaryContactFirstName, data.PrimaryContactMiddleName, data.PrimaryContactLastName, data.PrimaryContactDesignation, data.PrimaryContactPhone, data.PrimaryContactEmail, data.SecondaryContactFirstName, data.SecondaryContactMiddleName, data.SecondaryContactLastName, data.SecondaryContactDesignation, data.SecondaryContactPhone, data.SecondaryContactEmail, data.CorporateType, data.CorporateCategory, data.CorporateIndustry, data.CompanyProfile, data.Attachment, data.AttachmentName, data.YearOfEstablishment, data.AccountStatus, data.Password, expiryDate)

	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT004", fmt.Errorf("Error While inseting Corporate Table %v ", err.Error()), successResp,
		}
		return Job
	}

	// Print data in Console
	fmt.Printf("line 80 %+v %+v \n ", data, err)

	customError.ErrTyp = "000"
	successResp["Phone"] = data.PrimaryContactPhone
	successResp["StakeholderID"] = data.StakeholderID
	successResp["Email"] = data.PrimaryContactEmail
	customError.SuccessResp = successResp

	Job <- customError

	return Job

}

// UpdateVrfStatus ...
func (data *CorporateMasterDB) UpdateVrfStatus() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError

		return Job
	}
	successResp := map[string]string{}
	mobileVerfied, emailVerified, vrfDataExists := false, false, false

	dbSP, _ := RetriveSP("CORP_MBL_VRF_QRY")
	fmt.Printf("\n(((---- > ))) %+v\n", data)

	err := Db.QueryRow(dbSP, data.StakeholderID).Scan(&emailVerified, &mobileVerfied, &data.AccountStatus, &vrfDataExists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("query operation failed" + err.Error())
		Job <- DbModelError{
			"500", "S1VRF001", fmt.Errorf("Cannot Read Database %v ", err.Error()), successResp,
		}
		return Job
	}
	fmt.Printf("\nmv: %v , ev: %v, de : %v\n", mobileVerfied, emailVerified, vrfDataExists)
	if (data.PrimaryPhoneVerified && data.PrimaryPhoneVerified != mobileVerfied) || mobileVerfied {
		data.PrimaryPhoneVerified = true
	}
	if (data.PrimaryEmailVerified && data.PrimaryEmailVerified != emailVerified) || emailVerified {
		data.PrimaryEmailVerified = true

	}
	if data.PrimaryEmailVerified && data.PrimaryPhoneVerified {
		data.AccountStatus = "2"
	}

	dbSP, _ = RetriveSP("CORP_ACC_STATUS_UPD")
	stmt, err := Db.Prepare(dbSP)
	defer stmt.Close()
	if err != nil {
		fmt.Printf("\n ***-> %+v\n", err)
		customError.Err = fmt.Errorf("Account verification failed : %v", err.Error())
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	}
	result, err := stmt.Exec(data.PrimaryEmailVerified, data.PrimaryPhoneVerified, data.AccountStatus, data.StakeholderID)
	fmt.Printf("line 200 persis %v \n ", result)
	if err != nil {
		fmt.Printf("\n ***-> %+v\n", err)
		customError.Err = fmt.Errorf("Account verification failed : %v", err.Error())
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	} else {
		customError.ErrTyp = "000"
		successResp["AccountStatus"] = data.AccountStatus
	}

	successResp["emailVerified"] = fmt.Sprintf("%v", (emailVerified || data.PrimaryEmailVerified))
	successResp["mobileVerfied"] = fmt.Sprintf("%v", (mobileVerfied || data.PrimaryPhoneVerified))
	customError.SuccessResp = successResp

	Job <- customError

	return Job

}

// UpdateAccountStatus ...
func (data *CorporateMasterDB) UpdateAccountStatus(expiryDate string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)

	customError := ActivateAccount(data.StakeholderID, expiryDate, "CORP_ACC_ACTIVATION")

	Job <- customError

	return Job

}

// GetContactInfo ...
func (data *CorporateMasterDB) GetContactInfo(sid string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	customError := FetchContactInfo(sid, "CORP_VRF_ME_QRY")
	Job <- customError
	return Job
}

// VerifyAccountToken ...
func (data *CorporateMasterDB) VerifyAccountToken() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		Job <- customError

		return Job
	}
	dbSP, _ := RetriveSP("CORP_ACC_STATUS_UPD")
	fmt.Printf("\n(((---- > ))) %+v\n", data)
	stmt, err := Db.Prepare(dbSP)
	defer stmt.Close()
	if err != nil {
		customError.Err = fmt.Errorf("Account verification failed")
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	}
	result, err := stmt.Exec("PAYMENT_PENDING", data.StakeholderID, "VRF_PENDING")
	fmt.Printf("line 200 persis %v \n ", result)
	if err != nil {
		customError.Err = fmt.Errorf("Account verification failed")
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	} else {
		customError.ErrTyp = "000"
	}

	Job <- customError

	return Job

}

func createCrpSID(crTyp string, crCat string, estYear string) (string, DbModelError) {

	fmt.Printf("\n ---> crptype : %v , crpcat: %v yoe: %v\n", crTyp, crCat, estYear)
	lutSP, _ := RetriveSP("LUT_GET_CRP_TYPE")
	crpTypeExists, crpCatExists := false, false
	rowCount := 0
	//var crTyp, crCat string
	err := Db.QueryRow(lutSP, crTyp).Scan(&crTyp, &crpTypeExists)

	lutSP, _ = RetriveSP("LUT_GET_CRP_CAT")
	err = Db.QueryRow(lutSP, crCat).Scan(&crCat, &crpCatExists)

	lutSP, _ = RetriveSP("CORP_ROW_CNT")
	err = Db.QueryRow(lutSP).Scan(&rowCount)

	fmt.Printf("\n ---> updated crptype : %v , crpcat: %v yoe: %v \n", crTyp, crCat, estYear)
	fmt.Printf("\n ---> got  crptype : %v , crpcat: %v yoe: %v \n", crpTypeExists, crpCatExists, rowCount)
	if err != nil || !crpTypeExists || !crpCatExists {
		return "", DbModelError{
			"500", "S1AUT004", fmt.Errorf("Invalid Corporate Type / Sector  %+v", err), map[string]string{},
		}
	}
	partialID := fmt.Sprintf("%010d", (rowCount + 1))
	return fmt.Sprint("C", crTyp, crCat, estYear, partialID), DbModelError{
		"000", "S1AUT004", nil, map[string]string{},
	}
}

// Login ...
func (data *CorporateMasterDB) Login(userID string, password string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	customError := LoginStakehodler(userID, password, "CORP_LOGIN")
	Job <- customError
	return Job
}

// // Login ...
// func (data *CorporateMasterDB) Login(userID string, password string) <-chan DbModelError {
// 	Job := make(chan DbModelError, 1)
// 	var customError DbModelError
// 	if CheckPing(&customError); customError.Err != nil {
// 		Job <- customError
// 		return Job
// 	}

// 	// Verify as a new User
// 	var corporateExists bool
// 	customError.ErrCode = ""
// 	dbSP, _ := RetriveSP("CORP_LOGIN_WITH_EMAIL")
// 	fmt.Printf("\n Email: %v , password: %v \n", data.PrimaryContactEmail, data.Password)
// 	err := Db.QueryRow(dbSP, userID, userID, userID).Scan(&data.StakeholderID, &data.AccountStatus, &data.Password, &corporateExists)
// 	if err != nil || err == sql.ErrNoRows {
// 		customError.Err = fmt.Errorf("User details not Found err: %v ", err.Error())
// 		customError.ErrCode = "S1LGN001"
// 		customError.ErrTyp = "500"
// 		fmt.Println("query operation failed" + err.Error())
// 		Job <- customError

// 		return Job
// 	}
// 	if data.AccountStatus == "VRF_PENDING" {
// 		customError.Err = fmt.Errorf("User is not verified")
// 		customError.ErrCode = "S1LGN002"
// 		customError.ErrTyp = "403"
// 		Job <- customError

// 		return Job
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
// 	if err != nil {
// 		customError.Err = fmt.Errorf("Password Mismatch")
// 		customError.ErrCode = "S1LGN003"
// 		customError.ErrTyp = "403"
// 		fmt.Println("query operation failed" + err.Error())
// 		Job <- customError

// 		return Job
// 	}

// 	if data.AccountStatus == "ACTIVE" {
// 		customError.ErrCode = "/dashboard"
// 	} else {
// 		customError.ErrCode = "/payment"
// 	}
// 	customError.ErrTyp = "000"
// 	Job <- customError

// 	return Job

// }
