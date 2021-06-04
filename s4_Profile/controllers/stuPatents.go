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

// StudentPatents ...
var (
	StudentPatents studentPatents = studentPatents{}
)

type studentPatents struct{}

// AddPatents ...
func (saw *studentPatents) AddPatents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Patents")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPatentsModel
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

		err := models.StudentInfoService.AddToStudentInfo("STU_PATENTS_INS", getInterfaceValues(sa))
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Patents Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllPatents ...
func (saw *studentPatents) GetAllPatents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Patents")

	sa, err := getPatents(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Patents", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.Patents)
	c.Abort()
	return

}

func getPatents(ID string) (models.StudentAllPatentsModel, error) {
	var sa models.StudentAllPatentsModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_PATENTS_GETALL", ID)
	if err != nil {
		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentPatentsModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.PatentType, &newSl.PatentNumber, &newSl.PatentStatus, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate, &newSl.SentforVerification, &newSl.DateSentforVerification, &newSl.Verified, &newSl.DateVerified, &newSl.SentbackforRevalidation, &newSl.DateSentBackForRevalidation, &newSl.ValidatorRemarks, &newSl.VerificationType, &newSl.VerifiedByStakeholderID, &newSl.VerifiedByEmailID)
		if err != nil {
			return sa, err
		}
		sa.Patents = append(sa.Patents, newSl)
	}
	return sa, nil
}

// UpdatePatents ...
func (saw *studentPatents) UpdatePatents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update Patents")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPatentsModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Patents ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdatePatents()
		err := models.StudentInfoService.UpdateStudentInfo("STU_PATENTS_UPD", []interface{}{sa.Name, sa.PatentType, sa.PatentNumber, sa.PatentStatus, sa.Attachment, sa.AttachmentName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated Patents"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeletePatents ...
func (saw *studentPatents) DeletePatents(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete Patents")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPatentsModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeletePatents()
	err = models.StudentInfoService.UpdateStudentInfo("STU_PATENTS_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted Patents"})
	c.Abort()
	return
}
