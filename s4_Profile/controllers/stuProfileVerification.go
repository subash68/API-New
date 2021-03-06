package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// SRVNotificationType ...
const (
	SRVNotificationType   string = "RequestVerification"
	SRVNotificationTypeID string = "6"
	SRVRedirectURL        string = "/dashboard/univerity"
	UPVNotificationType   string = "ProcessVerification"
	UPVNotificationTypeID string = "7"
	UPVRedirectURL        string = "/dashboard/Academics"
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
	reqBody := map[string]string{"senderID": sd.StakeholderID, "senderUserRole": userType, "notificationType": SRVNotificationType, "content": "Student Profile Verification Request", "publishFlag": "false", "publishID": "", "receiverID": universityID, "redirectedURL": SRVRedirectURL, "isGeneric": "false", "notificationTypeID": SRVNotificationTypeID}
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
	studentsList, customErr := models.GetAllStudentProfileMetadata(ID, false)
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
	studentsList, customErr := models.GetAllStudentProfileMetadata(ID, true)
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
func (spv *studentProfileVerification) GetUnvStudentSVProfile(c *gin.Context) {
	ctx, _, _, successResp := getFuncReq(c, "Verifying Student Profile")
	studentID := c.Param("studentID")
	if studentID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find  studentID in Params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	sa, err := GetStudentVRFProfile(studentID, true, false)
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

// GetUnvStudentProfile ...
func (spv *studentProfileVerification) GetUnvStudentVerifiedProfile(c *gin.Context) {
	ctx, _, _, successResp := getFuncReq(c, "Verifying Student Profile")
	studentID := c.Param("studentID")
	if studentID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find  studentID in Params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	sa, err := GetStudentVRFProfile(studentID, false, true)
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

// VerificationReqModel ...
type VerificationReqModel struct {
	StudentID          string `form:"studentID" json:"studentID" binding:"required"`
	VerificationStatus bool   `form:"verificationStatus" json:"verificationStatus"`
	Remarks            string `form:"remarks" json:"remarks"`
}

// ProcessRequestVerification ...
func (spv *studentProfileVerification) ProcessRequestVerification(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Process Verification Request")
	var sa models.StuVrfDataModel
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
	if err == nil {

		nftContent := "University updated the verification request"

		err := sa.ValidateStudentProfile(ID)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}

		reqBody := map[string]string{"senderID": ID, "senderUserRole": userType, "receiverID": sa.StudentID, "notificationType": UPVNotificationType, "content": nftContent, "publishFlag": "false", "publishID": "", "ReceiverID": sa.StudentID, "redirectedURL": UPVRedirectURL, "isGeneric": "false", "notificationTypeID": UPVNotificationTypeID}
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

// GetStudentVRFProfile ...
func GetStudentVRFProfile(ID string, sv bool, vrf bool) (models.StudentCompleteProfileDataModel, error) {
	stuData, err := returnCompleteStuData(ID)
	if err != nil {
		return stuData, err
	}
	var vrfStuData models.StudentCompleteProfileDataModel

	vrfStuData.Profile = stuData.Profile
	vrfStuData.ContactInfo = stuData.ContactInfo
	if stuData.Academics.Tenth.SentforVerification == sv && stuData.Academics.Tenth.Verified == vrf {
		vrfStuData.Academics.Tenth = stuData.Academics.Tenth
		vrfStuData.Academics.Twelfth = stuData.Academics.Twelfth
	}
	semAdded := false
	for _, v := range stuData.Academics.Graduation.Semesters {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.Academics.Graduation.Semesters = append(vrfStuData.Academics.Graduation.Semesters, v)
			semAdded = true
		}
	}
	if semAdded {
		vrfStuData.Academics.Graduation = parseGrad(stuData.Academics.Graduation, vrfStuData.Academics.Graduation.Semesters)
	}

	semAdded = false
	for _, v := range stuData.Academics.PostGraduation.Semesters {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.Academics.PostGraduation.Semesters = append(vrfStuData.Academics.PostGraduation.Semesters, v)
			semAdded = true
		}
	}
	if semAdded {
		vrfStuData.Academics.PostGraduation = parsePG(stuData.Academics.PostGraduation, vrfStuData.Academics.PostGraduation.Semesters)
	}

	// certificates ...
	for _, v := range stuData.CertificationsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.CertificationsArray = append(vrfStuData.CertificationsArray, v)
		}
	}

	// certificates ...
	for _, v := range stuData.AssessmentsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.AssessmentsArray = append(vrfStuData.AssessmentsArray, v)
		}
	}

	// InternshipsArray ...
	for _, v := range stuData.InternshipsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.InternshipsArray = append(vrfStuData.InternshipsArray, v)
		}
	}

	// AwardsArray ...
	for _, v := range stuData.AwardsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.AwardsArray = append(vrfStuData.AwardsArray, v)
		}
	}

	// ExtraCurricularArray ...
	for _, v := range stuData.ExtraCurricularArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.ExtraCurricularArray = append(vrfStuData.ExtraCurricularArray, v)
		}
	}

	// PatentsArray ...
	for _, v := range stuData.PatentsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.PatentsArray = append(vrfStuData.PatentsArray, v)
		}
	}

	// ProjectsArray ...
	for _, v := range stuData.ProjectsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.ProjectsArray = append(vrfStuData.ProjectsArray, v)
		}
	}

	// PatentsArray ...
	for _, v := range stuData.PublicationsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.PublicationsArray = append(vrfStuData.PublicationsArray, v)
		}
	}

	// ScholarshipsArray ...
	for _, v := range stuData.ScholarshipsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.ScholarshipsArray = append(vrfStuData.ScholarshipsArray, v)
		}
	}

	// TestScoresArray ...
	for _, v := range stuData.TestScoresArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.TestScoresArray = append(vrfStuData.TestScoresArray, v)
		}
	}

	// VolunteerExperienceArray ...
	for _, v := range stuData.VolunteerExperienceArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.VolunteerExperienceArray = append(vrfStuData.VolunteerExperienceArray, v)
		}
	}

	// EventsArray ...
	for _, v := range stuData.EventsArray {
		if v.SentforVerification == sv && v.Verified == vrf {
			vrfStuData.EventsArray = append(vrfStuData.EventsArray, v)
		}
	}

	return vrfStuData, nil
}

func parseGrad(stuGrad models.StudentGradModel, sems []models.StudentSemesterModel) models.StudentGradModel {
	var vrfGrad models.StudentGradModel
	vrfGrad.UniversityStakeholderIDUniv = stuGrad.UniversityStakeholderIDUniv
	vrfGrad.CollegeRollNumber = stuGrad.CollegeRollNumber
	vrfGrad.ExpectedYearOfPassing = stuGrad.ExpectedYearOfPassing
	vrfGrad.ProgramID = stuGrad.ProgramID
	vrfGrad.ProgramName = stuGrad.ProgramName
	vrfGrad.BranchID = stuGrad.BranchID
	vrfGrad.BranchName = stuGrad.BranchName
	vrfGrad.FinalCGPA = stuGrad.FinalCGPA
	vrfGrad.FinalPercentage = stuGrad.FinalPercentage
	vrfGrad.ActiveBacklogsNumber = stuGrad.ActiveBacklogsNumber
	vrfGrad.TotalNumberOfBacklogs = stuGrad.TotalNumberOfBacklogs
	vrfGrad.Semesters = sems
	return vrfGrad
}

func parsePG(stuGrad models.StudentPGModel, sems []models.StudentSemesterModel) models.StudentPGModel {
	var vrfGrad models.StudentPGModel
	vrfGrad.UniversityStakeholderIDUniv = stuGrad.UniversityStakeholderIDUniv
	vrfGrad.CollegeRollNumber = stuGrad.CollegeRollNumber
	vrfGrad.ExpectedYearOfPassing = stuGrad.ExpectedYearOfPassing
	vrfGrad.ProgramID = stuGrad.ProgramID
	vrfGrad.ProgramName = stuGrad.ProgramName
	vrfGrad.BranchID = stuGrad.BranchID
	vrfGrad.BranchName = stuGrad.BranchName
	vrfGrad.FinalCGPA = stuGrad.FinalCGPA
	vrfGrad.FinalPercentage = stuGrad.FinalPercentage
	vrfGrad.Semesters = sems
	return vrfGrad
}
