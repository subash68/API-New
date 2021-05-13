package controllers

import (
	"context"
	"fmt"
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
		counts := models.GetRegisteredCounts()
		c.JSON(http.StatusOK, models.LoginRespModel{Token: token, RedirectURL: redirectURL, Stats: counts})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1LGN", ErrTyp: "Required Information not found", Err: err})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}
