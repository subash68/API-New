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

// StudentExtraCurricular ...
var (
	StudentExtraCurricular studentExtraCurricular = studentExtraCurricular{}
)

type studentExtraCurricular struct{}

// AddExtraCurricular ...
func (saw *studentExtraCurricular) AddExtraCurricular(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add ExtraCurricular")
	binding.Validator = &defaultValidator{}
	var sa models.StudentExtraCurricularModel
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

		err := models.StudentInfoService.AddToStudentInfo("STU_EXTRA_CURRICULAR_INS", getInterfaceValues(sa))
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"ExtraCurricular Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllExtraCurricular ...
func (saw *studentExtraCurricular) GetAllExtraCurricular(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get ExtraCurricular")

	sa, err := getAllExtraCurricular(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.ExtraCurricular)
	c.Abort()
	return

}

func getAllExtraCurricular(ID string) (models.StudentAllExtraCurricularModel, error) {
	var sa models.StudentAllExtraCurricularModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_EXTRA_CURRICULAR_GETALL", ID)
	if err != nil {
		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentExtraCurricularModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			return sa, err
		}
		sa.ExtraCurricular = append(sa.ExtraCurricular, newSl)
	}
	return sa, nil
}

// UpdateExtraCurricular ...
func (saw *studentExtraCurricular) UpdateExtraCurricular(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update ExtraCurricular")
	binding.Validator = &defaultValidator{}
	var sa models.StudentExtraCurricularModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find ExtraCurricular ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateExtraCurricular()
		err := models.StudentInfoService.UpdateStudentInfo("STU_EXTRA_CURRICULAR_UPD", []interface{}{sa.Name, sa.Attachment, sa.AttachmentName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated ExtraCurricular"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteExtraCurricular ...
func (saw *studentExtraCurricular) DeleteExtraCurricular(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete ExtraCurricular")
	binding.Validator = &defaultValidator{}
	var sa models.StudentExtraCurricularModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteExtraCurricular()
	err = models.StudentInfoService.UpdateStudentInfo("STU_EXTRA_CURRICULAR_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted ExtraCurricular"})
	c.Abort()
	return
}
