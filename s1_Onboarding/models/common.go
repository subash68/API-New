package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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

// CheckRedisPing Checks if connection exists with Database
func CheckRedisPing(customError *DbModelError) {
	_, err := RedisClient.Ping().Result()
	if err != nil {
		customError.Err = fmt.Errorf("Error While connecting to Redis %v ", err.Error())
		customError.ErrCode = "S1AUT912"
		customError.ErrTyp = "500"
		fmt.Printf(" line 21 %+v ", customError)
	}
	return
}

// UpdateVrfInfoInDB ...
func UpdateVrfInfoInDB(sid string, accStatus string, phoneVrf bool, emailVrf bool, querySP string, insSP string) DbModelError {
	var customError DbModelError

	if CheckPing(&customError); customError.Err != nil {
		return customError
	}
	successResp := map[string]string{}
	mobileVerfied, emailVerified, vrfDataExists := false, false, false

	dbSP, _ := RetriveSP(querySP)
	fmt.Printf("\n(((---- > ))) %v , %v,%v , %v,%v , %v\n", sid, accStatus, phoneVrf, emailVrf, querySP, insSP)

	err := Db.QueryRow(dbSP, sid).Scan(&emailVerified, &mobileVerfied, &accStatus, &vrfDataExists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("query operation failed" + err.Error())
		return DbModelError{
			"500", "S1VRF001", fmt.Errorf("Cannot Read Database %v ", err.Error()), successResp,
		}
	}
	fmt.Printf("\nmv: %v , ev: %v, de : %v\n", mobileVerfied, emailVerified, vrfDataExists)
	if (phoneVrf && phoneVrf != mobileVerfied) || mobileVerfied {
		phoneVrf = true
	}
	if (emailVrf && emailVrf != emailVerified) || emailVerified {
		emailVrf = true

	}
	if emailVrf && phoneVrf {
		accStatus = "2"
	}

	dbSP, _ = RetriveSP(insSP)
	stmt, err := Db.Prepare(dbSP)
	defer stmt.Close()
	if err != nil {
		fmt.Printf("\n ***-> %+v\n", err)
		customError.Err = fmt.Errorf("Account verification failed : %v", err.Error())
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	}
	result, err := stmt.Exec(emailVrf, phoneVrf, accStatus, sid)
	fmt.Printf("line 200 persis %v \n ", result)
	if err != nil {
		fmt.Printf("\n ***-> %+v\n", err)
		customError.Err = fmt.Errorf("Account verification failed : %v", err.Error())
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	} else {
		customError.ErrTyp = "000"
		successResp["AccountStatus"] = accStatus
	}
	customError.SuccessResp = successResp
	return customError
}

// ActivateAccount ...
func ActivateAccount(sid string, expiryDate string, insSP string) DbModelError {
	var customError DbModelError
	successResp := map[string]string{}
	dbSP, _ := RetriveSP(insSP)
	stmt, err := Db.Prepare(dbSP)
	defer stmt.Close()
	if err != nil {
		fmt.Printf("\n ***-> %+v\n", err)
		customError.Err = fmt.Errorf("Account Activation failed: %v", err.Error())
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	}
	result, err := stmt.Exec("A", expiryDate, "2", sid)
	fmt.Printf("line 200 persis %v \n ", result)
	if err != nil {
		fmt.Printf("\n ***-> %+v\n", err)
		customError.Err = fmt.Errorf("Account Activation failed: %v", err.Error())
		customError.ErrCode = "S1AUT008"
		customError.ErrTyp = "500"
	} else {
		customError.ErrTyp = "000"
		successResp["AccountStatus"] = "A"
	}
	customError.SuccessResp = successResp
	return customError
}

// LoginStakehodler ...
func LoginStakehodler(userID string, password string, loginSP string) DbModelError {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return customError
	}

	// Verify as a new User
	var userExists bool
	customError.ErrCode = ""
	accStatus, sid, dbPass := "", "", ""

	dbSP, _ := RetriveSP(loginSP)
	//fmt.Printf("\n Email: %v , password: %v \n", data.PrimaryContactEmail, data.Password)
	err := Db.QueryRow(dbSP, userID, userID, userID).Scan(&sid, &accStatus, &dbPass, &userExists)
	if err != nil || err == sql.ErrNoRows {
		customError.Err = fmt.Errorf("%s is not registered ", userID)
		customError.ErrCode = "S1LGN001"
		customError.ErrTyp = "500"
		fmt.Println("query operation failed" + err.Error())
		return customError
	}
	if accStatus == "1" {
		customError.Err = fmt.Errorf("User is not verified")
		customError.ErrCode = "S1LGN002"
		customError.ErrTyp = "403"
		return customError
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password))
	if err != nil {
		customError.Err = fmt.Errorf("Password Mismatch")
		customError.ErrCode = "S1LGN003"
		customError.ErrTyp = "403"
		fmt.Println("query operation failed" + err.Error())
		return customError
	}
	if accStatus == "A" {
		customError.ErrCode = "/dashboard"
	} else {
		customError.ErrCode = "/payment"
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = map[string]string{"StakeholderID": sid}

	return customError
}

// FetchContactInfo ...
func FetchContactInfo(userID string, qrySP string) DbModelError {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return customError
	}

	// Verify as a new User
	var userExists bool
	var phone, email string
	dbSP, _ := RetriveSP(qrySP)
	//fmt.Printf("\n Email: %v , password: %v \n", data.PrimaryContactEmail, data.Password)
	err := Db.QueryRow(dbSP, userID).Scan(&phone, &email, &userExists)
	if err != nil || err == sql.ErrNoRows {
		customError.Err = fmt.Errorf("User details not Found err: %v ", err.Error())
		customError.ErrCode = "S1VRF201"
		customError.ErrTyp = "500"
		fmt.Println("query operation failed" + err.Error())
		return customError
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = map[string]string{"Phone": phone, "Email": email}
	return customError
}

// GetPlatformUUID ...
func GetPlatformUUID(searchBy string, stakeholder string) (string, DbModelError) {
	var platformUUID string
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return "", customError
	}
	var dbQuery string
	switch stakeholder {
	case "Corporate":
		dbQuery, _ = RetriveSP("CORP_GET_PID")
		break
	case "University":
		dbQuery, _ = RetriveSP("UNV_GET_PID")
		break
	case "Student":
		dbQuery, _ = RetriveSP("STU_GET_PID")
		break
	default:
		customError.ErrTyp = "500"
		customError.ErrCode = "S1AUT"
		customError.Err = fmt.Errorf("" + stakeholder + " is invaild,  Expecting Corporate,University or Student")
		return "", customError
	}
	var userExists bool
	err := Db.QueryRow(dbQuery, searchBy, searchBy).Scan(&platformUUID, &userExists)
	fmt.Printf("\n --> %v,%v\n", platformUUID, userExists)
	if err != nil || err == sql.ErrNoRows || !userExists {
		customError.ErrTyp = "500"
		customError.ErrCode = "S1AUT"
		customError.Err = fmt.Errorf("User not registered with " + searchBy)
		return "", customError
	}
	customError.ErrTyp = "000"
	return platformUUID, customError

}

// ChangePassword ...
func ChangePassword(password string, id string, stakeholder string) DbModelError {
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return customError
	}
	var dbQuery string
	switch stakeholder {
	case "Corporate":
		dbQuery, _ = RetriveSP("CORP_CNG_PASS")
		break
	case "University":
		dbQuery, _ = RetriveSP("UNV_CNG_PASS")
		break
	case "Student":
		dbQuery, _ = RetriveSP("STU_CNG_PASS")
		break
	default:
		customError.ErrTyp = "500"
		customError.ErrCode = "S1AUT"
		customError.Err = fmt.Errorf("" + stakeholder + " is invaild,  Expecting Corporate,University or Student")
		return customError
	}
	stmt, err := Db.Prepare(dbQuery)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S1AUT"
		customError.Err = fmt.Errorf("Failed to Prepare the db  " + err.Error())
		return customError
	}
	_, err = stmt.Exec(password, id)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S1AUT"
		customError.Err = fmt.Errorf("Failed to Exec  the db for changing Password  " + err.Error())
		return customError
	}
	customError.ErrTyp = "000"
	return customError

}
