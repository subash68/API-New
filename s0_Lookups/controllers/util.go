// Package controllers ...
package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s0_Lookups/models"
	services "github.com/jaswanth-gorripati/PGK/s0_Lookups/services"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v8"
)

var log *logrus.Logger = services.Logger

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
	log.Debugf("\n Converting Dmr to string Err values %+v\n", err)
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

func getFuncReq(c *gin.Context, ctxKey string) (context.Context, string, string, map[string]string) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, ctxkey)

	defer cancel()
	ID, ok := c.Get("userID")

	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "s0LUT001", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return ctx, "", "", successResp
	}
	log.Debugf("-----> Got ID", ID.(string))
	userType, ok := c.Get("userType")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "s0LUT002", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User Type from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return ctx, "", "", successResp
	}
	return ctx, ID.(string), userType.(string), successResp
}
