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

// StudentProjects ...
var (
	StudentProjects studentProjects = studentProjects{}
)

type studentProjects struct{}

// AddProjects ...
func (saw *studentProjects) AddProjects(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Projects")
	binding.Validator = &defaultValidator{}
	var sa models.StudentProjectsModel
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
		vals := []interface{}{sa.StakeholderID, sa.Name, sa.ProjectAbstract, sa.GuideName, sa.GuideEmail, sa.StartDate, sa.EndDate, sa.Attachment, sa.AttachmentName, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_PROJECTS_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Projects Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllProjects ...
func (saw *studentProjects) GetAllProjects(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Projects")

	sa, err := getProjects(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Projects", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.Projects)
	c.Abort()
	return

}

func getProjects(ID string) (models.StudentAllProjectsModel, error) {
	var sa models.StudentAllProjectsModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_PROJECTS_GETALL", ID)
	if err != nil {
		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentProjectsModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.ProjectAbstract, &newSl.GuideName, &newSl.GuideEmail, &newSl.StartDate, &newSl.EndDate, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			return sa, err
		}
		sa.Projects = append(sa.Projects, newSl)
	}
	return sa, nil
}

// UpdateProjects ...
func (saw *studentProjects) UpdateProjects(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update Projects")
	binding.Validator = &defaultValidator{}
	var sa models.StudentProjectsModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Projects ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateProjects()
		err := models.StudentInfoService.UpdateStudentInfo("STU_PROJECTS_UPD", []interface{}{sa.Name, sa.ProjectAbstract, sa.GuideName, sa.GuideEmail, sa.StartDate, sa.EndDate, sa.Attachment, sa.AttachmentName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated Projects"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteProjects ...
func (saw *studentProjects) DeleteProjects(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete Projects")
	binding.Validator = &defaultValidator{}
	var sa models.StudentProjectsModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteProjects()
	err = models.StudentInfoService.UpdateStudentInfo("STU_PROJECTS_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted Projects"})
	c.Abort()
	return
}
