package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// StudentProfileVerification ...
var (
	StudentProfileVerification studentProfileVerification = studentProfileVerification{}
)

type studentProfileVerification struct{}

func (spv *studentProfileVerification) RequestVerification(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Request for verification")
	var sd models.StudentMasterDb
	sd.StakeholderID = ID
	universityID, err := sd.RequestProfileVerification()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	reqBody := map[string]string{"senderID": sd.StakeholderID, "senderUserRole": userType, "notificationType": "Direct", "content": "Student Profile Verification Request", "publishFlag": "false", "publishID": "", "ReceiverID": universityID}
	resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
	if err != nil {
		fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, models.MessageResp{"Request sent to University"})
	c.Abort()
	return
	return
}

// GetRequestedStudentInformation ...
func (spv *studentProfileVerification) GetAllStudentProfileValidationRequests(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Validating Student Profiles")
	studentsList, customErr := models.GetAllStudentProfileMetadata(ID, models.PvcPending)
	if customErr.ErrTyp != "000" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: customErr.Err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, studentsList)
	c.Abort()
	return
}

// GetAllVerifiedStudentProfile ...
func (spv *studentProfileVerification) GetAllVerifiedStudentProfile(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Students Profile")
	studentsList, customErr := models.GetAllStudentProfileMetadata(ID, models.PvcVerified)
	if customErr.ErrTyp != "000" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: customErr.Err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, studentsList)
	c.Abort()
	return
}

// GetUnvStudentProfile ...
func (spv *studentProfileVerification) GetUnvStudentProfile(c *gin.Context) {
	ctx, _, _, successResp := getFuncReq(c, "Student Profile")
	studentID := c.Param("studentID")
	if studentID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find  studentID in Params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	var sa models.StudentProfileVerificationDataModel
	customErr := sa.GetVrfProfileData(studentID)
	if customErr.ErrTyp != "000" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: customErr.Err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa)
	c.Abort()
	return
}

// VerificationReqModel ...
type VerificationReqModel struct {
	StudentID          string `form:"studentID" json:"studentID" binding:"required"`
	VerificationStatus bool   `form:"verificationStatus" json:"verificationStatus"`
	Remarks            string `form:"remarks" json:"remarks"`
}

// ProcessRequestVerification ...
func (spv *studentProfileVerification) ProcessRequestVerification(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Process Verification Request")
	var sa VerificationReqModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		nftContent := "Profile status updated to verified"
		if sa.VerificationStatus == false {
			if sa.Remarks == "" {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: fmt.Errorf("remarks are required"), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			nftContent = "University reviewed the request, Update below details and resubmit for verification \n" + sa.Remarks
		}

		var sd models.StudentMasterDb
		sd.StakeholderID = sa.StudentID
		err := sd.ValidateStudentProfile(sa.VerificationStatus, ID)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}

		reqBody := map[string]string{"senderID": ID, "senderUserRole": userType, "notificationType": "Direct", "content": nftContent, "publishFlag": "false", "publishID": "", "ReceiverID": sd.StakeholderID}
		resp, err := makeTokenServiceCall("/nft/addNotification", reqBody)
		if err != nil {
			fmt.Printf("\n==========Err Resp from Notification =======> %v", err)
		}
		fmt.Println(resp)
		c.JSON(http.StatusOK, models.MessageResp{"Status sent to Student"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}
