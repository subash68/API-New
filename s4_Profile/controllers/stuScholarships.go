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

// StudentScholarships ...
var (
	StudentScholarships studentScholarships = studentScholarships{}
)

type studentScholarships struct{}

// AddScholarships ...
func (saw *studentScholarships) AddScholarships(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Scholarships")
	binding.Validator = &defaultValidator{}
	var sa models.StudentScholarshipsModel
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
		vals := []interface{}{sa.StakeholderID, sa.Name, sa.ScholarshipIssuedBy, sa.ScholarshipDate, sa.Attachment, sa.AttachmentName, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_SCHOLARSHIPS_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Scholarships Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllScholarships ...
func (saw *studentScholarships) GetAllScholarships(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Scholarships")
	sa, err := getScholarships(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Scholarships", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, sa.Scholarships)
	c.Abort()
	return

}

func getScholarships(ID string) (models.StudentAllScholarshipsModel, error) {
	var sa models.StudentAllScholarshipsModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_SCHOLARSHIPS_GETALL", ID)
	if err != nil {

		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentScholarshipsModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.ScholarshipIssuedBy, &newSl.ScholarshipDate, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {

			return sa, err
		}
		sa.Scholarships = append(sa.Scholarships, newSl)
	}
	return sa, nil
}

// UpdateScholarships ...
func (saw *studentScholarships) UpdateScholarships(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update Scholarships")
	binding.Validator = &defaultValidator{}
	var sa models.StudentScholarshipsModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Scholarships ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateScholarships()
		err := models.StudentInfoService.UpdateStudentInfo("STU_SCHOLARSHIPS_UPD", []interface{}{sa.Name, sa.ScholarshipIssuedBy, sa.ScholarshipDate, sa.Attachment, sa.AttachmentName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated Scholarships"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteScholarships ...
func (saw *studentScholarships) DeleteScholarships(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete Scholarships")
	binding.Validator = &defaultValidator{}
	var sa models.StudentScholarshipsModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteScholarships()
	err = models.StudentInfoService.UpdateStudentInfo("STU_SCHOLARSHIPS_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted Scholarships"})
	c.Abort()
	return
}
