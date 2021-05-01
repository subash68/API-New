package models

import (
	"fmt"
	"net/url"
	"reflect"
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

// UpdatePublishedData ...
func UpdatePublishedData(updateQuery url.Values, spName string, spExt string, stakeholder string, updateID string, file []byte, dbField string) DbModelError {
	updateString, _ := RetriveSP(spName)
	values := []interface{}{}
	var customError DbModelError
	for key, val := range updateQuery {

		dbKey, exists := GetDbKey(key)
		if !exists {
			customError.ErrTyp = "500"
			customError.ErrCode = "S3UPDT001"
			customError.Err = fmt.Errorf("Invalid key , Cannot update " + key)
			return customError
			fmt.Println("key not exists: ", key)
		} else {
			updateString = updateString + " " + dbKey + "= ?,"
			fmt.Println(updateString)
			if val[0] == "false" {
				values = append(values, 0)
			} else if val[0] == "true" {
				values = append(values, 1)
			} else {
				values = append(values, val[0])
			}
		}

	}
	if file != nil && dbField != "" {
		updateString = updateString + " " + dbField + "= ?,"
		values = append(values, file)
	}
	valLength := reflect.ValueOf(values).Len()
	if valLength == 0 {
		customError.ErrTyp = "000"
		return customError
	}
	values = append(values, updateID)
	values = append(values, stakeholder)
	updateString = updateString[0 : len(updateString)-1]
	whereCond, _ := RetriveSP(spExt)
	updateString = updateString + "" + whereCond

	fmt.Printf("\n keys : %v \nvalues %v\n", updateString, values)

	updateStm, err := Db.Prepare(updateString)
	if err != nil {
		fmt.Println(updateString)
		customError.ErrTyp = "500"
		customError.ErrCode = "S3PJ002"
		customError.Err = fmt.Errorf("Cannot prepare database Prepare due to %v", err.Error())
		return customError
	}
	_, err = updateStm.Exec(values...)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to update the database due to : %v", err.Error())
		customError.ErrCode = "S3PJ002"
		return customError
	}
	customError.ErrTyp = "000"
	return customError
}

// InArray ...
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
