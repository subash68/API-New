package models

import (
	"encoding/json"
	"fmt"
)

// GetUnvPublishedData ...
func GetUnvPublishedData(publishID string, isOwner bool, subscriber string, subType string) (DbModelError, map[string]interface{}, string) {
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
	var pp, bp, ssp, ap, cp, rp, op, prp bool
	var pd string

	getByIDSP, _ := RetriveSP("UNV_PDH_GET_PID")
	err := Db.QueryRow(getByIDSP, publishID).Scan(&pp, &bp, &ssp, &ap, &cp, &rp, &op, &prp, &pd)
	fmt.Println(getByIDSP)
	if err != nil {
		customError.ErrTyp = "S3PJ003"
		customError.ErrCode = "500"
		customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v , %s ", err.Error(), getByIDSP)
		return customError, resp, ""
	}
	fmt.Println("====================>>>>>>>>>>>>>>>>>", pp, bp, ssp, ap, cp, rp, op, prp)

	if pp || bp || ssp || ap || cp || rp || prp {
		pd1 := pd[1:]
		pd = pd1[:len(pd)-2] //+ pd1[:1]
		resp = map[string]interface{}{"Profile": pd}
		fmt.Printf("\n============== %s =================== \n", resp["Jobs"])
		return customError, resp, "Profile"
	} else if op {
		jpdh := OtherInformationSubModel{}
		// getByPID, _ := RetriveSP("UNV_GET_OI_BY_PID")
		// getByPID = strings.ReplaceAll(getByPID, "34", fmt.Sprint('"'))
		// err := Db.QueryRow(getByPID, publishID).Scan(&jpdh.Title, &jpdh.Information, &jpdh.Attachment)
		err := json.Unmarshal([]byte(pd), &jpdh)
		if err != nil {
			customError.ErrTyp = "S3PJ003"
			customError.ErrCode = "500"
			customError.Err = fmt.Errorf("Failed to retrieve Published Data : %v ", err.Error())
			return customError, resp, ""
		}
		//jpdh.Attachment = []byte(fmt.Sprintf("%s", jpdh.tempAttach))
		resp = map[string]interface{}{"OI": jpdh}
		return customError, resp, "OI"
	}
	return customError, resp, ""
}
