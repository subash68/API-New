package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/services"
)

// MobileVerfData ...
type MobileVerfData struct {
	Stakeholder string `form:"stakeholder" binding:"required"`
	PlatformUID string `form:"platformUID" binding:"required"`
	OTP         string `form:"otp" binding:"required"`
	Phone       string `form:"phone" binding:"required"`
}

// EmailVerfData ...
type EmailVerfData struct {
	Stakeholder string `form:"stakeholder" binding:"required"`
	PlatformUID string `form:"platformUID" binding:"required"`
	Email       string `form:"email" binding:"required"`
	OTP         string `form:"otp" binding:"required"`
}

// CommonOtpVerifyModel ...
type CommonOtpVerifyModel struct {
	Stakeholder string `form:"stakeholder" binding:"required"`
	PlatformUID string `form:"platformUID" binding:"required"`
	Email       string `form:"email" binding:"required"`
	Phone       string `form:"phone" binding:"required"`
	EmailOTP    string `form:"emailOtp"`
	PhoneOTP    string `form:"phoneOtp"`
}

// VerfSucResp ...
type VerfSucResp struct {
	Message        string `json:"message"`
	MobileVerified bool   `json:"mobileVerified`
	EmailVerified  bool   `json:"emailVerified"`
}

// SucResp ...
type SucResp struct {
	Message string `json:"message"`
}

// CommonOTPVerifier ...
func CommonOTPVerifier(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "OTP Verification")

	defer cancel()

	var commonOtpData CommonOtpVerifyModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&commonOtpData, binding.Form)
	if err == nil {
		phoneVerified := false
		emailVerified := false
		if commonOtpData.Phone != "" && commonOtpData.PhoneOTP != "" {
			phoneVerified, err = services.ValidateOTP(commonOtpData.PhoneOTP, commonOtpData.Phone)
		}
		if commonOtpData.Email != "" && commonOtpData.EmailOTP != "" {
			emailVerified, err = services.VerifyEmailOtp(commonOtpData.PlatformUID, commonOtpData.EmailOTP)
		}
		processOtpValidation(ctx, c, err, successResp, commonOtpData.Stakeholder, commonOtpData.PlatformUID, phoneVerified, emailVerified, true)
	} else {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
}

// VerifyMobile ...
func VerifyMobile(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Mobile Verification")

	defer cancel()

	var mobileOtpData MobileVerfData
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&mobileOtpData, binding.Form)
	if err == nil {
		verified, err := services.ValidateOTP(mobileOtpData.OTP, mobileOtpData.Phone)
		processOtpValidation(ctx, c, err, successResp, mobileOtpData.Stakeholder, mobileOtpData.PlatformUID, verified, false, verified)
	} else {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
}

// VerifyEmail ...
func VerifyEmail(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Email Verification")

	defer cancel()

	var emailOtpData EmailVerfData
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&emailOtpData, binding.Form)
	if err == nil {
		verified, err := services.VerifyEmailOtp(emailOtpData.PlatformUID, emailOtpData.OTP)
		processOtpValidation(ctx, c, err, successResp, emailOtpData.Stakeholder, emailOtpData.PlatformUID, false, verified, verified)
	} else {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
}

func processOtpValidation(ctx context.Context, c *gin.Context, err error, successResp map[string]string, stakeholder string, pid string, phoneVrf bool, emailVrf bool, verified bool) {
	jobdb := make(chan models.DbModelError, 1)

	if verified {
		switch stakeholder {
		case "Corporate":
			corporateData := models.CorporateMasterDB{StakeholderID: pid, PrimaryPhoneVerified: phoneVrf, PrimaryEmailVerified: emailVrf}
			go func() {
				select {
				case insertJobChan := <-corporateData.UpdateVrfStatus():
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "University":
			universityData := models.UniversityMasterDb{StakeholderID: pid, PrimaryPhoneVerified: phoneVrf, PrimaryEmailVerified: emailVrf}
			go func() {
				select {
				case insertJobChan := <-universityData.UpdateVrfStatus():
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "Student":
			studentData := models.StudentMasterDb{StakeholderID: pid, PrimaryPhoneVerified: phoneVrf, PrimaryEmailVerified: emailVrf}
			go func() {
				select {
				case insertJobChan := <-studentData.UpdateVrfStatus():
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Invalid Stakehodler type", Err: fmt.Errorf("" + stakeholder + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		insertJob := <-jobdb

		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		// token, err := CreateToken(ctx, insertJob.SuccessResp["StakeholderID"], "VRF_TOK", c)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, resp)
		// }
		fmt.Printf("\n insertjob: %v\n", insertJob)

		c.JSON(http.StatusOK, VerfSucResp{"OTP verification successful", (insertJob.SuccessResp["mobileVerfied"] == "true"), (insertJob.SuccessResp["emailVerified"] == "true")})
		return

	}
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "OTP validation failed ", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Invalid OTP", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// ResendOtpModel ...
type ResendOtpModel struct {
	Stakeholder string `form:"stakeholder" binding:"required"`
	PlatformUID string `form:"platformUID" binding:"required"`
	OtpType     string `form:"otpType" binding:"required"`
}

// ResendOTP ...
func ResendOTP(c *gin.Context) {
	successResp = map[string]string{}
	jobdb := make(chan models.DbModelError, 1)
	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Resend OTP")

	defer cancel()

	var resendOtpData ResendOtpModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&resendOtpData, binding.Form)
	if err == nil {
		if resendOtpData.OtpType != "Email" && resendOtpData.OtpType != "Phone" {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Failed to send OTP", Err: fmt.Errorf("Invalid otpType, Expecting Email or Phone"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		switch resendOtpData.Stakeholder {
		case "Corporate":
			corporateDb := models.CorporateMasterDB{}
			go func() {
				select {
				case insertJobChan := <-corporateDb.GetContactInfo(resendOtpData.PlatformUID):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "University":
			universityDb := models.UniversityMasterDb{}
			go func() {
				select {
				case insertJobChan := <-universityDb.GetContactInfo(resendOtpData.PlatformUID):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "Student":
			studentDb := models.StudentMasterDb{}
			go func() {
				select {
				case insertJobChan := <-studentDb.GetContactInfo(resendOtpData.PlatformUID):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Invalid Stakehodler type", Err: fmt.Errorf("" + resendOtpData.Stakeholder + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		insertJob := <-jobdb

		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		otpSent := false
		sentTo := ""
		if resendOtpData.OtpType == "Email" {
			otpSent, err = services.SendOTPEmail(insertJob.SuccessResp["Email"], resendOtpData.PlatformUID)
			sentTo = insertJob.SuccessResp["Email"]
		} else {
			otpSent, err = services.SendSmsOtp(insertJob.SuccessResp["Phone"])
			sentTo = insertJob.SuccessResp["Phone"]
		}
		fmt.Printf("====================== OTP send %v , %v======", otpSent, err)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Failed to send OTP", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		} else if otpSent {
			c.JSON(http.StatusOK, SucResp{"OTP sent to " + sentTo})
			return
		}
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Failed to send OTP", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Invalid details", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}
