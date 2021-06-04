package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// StudentEvents ...
var (
	StudentEvents studentEvents = studentEvents{}
)

type studentEvents struct{}

// AddEvents ...
func (saw *studentEvents) AddEvents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Events")
	binding.Validator = &defaultValidator{}
	var sa models.StudentEventsModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID
		form, _ := c.MultipartForm()
		files := form.File["attachment"]
		if len(files) <= 0 {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: fmt.Errorf("Require Attachment file to be uploaded"), SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		for _, file := range files {
			fileContent, _ := file.Open()
			byteContainer, err := ioutil.ReadAll(fileContent)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				c.Abort()
				return
			}
			sa.Attachment = byteContainer
			sa.AttachmentName = file.Filename
		}
		currentTime := time.Now()
		sa.CreationDate = currentTime
		sa.LastUpdatedDate = currentTime
		sa.EnabledFlag = true
		vals := []interface{}{sa.StakeholderID, sa.Name, sa.Date, sa.Attachment, sa.AttachmentName, sa.OrganizedBy, sa.OrganizedByEmail, sa.OrganizedByPhone, sa.EventType, sa.EventTypeOther, sa.EventResult, sa.EventResultOther, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_Event_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Events Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllEvents ...
func (saw *studentEvents) GetAllEvents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Events")

	sa, err := getEvents(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Events", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.Events)
	c.Abort()
	return

}

func getEvents(ID string) (models.StudentAllEventsModel, error) {
	var sa models.StudentAllEventsModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_Event_GETALL", ID)
	if err != nil {
		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentEventsModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.Date, &newSl.Attachment, &newSl.AttachmentName, &newSl.OrganizedBy, &newSl.OrganizedByEmail, &newSl.OrganizedByPhone, &newSl.EventType, &newSl.EventTypeOther, &newSl.EventResult, &newSl.EventResultOther, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate, &newSl.SentforVerification, &newSl.DateSentforVerification, &newSl.Verified, &newSl.DateVerified, &newSl.SentbackforRevalidation, &newSl.DateSentBackForRevalidation, &newSl.ValidatorRemarks, &newSl.VerificationType, &newSl.VerifiedByStakeholderID, &newSl.VerifiedByEmailID)
		if err != nil {
			return sa, err
		}
		sa.Events = append(sa.Events, newSl)
	}
	return sa, nil
}

// UpdateEvents ...
func (saw *studentEvents) UpdateEvents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update Events")
	binding.Validator = &defaultValidator{}
	var sa models.StudentEventsModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID
		form, _ := c.MultipartForm()
		files := form.File["attachment"]
		for _, file := range files {
			fileContent, _ := file.Open()
			byteContainer, err := ioutil.ReadAll(fileContent)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				c.Abort()
				return
			}
			sa.Attachment = byteContainer
			sa.AttachmentName = file.Filename
		}
		sa.ID, err = strconv.Atoi(c.Param("id"))
		if sa.ID <= 0 || err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Events ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateEvents()
		err := models.StudentInfoService.UpdateStudentInfo("STU_Event_UPD", []interface{}{sa.Name, sa.Date, sa.Attachment, sa.AttachmentName, sa.OrganizedBy, sa.OrganizedByEmail, sa.OrganizedByPhone, sa.EventType, sa.EventTypeOther, sa.EventResult, sa.EventResultOther, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated Events"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteEvents ...
func (saw *studentEvents) DeleteEvents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete Events")
	binding.Validator = &defaultValidator{}
	var sa models.StudentEventsModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteEvents()
	err = models.StudentInfoService.UpdateStudentInfo("STU_Event_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted Events"})
	c.Abort()
	return
}
