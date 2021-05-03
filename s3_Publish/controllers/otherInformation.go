// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// OINotificationType ...
const (
	OINotificationType   string = "OtherInformation"
	OINotificationTypeID string = "2"
)

// AddOtherInfo ...
func AddOtherInfo(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.OtherInformationModel

	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Creating OtherInformation")

	defer cancel()
	defer close(jobdb)
	// if c.Param("jobID") != "" {
	// 	jc.JobID = c.Param("jobID")
	// }
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&oi, binding.Form)
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
			oi.Attachment = byteContainer
		}

		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		oi.StakeholderID = ID.(string)
		go func() {
			select {
			case insertJobChan := <-oi.Insert():
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
		oiID, _ := strconv.Atoi(insertJob.SuccessResp["insID"])
		c.JSON(http.StatusOK, AddOIResp{"Other information stored", oiID})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllOI ...
func GetAllOI(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.OtherInformationModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Other Information stored")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetAllOI("OI_GET_ALL")
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Other Information", Err: fmt.Errorf("Cannot find Other information : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n OI : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}

// GetAllPublishedOI ...
func GetAllPublishedOI(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.OtherInformationModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Hiring Criteria")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	oiArray, err := oi.GetAllOI("OI_GET_ALL_PUB")
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Other Information", Err: fmt.Errorf("Cannot find Other information : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n OI : %+v\n", oiArray)

	c.JSON(http.StatusOK, oiArray)
	return
}

// PublishOI ...
func PublishOI(c *gin.Context) {
	successResp = map[string]string{}
	var oi models.OtherInformationModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Publish Other information")

	defer cancel()
	defer close(jobdb)
	var err error
	oi.ID, err = strconv.Atoi(c.Param("id"))
	if oi.ID <= 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot Find Other information ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	oi.StakeholderID = ID.(string)
	go func() {
		select {
		case insertJobChan := <-oi.PublishOI():
			jobdb <- insertJobChan
		case <-ctx.Done():
			return
		}
	}()
	insertJob := <-jobdb

	if insertJob.ErrTyp != "000" {
		resp := ErrCheck(ctx, insertJob)
		c.Error(insertJob.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	reqBody := map[string]string{"senderID": oi.StakeholderID, "senderUserRole": "Corporate", "notificationType": OINotificationType, "content": "Other Information has been published", "publishFlag": "true", "publishID": insertJob.SuccessResp["publishID"], "isGeneric": "true", "notificationTypeID": OINotificationTypeID}
	resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
	if err != nil {
		fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
	}
	fmt.Println(resp)

	fmt.Printf("\n OI : %+v\n", oi)

	c.JSON(http.StatusOK, PubHCResp{"Other information has been Published", insertJob.SuccessResp["publishID"]})
	return
}
