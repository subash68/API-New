// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"strconv"

	// Blank import for initializing
	_ "github.com/go-sql-driver/mysql"
)

// Insert Corporate Table ....
func (data *UniversityMasterDb) Insert(expiryDate string) <-chan DbModelError {

	fmt.Printf("\n data : %+v \n", data)
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Verify as a new User
	var universityExists bool
	dbSP, _ := RetriveSP("UNV_EXISTS_WITH_EMAIL")
	err := Db.QueryRow(dbSP, data.PrimaryContactEmail).Scan(&data.StakeholderID, &data.PrimaryContactEmail, &universityExists)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("query operation failed" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT001", fmt.Errorf("Cannot Read Database %v ", err.Error()), successResp,
		}
		return Job
	}
	//fmt.Printf(" 49 %v  %+v\n ", universityExists, err)

	// Return if already exists
	if universityExists {
		Job <- DbModelError{
			"403", "S1AUT002", fmt.Errorf("University already Signed Up with this Email"), successResp,
		}
		return Job

	}

	sID, cbError := createUnvSID(data.UniversitySector, strconv.FormatInt(data.YearOfEstablishment, 10))
	if cbError.ErrCode != "000" {
		Job <- cbError
		return Job
	}
	data.StakeholderID = sID
	fmt.Println(data.StakeholderID)

	// Prepare Db Insert
	dbSP, _ = RetriveSP("UNV_INS_NEW_USR")
	stmt, err := Db.Prepare(dbSP)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT003", fmt.Errorf("Error While registering University %v ", err.Error()), successResp,
		}
		return Job
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.StakeholderID, data.UniversityName, data.UniversityCollageID, data.UniversityCollageName, data.UniversityHQAddressLine1, data.UniversityHQAddressLine2, data.UniversityHQAddressLine3, data.UniversityHQAddressCountry, data.UniversityHQAddressState, data.UniversityHQAddressCity, data.UniversityHQAddressDistrict, data.UniversityHQAddressZipcode, data.UniversityHQAddressPhone, data.UniversityHQAddressemail, data.UniversityLocalBranchAddressLine1, data.UniversityLocalBranchAddressLine2, data.UniversityLocalBranchAddressLine3, data.UniversityLocalBranchAddressCountry, data.UniversityLocalBranchAddressState, data.UniversityLocalBranchAddressCity, data.UniversityLocalBranchAddressDistrict, data.UniversityLocalBranchAddressZipcode, data.UniversityLocalBranchAddressPhone, data.UniversityLocalBranchAddressemail, data.PrimaryContactFirstName, data.PrimaryContactMiddleName, data.PrimaryContactLastName, data.PrimaryContactDesignation, data.PrimaryContactPhone, data.PrimaryContactEmail, data.SecondaryContactFirstName, data.SecondaryContactMiddleName, data.SecondaryContactLastName, data.SecondaryContactDesignation, data.SecondaryContactPhone, data.SecondaryContactEmail, data.UniversitySector, data.UniversityProfile, data.YearOfEstablishment, data.Attachment, data.AttachmentName, data.AccountStatus, data.Password, expiryDate, data.UniversityID)

	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S1AUT004", fmt.Errorf("Error While registering University %v ", err.Error()), successResp,
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
func (data *UniversityMasterDb) UpdateVrfStatus() <-chan DbModelError {
	Job := make(chan DbModelError, 1)

	customError := UpdateVrfInfoInDB(data.StakeholderID, "", data.PrimaryPhoneVerified, data.PrimaryEmailVerified, "UNV_MBL_VRF_QRY", "UNV_ACC_STATUS_UPD")

	Job <- customError

	return Job

}

// UpdateAccountStatus ...
func (data *UniversityMasterDb) UpdateAccountStatus(expiryDate string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)

	customError := ActivateAccount(data.StakeholderID, expiryDate, "UNV_ACC_ACTIVATION")

	Job <- customError

	return Job

}

// GetContactInfo ...
func (data *UniversityMasterDb) GetContactInfo(sid string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	customError := FetchContactInfo(sid, "UNV_VRF_ME_QRY")
	Job <- customError
	return Job
}

// Login ...
func (data *UniversityMasterDb) Login(userID string, password string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	customError := LoginStakehodler(userID, password, "UNV_LOGIN")
	Job <- customError
	return Job
}

func createUnvSID(unvCat string, estYear string) (string, DbModelError) {

	fmt.Printf("\n ---> unvcat: %v yoe: %v\n", unvCat, estYear)
	lutSP, _ := RetriveSP("LUT_GET_UNV_CAT")
	unvCatExists := false
	rowCount := 0
	//var crTyp, crCat string
	err := Db.QueryRow(lutSP, unvCat).Scan(&unvCat, &unvCatExists)

	lutSP, _ = RetriveSP("UNV_ROW_CNT")
	err = Db.QueryRow(lutSP).Scan(&rowCount)

	fmt.Printf("\n ---> updated unvcat: %v yoe: %v \n", unvCat, estYear)
	fmt.Printf("\n ---> got   unvcat: %v yoe: %v \n", unvCatExists, rowCount)
	if err != nil || !unvCatExists {
		return "", DbModelError{
			"500", "S1AUT004", fmt.Errorf("Invalid University Sector  %+v", err), map[string]string{},
		}
	}
	partialID := fmt.Sprintf("%010d", (rowCount + 1))
	return fmt.Sprint("U", unvCat, estYear, partialID), DbModelError{
		"000", "", nil, map[string]string{},
	}
}
