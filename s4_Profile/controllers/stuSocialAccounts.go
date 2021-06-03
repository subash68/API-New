package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// StudentSocialAccount ...
var (
	StudentSocialAccount studentSocialAccount = studentSocialAccount{}
)

type studentSocialAccount struct{}

// AddSocialAccount ...
func (saw *studentSocialAccount) AddSocialAccount(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add SocialAccount")
	binding.Validator = &defaultValidator{}
	var sa models.StudentSocialAccountModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID
		currentTime := time.Now()
		sa.CreationDate = currentTime
		sa.LastUpdatedDate = currentTime
		sa.EnabledFlag = true
		vals := []interface{}{sa.StakeholderID, sa.UserID, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_SOCIAL_ACCOUNTS_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"SocialAccount Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllSocialAccount ...
func (saw *studentSocialAccount) GetAllSocialAccount(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get SocialAccount")
	sa, err := getSocialAccounts(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get SocialAccount", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.SocialAccounts)
	c.Abort()
	return

}
func getSocialAccounts(ID string) (models.StudentAllSocialAccountModel, error) {

	var sa models.StudentAllSocialAccountModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_SOCIAL_ACCOUNTS_GETALL", ID)
	if err != nil {

		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentSocialAccountModel
		err = awardRows.Scan(&newSl.ID, &newSl.UserID, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {

			return sa, err
		}
		sa.SocialAccounts = append(sa.SocialAccounts, newSl)
	}
	return sa, nil
}

// UpdateSocialAccount ...
func (saw *studentSocialAccount) UpdateSocialAccount(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update SocialAccount")
	binding.Validator = &defaultValidator{}
	var sa models.StudentSocialAccountModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID

		sa.ID, err = strconv.Atoi(c.Param("id"))
		if sa.ID <= 0 || err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find SocialAccount ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateSocialAccount()
		err := models.StudentInfoService.UpdateStudentInfo("STU_SOCIAL_ACCOUNTS_UPD", []interface{}{sa.UserID, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated SocialAccount"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteSocialAccount ...
func (saw *studentSocialAccount) DeleteSocialAccount(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete SocialAccount")
	binding.Validator = &defaultValidator{}
	var sa models.StudentSocialAccountModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteSocialAccount()
	err = models.StudentInfoService.UpdateStudentInfo("STU_SOCIAL_ACCOUNTS_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted SocialAccount"})
	c.Abort()
	return
}
