package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// InsertToken ...
func (tkn *TokenMasterDB) InsertToken() <-chan DbModelError {
	job := make(chan DbModelError, 1)
	fmt.Println("insert started")
	var customError DbModelError
	at := time.Unix(tkn.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()
	if CheckRedisPing(&customError); customError.Err != nil {
		job <- customError
		return job
	}
	fmt.Println("ping done")
	byteClaims, err := json.Marshal(TokenClaims{TokenType: tkn.TokenType, UserType: tkn.UserType, UserID: tkn.UserID})
	if err != nil {
		customError.Err = fmt.Errorf("Cannot load add claims to the token %v ", err.Error())
		job <- customError
		return job
	}
	errAccess := RedisClient.Set(tkn.AccessUUID, byteClaims, at.Sub(now)).Err()
	fmt.Println(errAccess)
	if errAccess != nil {
		customError.Err = fmt.Errorf("Error While inserting token to Redis %v ", errAccess.Error())
		customError.ErrCode = "S1TOK001"
		customError.ErrTyp = "403"
		job <- customError
		return job
	}
	customError.ErrTyp = "000"
	job <- customError
	fmt.Println("insert done")
	return job
}

// GetClaimsForToken ...
func (tkn *TokenMasterDB) GetClaimsForToken() (TokenClaims, error) {
	var customError DbModelError
	if CheckRedisPing(&customError); customError.Err != nil {
		return TokenClaims{}, fmt.Errorf("Auth database not working")
	}
	fmt.Printf("\n tkn : %+v\n", tkn)
	usrInfo, err := RedisClient.Get(tkn.AccessToken).Result()
	fmt.Printf("\n User INFO : %+v\n", usrInfo)

	if err == redis.Nil || err != nil || usrInfo == "" {
		return TokenClaims{}, fmt.Errorf("Invalid token, nil db %v", err.Error())
	}
	var tokenClaims TokenClaims
	err = json.Unmarshal([]byte(usrInfo), &tokenClaims)
	if err != nil {
		return TokenClaims{}, fmt.Errorf("Invalid token marshaling : %v", err.Error())
	}
	fmt.Printf("\n get : %+v\n", tokenClaims)

	return tokenClaims, nil
}

// DeleteToken ...
func (tkn *TokenMasterDB) DeleteToken() bool {
	var customError DbModelError
	if CheckRedisPing(&customError); customError.Err != nil {
		return false
	}
	usr, err := RedisClient.Del(tkn.AccessToken).Result()
	fmt.Printf("\n---> %+v\n", usr)
	if err != nil {
		return false
	}
	return true
}
