// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// GetCrpSubscribedData ...
func GetCrpSubscribedData(c *gin.Context) {
	successResp = map[string]string{}
	//jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get Corporate Publish Deatails")

	defer cancel()
	//defer close(jobdb)

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	userType, ok := c.Get("userType")
	fmt.Println("-----> Got ID", userType.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode userType from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	pubID := c.Param("publishID")
	if pubID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Publish ID in params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	if userType.(string) == "Corporate" {
		customError, resp, respQry := models.GetUnvPublishedData(pubID, false, ID.(string), userType.(string))
		if customError.ErrTyp != "000" {
			res := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: customError.Err, SuccessResp: successResp})

			c.JSON(http.StatusUnprocessableEntity, res)
			return
		}
		_, ok := resp[respQry].(string)
		fmt.Printf("\n%+v\n", resp)
		if !ok {
			c.JSON(http.StatusOK, resp[respQry])
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, strings.ReplaceAll(resp[respQry].(string), "\\\"", "\""))
		c.Abort()
		return
	}
	customError, resp, respQry := models.GetCrpPublishedData(pubID, false, ID.(string), userType.(string))
	if customError.ErrTyp != "000" {
		res := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: customError.Err, SuccessResp: successResp})

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	_, ok = resp[respQry].(string)
	fmt.Printf("\n%+v\n", resp)
	if !ok {
		c.JSON(http.StatusOK, resp[respQry])
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, strings.ReplaceAll(resp[respQry].(string), "\\\"", "\""))
	c.Abort()
	return

}

// GetCrpPublishedDataByID ...
func GetCrpPublishedDataByID(c *gin.Context) {
	successResp = map[string]string{}
	//jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get Corporate Publish Deatails")

	defer cancel()
	//defer close(jobdb)

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information1", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	userType, ok := c.Get("userType")
	fmt.Println("-----> Got ID", userType.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information2", Err: fmt.Errorf("Cannot decode userType from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	pubID := c.Param("publishID")
	if pubID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information3", Err: fmt.Errorf("Cannot find Publish ID in params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	if userType.(string) == "Corporate" {
		customError, resp, respQry := models.GetCrpPublishedData(pubID, true, ID.(string), userType.(string))
		fmt.Println(customError)
		if customError.ErrTyp != "000" {
			res := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information4", Err: customError.Err, SuccessResp: successResp})

			c.JSON(http.StatusUnprocessableEntity, res)
			c.Abort()
			return
		}
		_, ok := resp[respQry].(string)
		fmt.Printf("\n%+v\n", resp)
		if !ok {
			c.JSON(http.StatusOK, resp[respQry])
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, strings.ReplaceAll(resp[respQry].(string), "\\\"", "\""))
		c.Abort()
		return
	}
	res := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information5", Err: fmt.Errorf("User type is not Corporate"), SuccessResp: successResp})

	c.JSON(http.StatusUnprocessableEntity, res)
	return
}

// GetAllCrpPublishedData ...
func GetAllCrpPublishedData(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.CorpPushedDataModel

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Corporate Other Information")

	defer cancel()

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetCrpPublishedDataByID()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Other Information", Err: fmt.Errorf("Cannot find Other information : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n OI : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}
