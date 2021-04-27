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

// AddPJResp ...
type AddPJResp struct {
	PjID []string `json:"pjID"`
}

// AddPublishedJobs ...
func AddPublishedJobs(c *gin.Context) {
	successResp = map[string]string{}
	var pj models.PublishJobs
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Adding Published Job")

	defer cancel()
	defer close(jobdb)
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&pj, binding.Form)
	if err == nil {

		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		go func() {
			select {
			case insertJobChan := <-pj.Insert(ID.(string)):
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
		pubIDs := insertJob.SuccessResp["pjIDs"]
		pubIDs = pubIDs[1 : len(pubIDs)-1]
		arrayPubIDs := strings.Split(pubIDs, " ")
		for i := 0; i < len(pj.PublishedJobs); i++ {
			reqBody := map[string]string{"senderID": ID.(string), "senderUserRole": "Corporate", "notificationType": "General", "content": "New Job has been published", "publishFlag": "true", "publishID": arrayPubIDs[i]}
			resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
			if err != nil {
				fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
			}
			fmt.Println(resp)
		}

		c.JSON(http.StatusOK, AddPJResp{arrayPubIDs})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetPublishedJobsByID ...
func GetPublishedJobsByID(c *gin.Context) {
	successResp = map[string]string{}
	var pj models.PublishedJobsDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get Published Job")

	defer cancel()
	defer close(jobdb)

	pj.PublishID = c.Param("pjID")
	if pj.PublishID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Piblished Job ID, require pjID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	go func() {
		select {
		case insertJobChan := <-pj.GetByID():
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

	fmt.Printf("\n PJ : %+v\n", pj)

	c.JSON(http.StatusOK, pj)
	return
}

// GetAllPublishedJobs ...
func GetAllPublishedJobs(c *gin.Context) {
	successResp = map[string]string{}
	var pj models.PublishedJobsDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Published Jobs")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	pj.StakeholderID = ID.(string)
	pjArray, err := pj.GetAllPJ()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Published Jobs created", Err: fmt.Errorf("Cannot find Published Jobs : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n PJ : %+v\n", pjArray)

	c.JSON(http.StatusOK, pjArray)
	return
}

// UpdatePublishedJobs ...
func UpdatePublishedJobs(c *gin.Context) {
	successResp = map[string]string{}
	var pj models.PublishedJobsDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Updating Published Job")

	defer cancel()
	defer close(jobdb)
	pj.PublishID = c.Param("pjID")
	if pj.PublishID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Piblished Job ID, require pjID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string), c.Request.PostFormValue("jobID"))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	pj.StakeholderID = ID.(string)
	updateReq := c.Request.PostForm
	customError := models.UpdatePublishedData(updateReq, "PJ_UPDATE_BY_ID", "PJ_UPDATE_WHERE", pj.StakeholderID, pj.PublishID, nil, "")
	if customError.ErrTyp != "000" {
		resp := ErrCheck(ctx, customError)
		c.Error(customError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, "Successfully updated")
	return
}

// DeletePublishedJobs ...
func DeletePublishedJobs(c *gin.Context) {
	successResp = map[string]string{}
	var pj models.PublishedJobsDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Delete Published Job")

	defer cancel()
	defer close(jobdb)

	pj.PublishID = c.Param("pjID")
	if pj.PublishID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find pjID"), SuccessResp: successResp})
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
	pj.StakeholderID = ID.(string)
	go func() {
		select {
		case insertJobChan := <-pj.DeleteByID():
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

	fmt.Printf("\n pj : %+v\n", pj)

	c.JSON(http.StatusOK, DelHCResp{"Published Job With ID :" + pj.PublishID + " has been delete"})
	return
}
