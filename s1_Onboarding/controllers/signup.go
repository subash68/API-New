// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/configuration"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/middleware"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/services"
	"golang.org/x/crypto/bcrypt"
)

// StakeholderType ...
type StakeholderType struct {
	Stakeholder string `form:"stakeholder"`
	Password    string `form:"password" binding:"required,min=8,max=15"`
}

// SignupResp ...
type SignupResp struct {
	AccountStatus string  `json:"accountStatus"`
	Message       string  `json:"message"`
	PlatformID    string  `json:"platformUID"`
	Email         string  `json:"email"`
	Phone         string  `json:"phoneNumber"`
	Stakeholder   string  `json:"stakeholder"`
	BonusTokens   float64 `json:"bonusTokens"`
}

// ctxFunc declaration for context use
type ctxFunc string

// ctxkey
var ctxkey ctxFunc

type RefCodeResp struct {
	StakeholderID string `json:"stakeholderID"`
}

// CheckRefCode ...
func CheckRefCode(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Check Referral Code")

	defer cancel()
	if c.Param("refCode") == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required referralCode information not found ", Err: fmt.Errorf("Require 'refCode' in params"), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		c.Abort()
		return
	}
	stakeholderID, err := models.CheckRefCode(c.Param(("refCode")))
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Inalid referral code", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, RefCodeResp{stakeholderID})
	c.Abort()
	return
}

// Signup ...
func Signup(c *gin.Context) {
	successResp = map[string]string{}
	var stakeholderInfo StakeholderType
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "SignUp")

	defer cancel()
	defer close(jobdb)
	if c.Param("stakeholder") == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required Stakeholder Type information not found ", Err: fmt.Errorf("Cannot decode Stakeholder type from post Query"), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		c.Abort()
		return
	}
	stakeholderInfo.Stakeholder = c.Param("stakeholder")
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&stakeholderInfo, binding.Form)
	if err == nil {
		fmt.Printf("\n stakeholderInfo --> %+v\n", stakeholderInfo)
		bcryptEncodedPassword, err := bcrypt.GenerateFromPassword([]byte(stakeholderInfo.Password), 10)
		if err != nil {
			fmt.Println("Failed To Bcrypt the Password " + err.Error())
			respErr.Code = "S1AUTHSNP001"
			respErr.Message = "Failed To store the Password " + err.Error()
			c.JSON(http.StatusUnprocessableEntity, respErr)
			c.Abort()
			return
		}
		expiryDate := time.Now().AddDate(0, 0, 1).Format(time.RFC3339)

		switch stakeholderInfo.Stakeholder {
		case "Corporate":
			corporateData := serializeCorporateData(ctx, c, string(bcryptEncodedPassword))
			if corporateData.CorporateName == "" {
				return
			}
			go func() {
				select {
				case insertJobChan := <-corporateData.Insert(expiryDate):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "University":
			universityData := serializeUniversityData(ctx, c, string(bcryptEncodedPassword))
			if universityData.UniversityName == "" {
				return
			}
			go func() {
				select {
				case insertJobChan := <-universityData.Insert(expiryDate):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "Student":
			studentData := serializeStudentData(ctx, c, string(bcryptEncodedPassword))
			if studentData.FirstName == "" {
				fmt.Println(studentData)
				// resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid STudentDetails", Err: fmt.Errorf("Invalid student information"), SuccessResp: successResp})
				// c.JSON(http.StatusUnprocessableEntity, resp)
				// c.Abort()
				// return
				return
			}
			go func() {
				select {
				case insertJobChan := <-studentData.Insert(expiryDate):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid Stakeholder type", Err: fmt.Errorf("" + stakeholderInfo.Stakeholder + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		insertJob := <-jobdb

		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		// token, err := CreateToken(ctx, insertJob.SuccessResp["StakeholderID"], "VRF_TOK", c)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, resp)
		// }
		fmt.Printf("\n insertjob: %v\n", insertJob)

		ms, err := services.SendSmsOtp(insertJob.SuccessResp["Phone"])
		if err != nil {
			fmt.Printf("\nFailed to send mobile otp\n")
			fmt.Println(err)
		}
		es, err := services.SendOTPEmail(insertJob.SuccessResp["Email"], insertJob.SuccessResp["StakeholderID"])
		if err != nil {
			fmt.Printf("\nFailed to send Email otp\n")
			fmt.Println(err)
		}
		fmt.Printf("Phone verification sent %v, email verification sent %v", ms, es)
		// tokenAdded, err := raiseBonusTokenReq(insertJob.SuccessResp["StakeholderID"])
		// if err != nil {
		// 	fmt.Println("Failed to assign Bonus tokens %v", err.Error())
		// }
		var respData SignupResp
		respData.AccountStatus = "1"
		respData.Message = "OTP sent to Mobile and Email for verification"
		respData.PlatformID = insertJob.SuccessResp["StakeholderID"]
		respData.Email = insertJob.SuccessResp["Email"]
		respData.Phone = insertJob.SuccessResp["Phone"]
		respData.Stakeholder = stakeholderInfo.Stakeholder
		// respData.BonusTokens = tokenAdded

		c.JSON(http.StatusOK, respData)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required information not found ", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return
}

func serializeCorporateData(ctx context.Context, c *gin.Context, encodedPass string) models.CorporateMasterDB {
	var corporateData models.CorporateMasterDB
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&corporateData, binding.Form)
	if err == nil {
		corporateData.Password = encodedPass
		corporateData.AccountStatus = "1"
		corporateData.Attachment, corporateData.AttachmentName, corporateData.ProfilePicture = attachFile(c)
		return corporateData
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return models.CorporateMasterDB{}
}

func serializeUniversityData(ctx context.Context, c *gin.Context, encodedPass string) models.UniversityMasterDb {
	var universityData models.UniversityMasterDb
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&universityData, binding.Form)
	if err == nil {
		universityData.Password = encodedPass
		universityData.AccountStatus = "1"
		universityData.Attachment, universityData.AttachmentName, universityData.ProfilePicture = attachFile(c)
		return universityData
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return models.UniversityMasterDb{}
}

func serializeStudentData(ctx context.Context, c *gin.Context, encodedPass string) models.StudentMasterDb {
	var studentData models.StudentMasterDb
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&studentData, binding.Form)
	if err == nil {
		studentData.Password = encodedPass
		studentData.AccountStatus = "1"
		studentData.Attachment, studentData.AttachmentName, studentData.ProfilePicture = attachFile(c)
		return studentData
	}
	fmt.Println(studentData)
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return models.StudentMasterDb{}

}

func attachFile(c *gin.Context) ([]byte, string, []byte) {
	form, _ := c.MultipartForm()
	files := form.File["attachment"]
	var err error
	var attchFile, prfPic []byte
	var attName string
	for _, file := range files {
		fileContent, _ := file.Open()
		attchFile, err = ioutil.ReadAll(fileContent)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			c.Abort()
			return nil, "", nil
		}
		attName = file.Filename

	}
	pfiles := form.File["profilePicture"]
	for _, file := range pfiles {
		fileContent, _ := file.Open()
		prfPic, err = ioutil.ReadAll(fileContent)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			c.Abort()
			return nil, "", nil
		}
	}
	return attchFile, attName, prfPic
}

func raiseBonusTokenReq(ID string) (float64, error) {
	tokenConfig := configuration.TokenConfig()
	reqData := map[string]string{"modeOfTokenissue": "Bonus", "stakeholderID": ID, "paymentID": "BT" + GetRandomID(15), "allocatedTokens": "1000"}
	_, err := middleware.MakeInternalServiceCall(tokenConfig.Host, tokenConfig.Port, "POST", "/t/addAllocation", reqData)
	if err != nil {
		return 0, err
	}
	allocatedTokens, err := strconv.ParseFloat(reqData["allocatedTokens"], 64)
	if err != nil {
		return 0, err
	}
	return allocatedTokens, nil
}
