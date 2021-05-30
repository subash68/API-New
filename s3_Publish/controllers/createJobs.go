// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// AddJCResp ...
type AddJCResp struct {
	JcID string `json:"jcID"`
}

// AddJobsCreation ...
func AddJobsCreation(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.FullJobDb
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Creating Job Creation")

	defer cancel()
	defer close(jobdb)
	// if c.Param("jobID") != "" {
	// 	jc.JobID = c.Param("jobID")
	// }
	fmt.Printf("\n +++++++---->  %+v\n", jc)
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		err = fmt.Errorf("Invalid content type %s , Required %s", reqContentType, "application/json")
	} else {
		binding.Validator = &defaultValidator{}
		//err := c.ShouldBindWith(&jc, binding.Form)
		err = c.ShouldBindWith(&jc, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))
		if err == nil {
			// form, _ := c.MultipartForm()
			// for index := range jc.Jobs {
			// 	files := form.File["attachment"+strconv.Itoa(index)]
			// 	if len(files) > 1 {
			// 		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Upload multiple files is not supported for single skill mapping "), SuccessResp: successResp})
			// 		c.JSON(http.StatusUnprocessableEntity, resp)
			// 		return
			// 	}
			// 	for index, file := range files {
			// 		fileContent, _ := file.Open()
			// 		byteContainer, err := ioutil.ReadAll(fileContent)
			// 		if err != nil {
			// 			c.JSON(http.StatusUnprocessableEntity, err.Error())
			// 			return
			// 		}
			// 		jc.Jobs[index].Attachment = byteContainer
			// 	}

			// }

			//c.MultipartForm()

			// jc.Skills = strings.Split(jc.Skills[0][1:len(jc.Skills[0])-1], ",")
			// for i, skill := range jc.Skills {
			// 	jc.Skills[i] = skill[1 : len(skill)-1]
			// }
			// c.JSON(http.StatusOK, jc)
			// return
			ID, ok := c.Get("userID")
			fmt.Println("-----> Got ID", ID.(string))
			if !ok {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			jc.StakeholderID = ID.(string)

			if (jc.HcID != "" || jc.HcName != "") && (jc.HcID == "" || jc.HcName == "") {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Required HiringCriteriaID along with hiringCreteriaName"), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				c.Abort()
				return
			}
			if (jc.Attachment != nil || jc.AttachmentName != "") && (jc.Attachment == nil || jc.AttachmentName == "") {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Required attachment along with attachmentName"), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				c.Abort()
				return
			}
			go func() {
				select {
				case insertJobChan := <-jc.Insert():
					jobdb <- insertJobChan
				case <-ctx.Done():
					return
				}
			}()
			insertJob := <-jobdb
			fmt.Printf("\n insertjob: %+v\n", insertJob)
			if insertJob.ErrTyp != "000" {
				resp := ErrCheck(ctx, insertJob)
				c.Error(insertJob.Err)
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			//
			c.JSON(http.StatusOK, AddJCResp{insertJob.SuccessResp["jcID"]})
			return
		}
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetJobsCreationByID ...
func GetJobsCreationByID(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.JobHcMappingDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get Jobs Created")

	defer cancel()
	defer close(jobdb)

	jc.JobID = c.Param("jobID")
	if jc.JobID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Job ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	go func() {
		select {
		case insertJobChan := <-jc.GetByID():
			jobdb <- insertJobChan
		case <-ctx.Done():
			return
		}
	}()
	insertJob := <-jobdb

	if insertJob.ErrTyp != "000" {
		resp := ErrCheck(ctx, insertJob)
		c.Error(insertJob.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	fmt.Printf("\n HC : %+v\n", jc)

	c.JSON(http.StatusOK, jc)
	return
}

// GetAllJobsCreated ...
func GetAllJobsCreated(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.JobHcMappingDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Jobs Created")

	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	jc.StakeholderID = ID.(string)
	jcArray, err := jc.GetAllJC()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Cannot get Jobs created", Err: fmt.Errorf("Cannot find Jobs created : %v", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	fmt.Printf("\n JC : %+v\n", jcArray)

	c.JSON(http.StatusOK, jcArray)
	return
}

// UpdateJobsCreation ...
func UpdateJobsCreation(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.FullJobDb
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Update Job Created")
	//var customError models.DbModelError
	defer cancel()
	defer close(jobdb)
	var err error
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		err = fmt.Errorf("Invalid content type %s , Required %s", reqContentType, "application/json")
	} else {
		binding.Validator = &defaultValidator{}
		//err := c.ShouldBindWith(&jc, binding.Form)
		err = c.ShouldBindWith(&jc, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))
		if err == nil {

			jc.JobID = c.Param("jobID")
			if jc.JobID == "" {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Job ID in query"), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			ID, ok := c.Get("userID")
			fmt.Println("-----> Got ID", ID.(string), c.PostForm("id"))
			if !ok {
				resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
				c.JSON(http.StatusUnprocessableEntity, resp)
				return
			}
			jc.StakeholderID = ID.(string)
			customError := jc.Update()
			if customError.ErrTyp != "000" {
				resp := ErrCheck(ctx, customError)
				c.Error(customError.Err)
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			c.JSON(http.StatusOK, DelHCResp{"Successfully updated"})
			return
		}
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}

// MapJobToHC ...
func MapJobToHC(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.JobHcMappingDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Jobs Created")
	//var customError models.DbModelError
	defer cancel()
	defer close(jobdb)
	jc.JobID = c.Param("jobID")
	if jc.JobID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Job ID in query"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string), c.PostForm("id"))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	jc.StakeholderID = ID.(string)
	hcid := c.Request.PostFormValue("hiringCriteriaID")
	hcName := c.Request.PostFormValue("hiringCriteriaName")
	if hcid == "" || hcName == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Hiring Criteria ID in form data"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	//jc.HiringCriteriaID = models.NullString{hcid}
	go func() {
		select {
		case insertJobChan := <-jc.MapHC(hcid, hcName):
			jobdb <- insertJobChan
		case <-ctx.Done():
			return
		}
	}()
	insertJob := <-jobdb
	fmt.Printf("\n insertjob: %+v\n", insertJob)
	if insertJob.ErrTyp != "000" {
		resp := ErrCheck(ctx, insertJob)
		c.Error(insertJob.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, DelHCResp{"Successfully updated"})
	return
}

// DeleteJobsCreation ...
func DeleteJobsCreation(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.JobHcMappingDB
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Delete Created Job")

	defer cancel()
	defer close(jobdb)

	jc.JobID = c.Param("jobID")
	if jc.JobID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Job created ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string), c.PostForm("id"))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	jc.StakeholderID = ID.(string)
	go func() {
		select {
		case insertJobChan := <-jc.DeleteByID():
			jobdb <- insertJobChan
		case <-ctx.Done():
			return
		}
	}()
	insertJob := <-jobdb

	if insertJob.ErrTyp != "000" {
		resp := ErrCheck(ctx, insertJob)
		c.Error(insertJob.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	fmt.Printf("\n JC : %+v\n", jc)

	c.JSON(http.StatusOK, DelHCResp{"Job Created With ID :" + jc.JobID + " Has been delete"})
	return
}

// AddSkills ...
func AddSkills(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.SkillsUpdateJobDb
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Creating Job Creation")

	defer cancel()
	defer close(jobdb)
	var err error
	if c.Param("jobID") == "" {
		err = fmt.Errorf("JobID not found in params")
	} else {
		jc.JobID = c.Param("jobID")

		err = c.ShouldBindWith(&jc, binding.Form)
		if len(jc.Jobs) <= 0 {
			err = fmt.Errorf("Require Skills in array format")
		}
	}
	if err == nil {

		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		jc.StakeholderID = ID.(string)
		go func() {
			select {
			case insertJobChan := <-jc.AddSkillsToJC():
				jobdb <- insertJobChan
			case <-ctx.Done():
				return
			}
		}()
		insertJob := <-jobdb
		fmt.Printf("\n insertjob: %+v\n", insertJob)
		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, DelHCResp{"Skills Added"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}

// UpdateJobSkill ...
func UpdateJobSkill(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.FullJobDb
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Update Job Created")
	//var customError models.DbModelError
	defer cancel()
	defer close(jobdb)
	jc.JobID = c.Param("jobID")
	if jc.JobID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Job ID in query"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	SkillID := c.Param("id")
	if SkillID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find id of updating record in query"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string), c.PostForm("id"))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	jc.StakeholderID = ID.(string)
	updateReq := c.Request.PostForm
	form, _ := c.MultipartForm()
	files := form.File["attachment0"]
	var updateFile []byte
	if len(files) > 1 {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Upload multiple files is not supported for single skill mapping "), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	for _, file := range files {
		fileContent, _ := file.Open()
		byteContainer, err := ioutil.ReadAll(fileContent)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		updateFile = byteContainer
	}

	customError := models.UpdatePublishedData(updateReq, "JOB_SKill_MAP_UPD", "JOB_SKill_MAP_WHR", jc.StakeholderID, SkillID, updateFile, "AttachFile")
	if customError.ErrTyp != "000" {
		resp := ErrCheck(ctx, customError)
		c.Error(customError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, DelHCResp{"Successfully updated"})
	return
}

// DeleteJobSkill ...
func DeleteJobSkill(c *gin.Context) {
	successResp = map[string]string{}
	var jc models.JobSkillsMapping
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Update Job Created")
	//var customError models.DbModelError
	defer cancel()
	defer close(jobdb)
	jc.JobID = c.Param("jobID")
	if jc.JobID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find created Job ID in query"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	var err error
	jc.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find id of updating record in query"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string), c.PostForm("id"))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	jc.StakeholderID = ID.(string)
	go func() {
		select {
		case insertJobChan := <-jc.DeleteSkillsByID():
			jobdb <- insertJobChan
		case <-ctx.Done():
			return
		}
	}()
	insertJob := <-jobdb

	if insertJob.ErrTyp != "000" {
		resp := ErrCheck(ctx, insertJob)
		c.Error(insertJob.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	fmt.Printf("\n JC : %+v\n", jc)

	c.JSON(http.StatusOK, DelHCResp{"Job Skills With ID " + c.Param("id") + " has been deleted"})
	return
}
