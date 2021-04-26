// Package controllers ...
package controllers

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
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

func getInterfaceValues(s ...interface{}) (val []interface{}) {

	v := reflect.ValueOf(s[0])
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if typeOfS.Field(i).Name != "ID" {
			val = append(val, v.Field(i).Interface())
		}
	}
	return val
}
