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

// StudentTestScores ...
var (
	StudentTestScores studentTestScores = studentTestScores{}
)

type studentTestScores struct{}

// AddTestScores ...
func (saw *studentTestScores) AddTestScores(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add TestScores")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTestScoresModel
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
		vals := []interface{}{sa.StakeholderID, sa.Name, sa.TestScoreDate, sa.TestScore, sa.TestScoreTotal, sa.Attachment, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_TEST_SCORES_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"TestScores Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllTestScores ...
func (saw *studentTestScores) GetAllTestScores(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get TestScores")

	var sa models.StudentAllTestScoresModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_TEST_SCORES_GETALL", ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get TestScores", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentTestScoresModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.TestScoreDate, &newSl.TestScore, &newSl.TestScoreTotal, &newSl.Attachment, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		sa.TestScores = append(sa.TestScores, newSl)
	}
	c.JSON(http.StatusOK, sa.TestScores)
	c.Abort()
	return

}

// UpdateTestScores ...
func (saw *studentTestScores) UpdateTestScores(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update TestScores")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTestScoresModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find TestScores ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateTestScores()
		err := models.StudentInfoService.UpdateStudentInfo("STU_TEST_SCORES_UPD", []interface{}{sa.Name, sa.TestScoreDate, sa.TestScore, sa.TestScoreTotal, sa.Attachment, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated TestScores"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteTestScores ...
func (saw *studentTestScores) DeleteTestScores(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete TestScores")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTestScoresModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteTestScores()
	err = models.StudentInfoService.UpdateStudentInfo("STU_TEST_SCORES_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted TestScores"})
	c.Abort()
	return
}
