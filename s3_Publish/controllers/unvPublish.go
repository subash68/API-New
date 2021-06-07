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

// UPNotificationType ...
const (
	UPNotificationType    string = "UniversityProfile"
	UPNotificationTypeID  string = "4"
	UOINotificationType   string = "UniversityOtherInformation"
	UOINotificationTypeID string = "5"
)

// PublishProfile ...
func PublishProfile(c *gin.Context) {
	successResp = map[string]string{}
	//var jc models.FullJobDb
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Publish University Proposal")
	defer cancel()
	defer close(jobdb)
	var up models.UnvPublishDBModel
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
		go func() {
			select {
			case insertJobChan := <-up.Publish(ID.(string)):
				jobdb <- insertJobChan
			case <-ctx.Done():
				return
			}
		}()

		insertJob := <-jobdb
		fmt.Printf("\n insertjob: %+v\n", insertJob)
		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		reqBody := map[string]string{"senderID": up.StakeholderID, "senderUserRole": "University", "notificationType": UPNotificationType, "content": "Profile has been published", "publishFlag": "true", "publishID": insertJob.SuccessResp["publishID"], "isGeneric": "true", "notificationTypeID": UPNotificationTypeID}
		resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
		if err != nil {
			fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
		}
		fmt.Println(resp)
		c.JSON(http.StatusOK, PubHCResp{"Profile has been Published", insertJob.SuccessResp["publishID"]})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3UNVPJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// PublishUnvOI ...
func PublishUnvOI(c *gin.Context) {
	successResp = map[string]string{}
	//var jc models.FullJobDb
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Publish Other information")
	defer cancel()
	defer close(jobdb)
	var up models.UnvOtherInformationModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&up, binding.Form)
	if err == nil {
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
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			up.Attachment = byteContainer
			up.AttachmentName = file.Filename
		}
		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		up.StakeholderID = ID.(string)
		go func() {
			select {
			case insertJobChan := <-up.PublishOI():
				jobdb <- insertJobChan
			case <-ctx.Done():
				return
			}
		}()
		insertJob := <-jobdb
		fmt.Printf("\n insertjob: %+v\n", insertJob)
		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		reqBody := map[string]string{"senderID": up.StakeholderID, "senderUserRole": "University", "notificationType": UOINotificationType, "content": "Other Information has been published", "publishFlag": "true", "publishID": insertJob.SuccessResp["publishID"], "isGeneric": "true", "notificationTypeID": UOINotificationTypeID}
		resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
		if err != nil {
			fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
		}
		fmt.Println(resp)
		//
		c.JSON(http.StatusOK, PubHCResp{"Information with title" + up.Title + " has been Published", insertJob.SuccessResp["publishID"]})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3UNVPJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetPublishedUnvOI ...
func GetPublishedUnvOI(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.UnvOtherInformationModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Other information published")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetAllOI("UNV_OI_GET_ALL_PUB")
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Hiring Criteria", Err: fmt.Errorf("Cannot find Hiring criteria : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n oi : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}

// GetUnvOI ...
func GetUnvOI(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.UnvOtherInformationModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Other information published")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetAllOI("UNV_OI_GET_ALL")
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Hiring Criteria", Err: fmt.Errorf("Cannot find Hiring criteria : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n oi : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}

// GetPublishedUnvData ...
func GetPublishedUnvData(c *gin.Context) {
	successResp = map[string]string{}
	var up models.UnvPublishDBModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Other information published")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	up.StakeholderID = ID.(string)
	upArray, err := up.GetAllPublishedData()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Hiring Criteria", Err: fmt.Errorf("Cannot find Hiring criteria : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n Published Data : %+v\n", upArray)

	c.JSON(http.StatusOK, upArray)
	return
}

// GetUnvPublishedDataByID ...
func GetUnvPublishedDataByID(c *gin.Context) {
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
	if userType.(string) == "University" {
		customError, resp, respQry := models.GetUnvPublishedData(pubID, true, ID.(string), userType.(string))
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
	res := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("User type is not University"), SuccessResp: successResp})

	c.JSON(http.StatusUnprocessableEntity, res)
	return
}
