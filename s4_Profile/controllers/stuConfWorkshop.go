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

// StudentConfWorkshop ...
var (
	StudentConfWorkshop studentConfWorkshop = studentConfWorkshop{}
)

type studentConfWorkshop struct{}

// AddConfWorkshop ...
func (saw *studentConfWorkshop) AddConfWorkshop(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add ConfWorkshop")
	binding.Validator = &defaultValidator{}
	var sa models.StudentConfWorkshopModel
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
		}
		currentTime := time.Now()
		sa.CreationDate = currentTime
		sa.LastUpdatedDate = currentTime
		sa.EnabledFlag = true

		err := models.StudentInfoService.AddToStudentInfo("STU_CONF_WORKSHOP_INS", getInterfaceValues(sa))
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"ConfWorkshop Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllConfWorkshop ...
func (saw *studentConfWorkshop) GetAllConfWorkshop(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get ConfWorkshop")

	var sa models.StudentAllConfWorkshopModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_CONF_WORKSHOP_GETALL", ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get ConfWorkshop", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentConfWorkshopModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.Date, &newSl.Attachment, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		sa.ConfWorkshop = append(sa.ConfWorkshop, newSl)
	}
	c.JSON(http.StatusOK, sa.ConfWorkshop)
	c.Abort()
	return

}

// UpdateConfWorkshop ...
func (saw *studentConfWorkshop) UpdateConfWorkshop(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update ConfWorkshop")
	binding.Validator = &defaultValidator{}
	var sa models.StudentConfWorkshopModel
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
		}
		sa.ID, err = strconv.Atoi(c.Param("id"))
		if sa.ID <= 0 || err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find ConfWorkshop ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateConfWorkshop()
		err := models.StudentInfoService.UpdateStudentInfo("STU_CONF_WORKSHOP_UPD", []interface{}{sa.Name, sa.Date, sa.Attachment, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated ConfWorkshop"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteConfWorkshop ...
func (saw *studentConfWorkshop) DeleteConfWorkshop(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete ConfWorkshop")
	binding.Validator = &defaultValidator{}
	var sa models.StudentConfWorkshopModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteConfWorkshop()
	err = models.StudentInfoService.UpdateStudentInfo("STU_CONF_WORKSHOP_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted ConfWorkshop"})
	c.Abort()
	return
}
