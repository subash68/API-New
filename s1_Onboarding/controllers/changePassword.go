// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/services"
	"golang.org/x/crypto/bcrypt"
)

// ChangePasswordModel ...
type ChangePasswordModel struct {
	Stakeholder  string `form:"stakeholder" binding:"required"`
	Otp          string `form:"otp" binding:"required"`
	VrfBy        string `form:"vrfBy" binding:"required"`
	Phone        string `form:"phone"`
	Email        string `form:"email"`
	PlatformUUID string `form:"platformUUID"`
	NewPassword  string `form:"newPassword" binding:"required,min=8,max=15"`
}

// SendOtpCPModel ...
type SendOtpCPModel struct {
	Stakeholder  string `form:"stakeholder" binding:"required"`
	VrfBy        string `form:"vrfBy" binding:"required"`
	Phone        string `form:"phone"`
	Email        string `form:"email"`
	PlatformUUID string `form:"platformUUID"`
}

// ChangePassRespModel ...
type ChangePassRespModel struct {
	Message      string `json:"message"`
	RedirectPath string `json:"Redirectpath"`
}

// SendChangePasswordOTP ...
func SendChangePasswordOTP(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Change Password")

	defer cancel()

	var stcp SendOtpCPModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&stcp, binding.Form)
	if err == nil {
		var dbError models.DbModelError
		var search string
		switch stcp.VrfBy {
		case "Phone":
			search = stcp.Phone
			stcp.PlatformUUID, dbError = models.GetPlatformUUID(search, stcp.Stakeholder)
			if dbError.ErrTyp != "000" {
				resp := ErrCheck(ctx, dbError)
				c.Error(dbError.Err)
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			otpSent, err := services.SendSmsOtp(stcp.Phone)
			if err != nil || !otpSent {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Failed to send OTP to " + stcp.Phone, Err: err, SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			break
		case "Email":
			search = stcp.Email
			stcp.PlatformUUID, dbError = models.GetPlatformUUID(search, stcp.Stakeholder)
			if dbError.ErrTyp != "000" {
				resp := ErrCheck(ctx, dbError)
				c.Error(dbError.Err)
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			otpSent, err := services.SendOTPEmail(stcp.Email, stcp.PlatformUUID)
			if err != nil || !otpSent {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Failed to send OTP to " + stcp.Email, Err: err, SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid vrfBy type", Err: fmt.Errorf("" + stcp.VrfBy + " is invaild,  Expecting Phone or Email"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		c.JSON(http.StatusOK, "OTP Sent successfully")
		return

		//verified, err := services.ValidateOTP(cp.Otp, cp.Phone)
		//processOtpValidation(ctx, c, err, successResp, mobileOtpData.Stakeholder, mobileOtpData.PlatformUID, verified, false, verified)
	} else {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
}

// ChangePassword ...
func ChangePassword(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Change Password")

	defer cancel()

	var cp ChangePasswordModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&cp, binding.Form)
	if err == nil {
		var dbError models.DbModelError
		var search string
		switch cp.VrfBy {
		case "Phone":
			search = cp.Phone
			verified, err := services.ValidateOTP(cp.Otp, cp.Phone)
			if err != nil || !verified {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "OTP validation fialed ", Err: err, SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			cp.PlatformUUID, dbError = models.GetPlatformUUID(search, cp.Stakeholder)
			if dbError.ErrTyp != "000" {
				resp := ErrCheck(ctx, dbError)
				c.Error(dbError.Err)
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			break
		case "Email":
			search = cp.Email
			cp.PlatformUUID, dbError = models.GetPlatformUUID(search, cp.Stakeholder)
			if dbError.ErrTyp != "000" {
				resp := ErrCheck(ctx, dbError)
				c.Error(dbError.Err)
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			verified, err := services.VerifyEmailOtp(cp.PlatformUUID, cp.Otp)
			if err != nil || !verified {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "OTP validation fialed ", Err: err, SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid vrfBy type", Err: fmt.Errorf("" + cp.VrfBy + " is invaild,  Expecting Phone or Email"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		bcryptEncodedPassword, err := bcrypt.GenerateFromPassword([]byte(cp.NewPassword), 10)
		if err != nil {
			fmt.Println("Failed To Bcrypt the Password " + err.Error())
			respErr.Code = "S1AUTHSNP001"
			respErr.Message = "Failed To store the Password " + err.Error()
			c.JSON(http.StatusUnprocessableEntity, respErr)
			return
		}
		dbError = models.ChangePassword(string(bcryptEncodedPassword), cp.PlatformUUID, cp.Stakeholder)
		if dbError.ErrTyp != "000" {
			resp := ErrCheck(ctx, dbError)
			c.Error(dbError.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		c.JSON(http.StatusOK, ChangePassRespModel{"Password Changed Succesfully", "/login"})
		return

		//verified, err := services.ValidateOTP(cp.Otp, cp.Phone)
		//processOtpValidation(ctx, c, err, successResp, mobileOtpData.Stakeholder, mobileOtpData.PlatformUID, verified, false, verified)
	} else {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
}
