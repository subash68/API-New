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

// StudentPublications ...
var (
	StudentPublications studentPublications = studentPublications{}
)

type studentPublications struct{}

// AddPublications ...
func (saw *studentPublications) AddPublications(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add Publications")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPublicationsModel
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
		vals := []interface{}{sa.StakeholderID, sa.Name, sa.PublishingAuthority, sa.GuideName, sa.GuideEmail, sa.StartDate, sa.EndDate, sa.Attachment, sa.AttachmentName, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_PUBLICATIONS_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Publications Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllPublications ...
func (saw *studentPublications) GetAllPublications(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Publications")

	sa, err := getPublications(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.Publications)
	c.Abort()
	return

}

func getPublications(ID string) (models.StudentAllPublicationsModel, error) {

	var sa models.StudentAllPublicationsModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_PUBLICATIONS_GETALL", ID)
	if err != nil {
		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentPublicationsModel
		err = awardRows.Scan(&newSl.ID, &newSl.Name, &newSl.PublishingAuthority, &newSl.GuideName, &newSl.GuideEmail, &newSl.StartDate, &newSl.EndDate, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate, &newSl.SentforVerification, &newSl.DateSentforVerification, &newSl.Verified, &newSl.DateVerified, &newSl.SentbackforRevalidation, &newSl.DateSentBackForRevalidation, &newSl.ValidatorRemarks, &newSl.VerificationType, &newSl.VerifiedByStakeholderID, &newSl.VerifiedByEmailID)
		if err != nil {
			return sa, err
		}
		sa.Publications = append(sa.Publications, newSl)
	}
	return sa, nil
}

// UpdatePublications ...
func (saw *studentPublications) UpdatePublications(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update Publications")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPublicationsModel
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
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Publications ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdatePublications()
		err := models.StudentInfoService.UpdateStudentInfo("STU_PUBLICATIONS_UPD", []interface{}{sa.Name, sa.PublishingAuthority, sa.GuideName, sa.GuideEmail, sa.StartDate, sa.EndDate, sa.Attachment, sa.AttachmentName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated Publications"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeletePublications ...
func (saw *studentPublications) DeletePublications(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete Publications")
	binding.Validator = &defaultValidator{}
	var sa models.StudentPublicationsModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeletePublications()
	err = models.StudentInfoService.UpdateStudentInfo("STU_PUBLICATIONS_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted Publications"})
	c.Abort()
	return
}
