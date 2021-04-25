// Package models ...
package models

import (
	"fmt"
)

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
