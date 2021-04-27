// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// ProfilePublish ...
func ProfilePublish(c *gin.Context) {
	successResp = map[string]string{}
	var pubData models.PublishDataModel
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Publishing Corporate Profile")

	defer cancel()
	defer close(jobdb)
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&pubData, binding.Form)

	if err != nil || pubData.PublishData == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Profile Publish Data"), SuccessResp: successResp})
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
	pubData.StakeholderID = ID.(string)
	go func() {
		select {
		case insertJobChan := <-pubData.Insert():
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
	reqBody := map[string]string{"senderID": pubData.StakeholderID, "senderUserRole": "Corporate", "notificationType": "General", "content": "Profile Information has been published", "publishFlag": "true", "publishID": insertJob.SuccessResp["publishID"]}
	resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
	if err != nil {
		fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
	}
	fmt.Println(resp)

	fmt.Printf("\n Profile : %+v\n", pubData)

	c.JSON(http.StatusOK, PubHCResp{"Profile has been Published", insertJob.SuccessResp["publishID"]})
	return
}
