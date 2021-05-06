// Package controllers ...
package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/configuration"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/middleware"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/models"
	"gopkg.in/go-playground/validator.v8"
)

// ctxFunc declaration for context use
type ctxFunc string

// ctxkey
var ctxkey ctxFunc

// respErr this struct will be returned as a JSON object
var respErr models.RespErrModel

// detailedErr will be returned along with respErr
var detailedErr models.DetailedErrModel

// DtlErr ...
type DtlErr models.DetailedErrModel

//Convert ...
func (data DtlErr) Convert() *gin.Error {

	var loger gin.Error

	loger.Err = errors.New(data.Message)
	loger.Type = gin.ErrorTypePrivate
	loger.Meta = data

	return &loger

}

var successResp map[string]string

// Dmr ...
type Dmr models.DbModelError

// String  ...
func (err Dmr) String() models.DetailedErrModel {
	fmt.Printf("%+v", err)
	var errormessage models.DetailedErrModel

	errormessage.Code = err.ErrCode
	errormessage.Message = err.Err.Error()

	return errormessage

}

// ErrCheck ...
func ErrCheck(ctx context.Context, result models.DbModelError) models.RespErrModel {

	var respErrMsg models.RespErrModel
	var detailErrs DtlErr
	target, _ := ctx.Value(ctxkey).(string)
	var errormessage models.ErrorMessages
	_, ok := result.Err.(validator.ValidationErrors)
	if !ok {
		fmt.Println("V")
	}
	customError := Dmr(result).String()
	respErrMsg.Message = result.ErrTyp
	respErrMsg.Target = target
	respErrMsg.Code = customError.Code
	if ok {
		for _, fieldErr := range result.Err.(validator.ValidationErrors) {

			ErrorConverter := models.FieldError{*fieldErr}.String()

			errormessage.Err.Detail = append(errormessage.Err.Detail, ErrorConverter)

		}
		respErrMsg.Errors = append(respErrMsg.Errors, errormessage.Err.Detail...)
		return respErrMsg
	}

	detailErrs.Code = customError.Code
	detailErrs.Target = "All"
	detailErrs.Message = result.Err.Error()
	respErrMsg.Errors = append(respErrMsg.Errors, models.DetailError(detailErrs))

	return respErrMsg

}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GetRandomID ...
func GetRandomID(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))

}

func getFuncReq(c *gin.Context, ctxKey string) (context.Context, string, string, map[string]string) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, ctxkey)

	defer cancel()
	ID, ok := c.Get("userID")

	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6Profile001", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return ctx, "", "", successResp
	}
	fmt.Println("-----> Got ID", ID.(string))
	userType, ok := c.Get("userType")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile001", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User Type from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return ctx, "", "", successResp
	}
	return ctx, ID.(string), userType.(string), successResp
}

func makeNFTServiceCall(endpoint string, reqData map[string]string) (string, error) {
	tokenConfig := configuration.NftConfig()
	resBody, err := middleware.MakeInternalServiceCall(tokenConfig.Host, tokenConfig.Port, "POST", endpoint, reqData)
	if err != nil {
		return "", err
	}
	var tokenResp TokenDbResp
	err = json.Unmarshal(resBody, &tokenResp)
	if err != nil {
		return "", err
	}
	return tokenResp.NftID, nil
}
