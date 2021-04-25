package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/configuration"
)

// AddPayment ...
func (pay *PaymentDbModel) AddPayment(userType string) DbModelError {

	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		return customError
	}

	sp, _ := RetriveSP("Payment_Insert")
	dbConfig := configuration.DbConfig()
	replaceStr := ""
	switch userType {
	case "Corporate":
		replaceStr = dbConfig.CrpPayDbName
		break
	case "University":
		replaceStr = dbConfig.UnvPayDbName
		break
	case "Student":
		replaceStr = dbConfig.StuPayDbName
		break
	default:
		fmt.Println("Invalid userType specified")
		return DbModelError{
			"500", "S7PG002", fmt.Errorf("Invalid Usertype %v" + userType), successResp,
		}
	}
	sp = strings.ReplaceAll(sp, "RPLCE", replaceStr)
	fmt.Println(sp)
	stmt, err := Db.Prepare(sp)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		return DbModelError{
			"500", "S7PG001", fmt.Errorf("Error While While Preparing Payment insertion %v ", err.Error()), successResp,
		}
	}

	defer stmt.Close()
	currentTime := time.Now()

	var parsedPayAmount float64
	parsedPayAmount, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", pay.PayedAmount/100), 64)

	stmt.Exec(pay.StakeholderID, parsedPayAmount, pay.PaymentID, pay.PaymentMode, currentTime)

	if err != nil {

		fmt.Println("error while inserting " + err.Error())
		return DbModelError{
			"500", "S7PG002", fmt.Errorf("Error While inseting Payment Table %v ", err.Error()), successResp,
		}
	}

	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	return customError
}

// GetPayments ...
func (pays *AllPaymentsModel) GetPayments(ID string, userType string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	sp, _ := RetriveSP("GET_Payments_BY_ID")
	dbConfig := configuration.DbConfig()
	replaceStr := ""
	switch userType {
	case "Corporate":
		replaceStr = dbConfig.CrpPayDbName
		break
	case "University":
		replaceStr = dbConfig.UnvPayDbName
		break
	case "Student":
		replaceStr = dbConfig.StuPayDbName
		break
	default:
		fmt.Println("Invalid userType specified")
		Job <- DbModelError{
			"500", "S7PG002", fmt.Errorf("Invalid Usertype %v" + userType), successResp,
		}
		return Job
	}
	sp = strings.ReplaceAll(sp, "RPLCE", replaceStr)
	fmt.Println(sp)

	rows, err := Db.Query(sp, ID)
	if err != nil {

		fmt.Println("error while Fetching payments " + err.Error())
		Job <- DbModelError{
			"500", "S7PG003", fmt.Errorf("Error While Getting payments %v ", err.Error()), successResp,
		}
		return Job
	}

	defer rows.Close()

	for rows.Next() {
		var newPay PaymentDbModel
		err = rows.Scan(&newPay.PayedAmount, &newPay.PaymentID, &newPay.PaymentMode, &newPay.PayedDate)
		if err != nil {
			customError.ErrTyp = "S3PJ003"
			customError.ErrCode = "500"
			customError.Err = fmt.Errorf("Cannot read the Allocation  Rows %v", err.Error())
			Job <- customError
			return Job
		}
		pays.Payments = append(pays.Payments, newPay)
	}

	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError
	fmt.Println("========================================== Payment inserted")
	return Job
}
