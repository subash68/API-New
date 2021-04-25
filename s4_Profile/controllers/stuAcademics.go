package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// AddAcademics ...
func AddAcademics(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Academics")
	binding.Validator = &defaultValidator{}
	var sa models.StudentAcademicsModelReq
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID
		form, _ := c.MultipartForm()
		err := sa.InsertAcademics(form)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Academics Saved"})
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
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
