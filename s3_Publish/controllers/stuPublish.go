// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// PublishStudentProfile ...
func PublishStudentProfile(c *gin.Context) {
	successResp = map[string]string{}
	//var jc models.FullJobDb

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Publish Student Profile")
	defer cancel()
	var up models.StuPublishDBModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&up, binding.Form)

	if err == nil {
		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		up.DateOfPublish = time.Now()
		up.StakeholderID = ID.(string)
		up.GeneralNote = "Profile has been published"
		publishID, err := up.Publish()

		fmt.Printf("\n insertjob: %+v\n", err)
		var errresp models.DbModelError
		if err != nil {
			errresp.ErrCode = "S3STUPUB"
			errresp.Err = err
			errresp.ErrTyp = "Database Insert Error"
			resp := ErrCheck(ctx, errresp)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		//
		c.JSON(http.StatusOK, PubHCResp{"Profile has been Published", publishID})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3STUPUB", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// PublishStudentOI ...
func PublishStudentOI(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Publish Student Other information")
	defer cancel()
	var up models.StuOtherInformationModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&up, binding.Form)

	if err == nil {
		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		up.StakeholderID = ID.(string)
		// up.GeneralNote = "Profile has been published"
		form, _ := c.MultipartForm()
		files := form.File["attachment"]
		if len(files) > 1 {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Upload multiple files is not supported"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		for _, file := range files {
			fileContent, _ := file.Open()
			byteContainer, err := ioutil.ReadAll(fileContent)
			if err != nil {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Reading uploaded files err %v", err), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			up.Attachment = byteContainer
		}
		publishID, err := up.PublishOtherInfo()

		fmt.Printf("\n insertjob: %+v\n", err)
		var errresp models.DbModelError
		if err != nil {
			errresp.ErrCode = "S3STUPUB"
			errresp.Err = err
			errresp.ErrTyp = "Database Insert Error"
			resp := ErrCheck(ctx, errresp)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		//
		c.JSON(http.StatusOK, PubHCResp{"Profile has been Published", publishID})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3STUPUB", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllStudentOI ...
func GetAllStudentOI(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.StuOtherInformationModel

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Student Other Information")

	defer cancel()

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetAllOI()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Other Information", Err: fmt.Errorf("Cannot find Other information : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n OI : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}

// GetStudentPublishedData ...
func GetStudentPublishedData(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.StuPublishDBModel

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Student Other Information")

	defer cancel()

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetAllPublishHistory()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Other Information", Err: fmt.Errorf("Cannot find Other information : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n OI : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}

// GetStudentPublishedDataByID ...
func GetStudentPublishedDataByID(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.StuPublishDBModel

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Student Other Information")

	defer cancel()

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	pubID := c.Param("publishID")
	if pubID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Publish ID in params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	customErr, oiArray, queryStr := models.GetStuPublishedDataByID(pubID, true, "", "")
	if customErr.Err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Other Information", Err: customErr.Err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n OI : %+v\n", oiArray)

	_, ok = oiArray[queryStr].(string)
	fmt.Printf("\n%+v\n", oiArray)
	if !ok {
		c.JSON(http.StatusOK, oiArray[queryStr])
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, strings.ReplaceAll(oiArray[queryStr].(string), "\\\"", "\""))
	c.Abort()
	return
}
