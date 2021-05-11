package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// AddTenth ...
func AddTenth(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Academics")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTenthAcademicsModelReq
	fmt.Println((c.Request.Body))
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		form, _ := c.MultipartForm()
		var stt models.StudentAcademicsModelReq
		stt.StakeholderID = ID
		stt.Tenth = sa.Tenth
		err := stt.InsertAcademics(form)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Academics Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found in Add academics", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// AddTwelfth ...
func AddTwelfth(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Academics")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTwelfthAcademicsModelReq
	fmt.Println((c.Request.Body))
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		form, _ := c.MultipartForm()
		var stt models.StudentAcademicsModelReq
		stt.StakeholderID = ID
		stt.Twelfth = sa.Twelfth
		err := stt.InsertAcademics(form)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Academics Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found in Add academics", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// AddGraduation ...
func AddGraduation(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Graduation")
	binding.Validator = &defaultValidator{}
	var sa models.StudentGradAcademicsModelReq
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid Headers", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	binding.Validator = &defaultValidator{}

	err = c.ShouldBindWith(&sa, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))

	if len(sa.Graduation.Semesters) <= 0 {
		err = fmt.Errorf("semester details not found")
	}
	if err == nil {
		form, _ := c.MultipartForm()
		var stt models.StudentAcademicsModelReq
		stt.StakeholderID = ID
		stt.Graduation = sa.Graduation
		err := stt.InsertAcademics(form)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Academics Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found in Add academics", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// AddPostGraduation ...
func AddPostGraduation(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Post Graduation")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPGAcademicsModelReq
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid Headers", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	binding.Validator = &defaultValidator{}

	err = c.ShouldBindWith(&sa, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))

	if len(sa.PostGraduation.Semesters) <= 0 {
		err = fmt.Errorf("semester details not found")
	}
	if err == nil {
		form, _ := c.MultipartForm()
		var stt models.StudentAcademicsModelReq
		stt.StakeholderID = ID
		stt.PostGraduation = sa.PostGraduation
		err := stt.InsertAcademics(form)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Academics Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found in Add academics", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAcademics ...
func GetAcademics(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Academics")
	sa, err := models.GetAcademics(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa)
	c.Abort()
	return

}
