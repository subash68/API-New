package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
)

// ActivateAccountModel ...
type ActivateAccountModel struct {
	Stakeholder   string `form:"stakeholder" binding:"required"`
	StakeholderID string `form:"stakeholderID" binding:"required"`
}

// ActivateAccount ...
func ActivateAccount(c *gin.Context) {
	jobdb := make(chan models.DbModelError, 1)
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Account Activation")

	defer cancel()

	var accActivate ActivateAccountModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&accActivate, binding.Form)
	if err == nil {
		expiryDate := time.Now().AddDate(1, 0, 0).Format(time.RFC3339)
		switch accActivate.Stakeholder {
		case "Corporate":
			corporateData := models.CorporateMasterDB{StakeholderID: accActivate.StakeholderID}
			go func() {
				select {
				case insertJobChan := <-corporateData.UpdateAccountStatus(expiryDate):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "University":
			universityData := models.UniversityMasterDb{StakeholderID: accActivate.StakeholderID}
			go func() {
				select {
				case insertJobChan := <-universityData.UpdateAccountStatus(expiryDate):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		case "Student":
			studentData := models.StudentMasterDb{StakeholderID: accActivate.StakeholderID}
			go func() {
				select {
				case insertJobChan := <-studentData.UpdateAccountStatus(expiryDate):
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Invalid Stakeholder type", Err: fmt.Errorf("" + accActivate.Stakeholder + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
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

		c.JSON(http.StatusOK, VerfSucResp{"Account activation successful"})
		return
	} else {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1VRF", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
}
