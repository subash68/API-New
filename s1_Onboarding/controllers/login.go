package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
)

// UserCredsModel ...
type UserCredsModel struct {
	StakeholderType string `form:"stakeholder" binding:"required"`
	UserID          string `form:"userID" binding:"required"`
	Password        string `form:"password" binding:"required,min=8,max=15"`
}

// LoginRespModel ...
type LoginRespModel struct {
	Token       string `json:"token" binding:"required"`
	RedirectURL string `json:"redirectURL" binding:"required"`
}

// Offchain ...
type Offchain struct {
	FromMSP     string `json:"fromMSP"`
	ToMSP       string `json:"toMSP"`
	Data        string `json:"data"`
	DataHash    string `json:"dataHash"`
	TimeStamp   string `json:"timeStamp"`
	ReferenceID string `json:"id"`
}

// RefID ...
var RefID string

// APICheck ...
func APICheck(c *gin.Context) {
	RefID = c.Param("refID")
	body, _ := ioutil.ReadAll(c.Request.Body)
	var od Offchain
	err := json.Unmarshal(body, &od)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(od)
	c.JSON(http.StatusOK, od)
	return
}

// APIGetCheck ...
func APIGetCheck(c *gin.Context) {
	RefID = c.Query("refID")
	od := Offchain{"ORG1", "ORG2", "ZGF0YTEyMzQ=", "4508767822a67b7a051a8fe50250897c38c044ab10d9070d740a05a21deb4499", "1617606953358906651", RefID}
	fmt.Println(od)
	c.JSON(http.StatusOK, od)
	return
}

// APIDelCheck ...
func APIDelCheck(c *gin.Context) {
	rf := c.Param("refID")
	fmt.Println(rf, RefID)
	if rf == RefID {
		RefID = "Del"
	}
	fmt.Println("DELETE -------> ", RefID)
	c.JSON(http.StatusOK, "Deleted resp from api")
	return
}

// APIGetAllCheck ...
func APIGetAllCheck(c *gin.Context) {
	str := struct {
		ReferenceIds []string `json:"referenceIds"`
	}{}
	fmt.Println("get all -------> ", RefID)
	if RefID == "" {
		str.ReferenceIds = []string{"ref1", "ref2"}
	} else if RefID == "Del" {
		str.ReferenceIds = []string{}
	} else {
		str.ReferenceIds = []string{RefID}
	}
	fmt.Println(str)
	c.JSON(http.StatusOK, str)
	return
}

// Login ...
func Login(c *gin.Context) {
	var credsModel UserCredsModel
	jobdb := make(chan models.DbModelError, 1)
	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "LoginInformation")
	redirectURL := "/payment"

	defer cancel()
	defer close(jobdb)

	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&credsModel, binding.Form)
	if err == nil {
		switch credsModel.StakeholderType {
		case "Corporate":
			go func() {
				corporateDb := models.CorporateMasterDB{}
				select {
				case insertJobChan := <-corporateDb.Login(credsModel.UserID, credsModel.Password):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "University":
			go func() {
				universityDb := models.UniversityMasterDb{}
				select {
				case insertJobChan := <-universityDb.Login(credsModel.UserID, credsModel.Password):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "Student":
			go func() {
				studentDb := models.StudentMasterDb{}
				select {
				case insertJobChan := <-studentDb.Login(credsModel.UserID, credsModel.Password):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid Stakehodler type", Err: fmt.Errorf("" + credsModel.StakeholderType + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		insertJob := <-jobdb

		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusMethodNotAllowed, resp)
			return
		}
		redirectURL = insertJob.ErrCode
		token, err := CreateToken(ctx, credsModel.StakeholderType, insertJob.SuccessResp["StakeholderID"], "AUTH_TOK")
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1LGN", ErrTyp: "Token Creation ", Err: err})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		c.JSON(http.StatusOK, LoginRespModel{Token: token, RedirectURL: redirectURL})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1LGN", ErrTyp: "Required Information not found", Err: err})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}
