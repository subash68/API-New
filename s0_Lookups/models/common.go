package models

import (
	"fmt"
)

// CheckPing Checks if connection exists with Database
func CheckPing(customError *DbModelError) {
	err := Db.Ping()
	if err != nil {
		InitDataModel()
		err1 := Db.Ping()
		if err1 != nil {
			customError.Err = fmt.Errorf("Error While connecting CORPORATE Table %w ", err)
			customError.ErrCode = "S1AUT912"
			customError.ErrTyp = "500"
			customError.SuccessResp = map[string]string{}
			fmt.Printf(" line 21 %+v ", customError)
		}
	}
	return
}
