// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// HCNotificationType ...
const (
	HCNotificationType   string = "HiringCriteria"
	HCNotificationTypeID string = "1"
)

// AddHCResp ...
type AddHCResp struct {
	HcID string `json:"hcID"`
}

// DelHCResp ...
type DelHCResp struct {
	Message string `json:"message"`
}

// PubHCResp ...
type PubHCResp struct {
	Message   string `json:"message"`
	PublishID string `json:"publishID"`
}

// AddOIResp ...
type AddOIResp struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// AddHiringCriteria ...
func AddHiringCriteria(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.MultipleHC
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Creating Hiring Criteria")
	// lc := c.Bind()
	defer cancel()
	defer close(jobdb)
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}

	binding.Validator = &defaultValidator{}

	err = c.ShouldBindWith(&hc, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))

	fmt.Printf("\nhc: %+v\n", c.Request.PostForm)
	if len(hc.HiringCriterias) <= 0 {
		err = fmt.Errorf("Require HiringCriterias in Array format")
	}
	if err == nil {

		ID, ok := c.Get("userID")
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		fmt.Printf("\n=====================> HC: %+v\n", hc.HiringCriterias)
		go func() {
			select {
			case insertJobChan := <-hc.Insert(ID.(string)):
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
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, AddHCResp{insertJob.SuccessResp["hcIDs"]})
		c.Abort()
		return
	}

	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}

// GetHiringCriteriaByID ...
func GetHiringCriteriaByID(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.HiringCriteriaDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get Hiring Criteria")

	defer cancel()
	defer close(jobdb)

	hc.HiringCriteriaID = c.Param("hcID")
	if hc.HiringCriteriaID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot fing Hiring criteria ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	go func() {
		select {
		case insertJobChan := <-hc.GetByID():
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

	fmt.Printf("\n HC : %+v\n", hc)

	c.JSON(http.StatusOK, hc)
	return
}

// GetAllHiringCriteria ...
func GetAllHiringCriteria(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.HiringCriteriaDB
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
	hc.StakeholderID = ID.(string)
	hcArray, err := hc.GetAllHC("HC_GET_ALL")
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Hiring Criteria", Err: fmt.Errorf("Cannot find Hiring criteria : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n HC : %+v\n", hcArray)

	c.JSON(http.StatusOK, hcArray)
	return
}

// GetAllPublishedHC ...
func GetAllPublishedHC(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.HiringCriteriaDB
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
	hc.StakeholderID = ID.(string)
	hcArray, err := hc.GetAllHC("HC_GET_ALL_PUB")
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Hiring Criteria", Err: fmt.Errorf("Cannot find Hiring criteria : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n HC : %+v\n", hcArray)

	c.JSON(http.StatusOK, hcArray)
	return
}

// UpdateHiringCriteria ...
func UpdateHiringCriteria(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.HiringCriteriaDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Updating Hiring Criteria")

	defer cancel()
	defer close(jobdb)
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "InvalidContentType", Err: fmt.Errorf("Content type is not 'application/json'"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}

	hc.HiringCriteriaID = c.Param("hcID")
	if hc.HiringCriteriaID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot fing Hiring criteria ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string), c.Request.PostFormValue("mi"))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	binding.Validator = &defaultValidator{}
	err = c.ShouldBindWith(&hc, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))
	if err == nil {

		hc.StakeholderID = ID.(string)
		customError := hc.Update()
		if customError.ErrTyp != "000" {
			resp := ErrCheck(ctx, customError)
			c.Error(customError.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, DelHCResp{"Successfully updated"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteHiringCriteria ...
func DeleteHiringCriteria(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.HiringCriteriaDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Delete Hiring Criteria")

	defer cancel()
	defer close(jobdb)

	hc.HiringCriteriaID = c.Param("hcID")
	if hc.HiringCriteriaID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot fing Hiring criteria ID"), SuccessResp: successResp})
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
	hc.StakeholderID = ID.(string)
	go func() {
		select {
		case insertJobChan := <-hc.DeleteByID():
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

	fmt.Printf("\n HC : %+v\n", hc)

	c.JSON(http.StatusOK, DelHCResp{"Hiring Criteria With ID :" + hc.HiringCriteriaID + " Has been delete"})
	return
}

// PublishHiringCriteria ...
func PublishHiringCriteria(c *gin.Context) {
	successResp = map[string]string{}
	var hc models.PublishHiringCriteriasModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Delete Hiring Criteria")

	defer cancel()
	defer close(jobdb)
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		err = fmt.Errorf("Invalid content type %s , Required %s", reqContentType, "application/json")
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	binding.Validator = &defaultValidator{}
	err = c.ShouldBindWith(&hc, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))
	if err == nil {
		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		hc.StakeholderID = ID.(string)
		go func() {
			select {
			case insertJobChan := <-hc.PublishHC():
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

		reqBody := map[string]string{"senderID": hc.StakeholderID, "senderUserRole": "Corporate", "notificationType": HCNotificationType, "content": "Hiring Criteria has been published", "publishFlag": "true", "publishID": insertJob.SuccessResp["publishID"], "isGeneric": "true", "notificationTypeID": HCNotificationTypeID}
		resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
		if err != nil {
			fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
		}
		fmt.Println(resp)

		fmt.Printf("\n HC : %+v\n", hc)

		c.JSON(http.StatusOK, PubHCResp{"Hiring Criteria has been Published", insertJob.SuccessResp["publishID"]})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}
