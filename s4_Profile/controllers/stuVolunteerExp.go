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

// StudentVolunteerExperience ...
var (
	StudentVolunteerExperience studentVolunteerExperience = studentVolunteerExperience{}
)

type studentVolunteerExperience struct{}

// AddVolunteerExperience ...
func (saw *studentVolunteerExperience) AddVolunteerExperience(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add VolunteerExperience")
	binding.Validator = &defaultValidator{}
	var sa models.StudentVolunteerExperienceModel
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
		vals := []interface{}{sa.StakeholderID, sa.Name, sa.Organisation, sa.Location, sa.StartDate, sa.EndDate, sa.Attachment, sa.AttachmentName, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_VOLUNTEER_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"VolunteerExperience Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllVolunteerExperience ...
func (saw *studentVolunteerExperience) GetAllVolunteerExperience(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get VolunteerExperience")

	sa, err := getAllVolunteerExperience(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.VolunteerExperience)
	c.Abort()
	return

}

func getAllVolunteerExperience(ID string) (models.StudentAllVolunteerExperienceModel, error) {

	var sa models.StudentAllVolunteerExperienceModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_VOLUNTEER_GETALL", ID)
	if err != nil {
		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentVolunteerExperienceModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.Organisation, &newSl.Location, &newSl.StartDate, &newSl.EndDate, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			return sa, err
		}
		sa.VolunteerExperience = append(sa.VolunteerExperience, newSl)
	}
	return sa, nil
}

// UpdateVolunteerExperience ...
func (saw *studentVolunteerExperience) UpdateVolunteerExperience(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update VolunteerExperience")
	binding.Validator = &defaultValidator{}
	var sa models.StudentVolunteerExperienceModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find VolunteerExperience ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateVolunteerExperience()
		err := models.StudentInfoService.UpdateStudentInfo("STU_VOLUNTEER_UPD", []interface{}{sa.Name, sa.Organisation, sa.Location, sa.StartDate, sa.EndDate, sa.Attachment, sa.AttachmentName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated VolunteerExperience"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteVolunteerExperience ...
func (saw *studentVolunteerExperience) DeleteVolunteerExperience(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete VolunteerExperience")
	binding.Validator = &defaultValidator{}
	var sa models.StudentVolunteerExperienceModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteVolunteerExperience()
	err = models.StudentInfoService.UpdateStudentInfo("STU_VOLUNTEER_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted VolunteerExperience"})
	c.Abort()
	return
}
