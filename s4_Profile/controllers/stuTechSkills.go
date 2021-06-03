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

// StudentTechSkills ...
var (
	StudentTechSkills studentTechSkills = studentTechSkills{}
)

type studentTechSkills struct{}

// AddTechSkills ...
func (saw *studentTechSkills) AddTechSkills(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Add TechSkills")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTechSkillsModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID
		currentTime := time.Now()
		sa.CreationDate = currentTime
		sa.LastUpdatedDate = currentTime
		sa.EnabledFlag = true
		vals := []interface{}{sa.StakeholderID, sa.SkillID, sa.SkillName, sa.EnabledFlag, sa.CreationDate, sa.LastUpdatedDate}
		err := models.StudentInfoService.AddToStudentInfo("STU_TECH_SKILLS_INS", vals)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"TechSkills Saved"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllTechSkills ...
func (saw *studentTechSkills) GetAllTechSkills(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get TechSkills")

	sa, err := getTechSkills(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get TechSkills", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sa.TechSkills)
	c.Abort()
	return

}

func getTechSkills(ID string) (models.StudentAllTechSkillsModel, error) {
	var sa models.StudentAllTechSkillsModel
	awardRows, err := models.StudentInfoService.GetAllStudentInfo("STU_TECH_SKILLS_GETALL", ID)
	if err != nil {

		return sa, err
	}
	defer awardRows.Close()
	for awardRows.Next() {
		var newSl models.StudentTechSkillsModel
		err = awardRows.Scan(&newSl.ID, &newSl.SkillID, &newSl.SkillName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {

			return sa, err
		}
		sa.TechSkills = append(sa.TechSkills, newSl)
	}
	return sa, nil
}

// UpdateTechSkills ...
func (saw *studentTechSkills) UpdateTechSkills(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Update TechSkills")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTechSkillsModel
	err := c.ShouldBindWith(&sa, binding.Form)
	if err == nil {
		sa.StakeholderID = ID

		sa.ID, err = strconv.Atoi(c.Param("id"))
		if sa.ID <= 0 || err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find TechSkills ID"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		//err := sa.UpdateTechSkills()
		err := models.StudentInfoService.UpdateStudentInfo("STU_TECH_SKILLS_UPD", []interface{}{sa.SkillID, sa.SkillName, time.Now(), sa.ID, sa.StakeholderID})
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.MessageResp{"Updated TechSkills"})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// DeleteTechSkills ...
func (saw *studentTechSkills) DeleteTechSkills(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Delete TechSkills")
	binding.Validator = &defaultValidator{}
	var sa models.StudentTechSkillsModel
	var err error
	sa.StakeholderID = ID
	sa.ID, err = strconv.Atoi(c.Param("id"))
	if sa.ID == 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Language ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	//err = sa.DeleteTechSkills()
	err = models.StudentInfoService.UpdateStudentInfo("STU_TECH_SKILLS_DLT", []interface{}{sa.ID, sa.StakeholderID})
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.MessageResp{"Deleted TechSkills"})
	c.Abort()
	return
}
