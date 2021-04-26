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

// StudentCompetitions ...
var (
	StudentCompetitions studentCompetitions = studentCompetitions{}
)

type studentCompetitions struct{}

// AddCompetitions ...
func (saw *studentCompetitions) AddCompetitions(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Competitions")
	binding.Validator = &defaultValidator{}
	var sa models.StudentCompetitionModel
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

		err := models.StudentInfoService.AddToStudentInfo("STU_Competitions_INS", getInterfaceValues(sa))
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Competitions Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllCompetitions ...
func (saw *studentCompetitions) GetAllCompetitions(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Competitions")

	var sa models.StudentAllCompetitionModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_Competitions_GETALL", ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Competitions", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentCompetitionModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.Date, &newSl.Rank, &newSl.Attachment, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		sa.Competitions = append(sa.Competitions, newSl)
	}
	c.JSON(http.StatusOK, sa.Competitions)
	c.Abort()
	return

}

// UpdateCompetitions ...
func (saw *studentCompetitions) UpdateCompetitions(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update Competitions")
	binding.Validator = &defaultValidator{}
	var sa models.StudentCompetitionModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Competitions ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateCompetitions()
		err := models.StudentInfoService.UpdateStudentInfo("STU_Competitions_UPD", []interface{}{sa.Name, sa.Date, sa.Rank, sa.Attachment, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated Competitions"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteCompetitions ...
func (saw *studentCompetitions) DeleteCompetitions(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete Competitions")
	binding.Validator = &defaultValidator{}
	var sa models.StudentCompetitionModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteCompetitions()
	err = models.StudentInfoService.UpdateStudentInfo("STU_Competitions_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted Competitions"})
	c.Abort()
	return
}
