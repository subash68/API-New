// Package models ...
package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Insert ....
func (subs *SubscriptionReq) Insert(stakeholder string, subscriberID string, pubStakeholder string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	customError.ErrTyp = "000"
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	rsp := ""
	shRsp := ""
	switch stakeholder {
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
		customError = DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid Stakeholder type", Err: fmt.Errorf("" + stakeholder + " is invalid,  Expecting Corporate,University or Student"), SuccessResp: successResp}
		Job <- customError
		return Job
	}
	switch pubStakeholder {
	case "Corporate":
		shRsp = "CRP"
		break
	case "University":
		shRsp = "UNV"
		break
	case "Student":
		shRsp = "STU"
		break
	default:
		customError = DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid Stakeholder type", Err: fmt.Errorf("" + pubStakeholder + " is invalid,  Expecting Corporate,University or Student"), SuccessResp: successResp}
		Job <- customError
		return Job
	}
	currentTime := time.Now()

	// Getting Publisher ID
	shSP, _ := RetriveSP(shRsp + "_GET_SH_PUB")
	fmt.Println(shSP, "===SHSP====", subs.PublishID)
	err := Db.QueryRow(shSP, subs.PublishID).Scan(&subs.Publisher)
	if err != nil {
		fmt.Println(err)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Failed to get Publisher details of ID %s", subs.PublishID)
		Job <- customError
		return Job
	}

	newSubIns, _ := RetriveSP(rsp + "_SUB_INS")

	vals := []interface{}{}
	newSubIns += "(?,?,?,?,?),"

	vals = append(vals, subscriberID, subs.Publisher, currentTime, subs.PublishID, subs.TransactionID)
	newSubIns = newSubIns[0 : len(newSubIns)-1]
	fmt.Println(newSubIns)
	fmt.Printf("\n%+v\n", vals)
	subInsStmt, err := Db.Prepare(newSubIns)
	if err != nil {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare  Subscription insert due to %v %v", newSubIns, err.Error())
		Job <- customError
		return Job
	}

	_, err = subInsStmt.Exec(vals...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert into Subscription database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	customError.ErrTyp = "000"
	customError.SuccessResp = successResp
	Job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return Job
}

// GetAllSubscriptions ...
func (subs *AllSubscriptionsModel) GetAllSubscriptions(ID string, stakeholder string) <-chan DbModelError {
	rsp := ""
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	publisherType := ""
	switch stakeholder {
	case "Corporate":
		rsp = "CRP"
		publisherType = "UNIV"
		break
	case "University":
		rsp = "UNV"
		publisherType = "CORP"
		break
	case "Student":
		rsp = "STU"
		publisherType = "CORP"
		break
	default:
		customError = DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid Stakeholder type", Err: fmt.Errorf("" + stakeholder + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp}
		Job <- customError
		return Job
	}
	// Getting Publisher ID
	shSP, _ := RetriveSP(rsp + "_GET_ALL_SUBS")
	rows, err := Db.Query(shSP, ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Failed to get Subscriptions for ID %s", ID)
		Job <- customError
		return Job
	} else if err == sql.ErrNoRows {
		customError.ErrTyp = "000"
		Job <- customError
		return Job
	}
	defer rows.Close()
	for rows.Next() {
		var newsub SubscriptionModel
		err = rows.Scan(&newsub.Publisher, &newsub.DateOfSubscription, &newsub.PublishID, &newsub.TransactionID, &newsub.CorporateName, &newsub.GeneralNote)
		newsub.GeneralNote = strings.Split(newsub.GeneralNote, " has been published")[0]
		newsub.PublisherType = publisherType
		newsub.SubscriptionType = parseSubscriptionType(publisherType[:1] + newsub.GeneralNote)
		if err != nil {
			customError.ErrTyp = "500"
			customError.ErrCode = "S3PJ002"
			customError.Err = fmt.Errorf("Failed to get Subscriptions for ID %s", ID)
			Job <- customError
			return Job
		}
		subs.Subscriptions = append(subs.Subscriptions, newsub)
	}
	subSP, _ := RetriveSP("UNV_STU_DB_SUB_GET_ALL_SUB")
	fmt.Println("========================== UNV_STU_DB_SUB_GET_ALL==========", subSP)
	subrow, err := Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		Job <- customError
		return Job
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionModel
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Publisher, &newsub.Subscriber, &newsub.DateOfSubscription, &newsub.CorporateName, &newsub.PublisherLocation, &newsub.SearchCriteria)
			newsub.GeneralNote = "Student Database" // strings.Split(newsub.GeneralNote, " has been published")[0]
			newsub.PublisherLocation = newsub.PublisherLocation[:len(newsub.PublisherLocation)-2]
			newsub.PublisherType = "UNIV"
			newsub.SubscriptionType = parseSubscriptionType(newsub.GeneralNote)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				Job <- customError
				return Job
			}
			subs.Subscriptions = append(subs.Subscriptions, newsub)
		}
	}

	subSP, _ = RetriveSP("UNV_INSIGHTS_GET_ALL_SUB")
	fmt.Println("========================== UNV_INSIGHTS_GET_ALL==========", subSP)
	subrow, err = Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		Job <- customError
		return Job
	} else if err == sql.ErrNoRows {

	} else {
		fmt.Println("=============== ")
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionModel
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Publisher, &newsub.Subscriber, &newsub.DateOfSubscription, &newsub.CorporateName, &newsub.PublisherLocation)
			newsub.GeneralNote = "University Information"
			newsub.PublisherLocation = newsub.PublisherLocation[:len(newsub.PublisherLocation)-2]
			newsub.PublisherType = "UNIV"
			newsub.SubscriptionType = parseSubscriptionType(newsub.GeneralNote)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				Job <- customError
				return Job
			}
			subs.Subscriptions = append(subs.Subscriptions, newsub)
		}
	}
	subSP, _ = RetriveSP("CORP_HCI_GET_ALL_SUB")
	fmt.Println("========================== CORP_HCI_GET_ALL_SUB==========", subSP)
	subrow, err = Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		Job <- customError
		return Job
	} else if err == sql.ErrNoRows {
	} else {
		fmt.Println("=============== ")
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionModel
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Publisher, &newsub.Subscriber, &newsub.DateOfSubscription, &newsub.CorporateName, &newsub.PublisherLocation)
			newsub.GeneralNote = "Hiring Insights"
			newsub.PublisherLocation = newsub.PublisherLocation[:len(newsub.PublisherLocation)-2]
			newsub.PublisherType = "CORP"
			newsub.SubscriptionType = parseSubscriptionType(newsub.GeneralNote)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				Job <- customError
				return Job
			}
			subs.Subscriptions = append(subs.Subscriptions, newsub)
		}
	}
	subSP, _ = RetriveSP("CORP_CD_GET_ALL")
	fmt.Println("========================== CORP_CD_GET_ALL==========", subSP, ID)
	subrow, err = Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		Job <- customError
		return Job
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionModel
			var cdReq, cdAr bool
			var rqDate, arDate time.Time
			var reqNftID, arNftID string
			err = subrow.Scan(&newsub.Subscriber, &newsub.Publisher, &newsub.CampusDriveID, &cdReq, &rqDate, &reqNftID, &cdAr, &arDate, &arNftID, &newsub.CorporateName, &newsub.PublisherLocation)
			newsub.GeneralNote = "Campus Hiring" // strings.Split(newsub.GeneralNote, " has been published")[0]
			newsub.PublisherType = "CORP"
			newsub.SubscriptionType = parseSubscriptionType(newsub.GeneralNote)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				Job <- customError
				return Job
			}
			fmt.Printf("\n\n==== Campus details %+v , cdr %v , adAr %v===\n\n", newsub, cdReq, cdAr)
			if cdReq == true && cdAr == true && arNftID != "" {
				newsub.CampusDriveStatus = "Accepted"
				newsub.NftID = arNftID
				newsub.DateOfSubscription = rqDate
			} else if cdReq == true && cdAr == false && arNftID != "" {
				newsub.CampusDriveStatus = "Rejected"
				newsub.NftID = arNftID
				newsub.DateOfSubscription = rqDate
			} else if cdReq == true && cdAr == false && arNftID == "" {
				newsub.CampusDriveStatus = "Pending"
				newsub.NftID = reqNftID
				newsub.DateOfSubscription = rqDate
			}
			subs.Subscriptions = append(subs.Subscriptions, newsub)
		}
	}
	subSP, _ = RetriveSP("UNV_CD_GET_ALL")
	fmt.Println("========================== UNV_CD_GET_ALL==========", subSP, ID)
	subrow, err = Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		Job <- customError
		return Job
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionModel
			var cdReq, cdAr bool
			var rqDate, arDate time.Time
			var reqNftID, arNftID string
			err = subrow.Scan(&newsub.Subscriber, &newsub.Publisher, &newsub.CampusDriveID, &cdReq, &rqDate, &reqNftID, &cdAr, &arDate, &arNftID, &newsub.CorporateName, &newsub.PublisherLocation)
			newsub.GeneralNote = "Campus Hiring" // strings.Split(newsub.GeneralNote, " has been published")[0]
			newsub.PublisherType = "UNIV"
			newsub.SubscriptionType = parseSubscriptionType(newsub.GeneralNote)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				Job <- customError
				return Job
			}
			fmt.Printf("\n\n==== Campus details %+v , cdr %v , adAr %v===\n\n", newsub, cdReq, cdAr)
			if cdReq == true && cdAr == true && arNftID != "" {
				newsub.CampusDriveStatus = "Accepted"
				newsub.NftID = arNftID
				newsub.DateOfSubscription = rqDate
			} else if cdReq == true && cdAr == false && arNftID != "" {
				newsub.CampusDriveStatus = "Rejected"
				newsub.NftID = arNftID
				newsub.DateOfSubscription = rqDate
			} else if cdReq == true && cdAr == false && arNftID == "" {
				newsub.CampusDriveStatus = "Pending"
				newsub.NftID = reqNftID
				newsub.DateOfSubscription = rqDate
			}
			subs.Subscriptions = append(subs.Subscriptions, newsub)
		}
	}
	customError.ErrTyp = "000"
	Job <- customError
	return Job

}

// GetCDForSH ...
func (subs *AllCdInvites) GetCDForSH(ID string, stakeholder string) DbModelError {
	var sentReqSp, receivedReqSp string
	var customError DbModelError
	switch stakeholder {
	case "Corporate":
		sentReqSp = "CORP_CD_GET_ALL"
		receivedReqSp = "UNV_CD_RETURN_RECEIVED"
		break
	case "University":
		sentReqSp = "UNV_CD_GET_ALL"
		receivedReqSp = "CORP_CD_RETURN_RECEIVED"
		break
	default:
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Invalid user Type %s", stakeholder)
		return customError
	}

	subSP, _ := RetriveSP(sentReqSp)
	fmt.Println("========================== sentReqSp==========", subSP, ID)
	subrow, err := Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		return customError
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub CDSubscriptionReq
			var cdReq, cdAr bool
			var rqDate, arDate time.Time
			var reqNftID, arNftID string
			err = subrow.Scan(&newsub.InitiatorID, &newsub.ReceiverID, &newsub.CampusDriveID, &cdReq, &rqDate, &reqNftID, &cdAr, &arDate, &arNftID, &newsub.ReceiverName, &newsub.ReceiverLocation)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				return customError
			}
			fmt.Printf("\n\n==== Campus details %+v , cdr %v , adAr %v===\n\n", newsub, cdReq, cdAr)
			if cdReq == true && cdAr == true && arNftID != "" {
				newsub.CampusDriveStatus = "Accepted"
				newsub.NftID = arNftID
				newsub.RequestedDate = rqDate
				newsub.CampusDriveRequested = true
				newsub.ResponseDate = arDate.Format(time.RFC3339)
			} else if cdReq == true && cdAr == false && arNftID != "" {
				newsub.CampusDriveStatus = "Rejected"
				newsub.NftID = arNftID
				newsub.RequestedDate = rqDate
				newsub.CampusDriveRequested = true
				newsub.ResponseDate = arDate.Format(time.RFC3339)
			} else if cdReq == true && cdAr == false && arNftID == "" {
				newsub.CampusDriveStatus = "Pending"
				newsub.NftID = reqNftID
				newsub.RequestedDate = rqDate
				newsub.CampusDriveRequested = true
			}
			if !newsub.CampusDriveRequested {
				newsub.CampusDriveStatus = "subscribed"
				newsub.RequestedDate = rqDate
			}
			subs.CampusInviteSent = append(subs.CampusInviteSent, newsub)
		}
	}
	subSP, _ = RetriveSP(receivedReqSp)
	fmt.Println("========================== receivedReqSp==========", subSP, ID)
	subrow, err = Db.Query(subSP, ID)
	if err != nil && err != sql.ErrNoRows {
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		return customError
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub CDSubscriptionReq
			var cdReq, cdAr bool
			var rqDate, arDate time.Time
			var reqNftID, arNftID string
			err = subrow.Scan(&newsub.InitiatorID, &newsub.ReceiverID, &newsub.CampusDriveID, &cdReq, &rqDate, &reqNftID, &cdAr, &arDate, &arNftID, &newsub.InitiatorName, &newsub.InitiatorLocation)
			if err != nil {
				customError.ErrTyp = "500"
				customError.ErrCode = "S3PJ002"
				customError.Err = fmt.Errorf("Cannot read the Rows %v", err.Error())
				return customError
			}
			fmt.Printf("\n\n==== Campus details %+v , cdr %v , adAr %v===\n\n", newsub, cdReq, cdAr)
			if cdReq == true && cdAr == true && arNftID != "" {
				newsub.CampusDriveStatus = "Accepted"
				newsub.NftID = arNftID
				newsub.RequestedDate = rqDate
				newsub.CampusDriveRequested = true
				newsub.ResponseDate = arDate.Format(time.RFC3339)
			} else if cdReq == true && cdAr == false && arNftID != "" {
				newsub.CampusDriveStatus = "Rejected"
				newsub.NftID = arNftID
				newsub.RequestedDate = rqDate
				newsub.CampusDriveRequested = true
				newsub.ResponseDate = arDate.Format(time.RFC3339)
			} else if cdReq == true && cdAr == false && arNftID == "" {
				newsub.CampusDriveStatus = "Pending"
				newsub.NftID = reqNftID
				newsub.RequestedDate = rqDate
				newsub.CampusDriveRequested = true
			}
			if newsub.NftID != "" {
				subs.CampusInviteReceived = append(subs.CampusInviteReceived, newsub)
			}
		}
	}
	customError.ErrTyp = "000"
	return customError
}
