// Package controllers ...
package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// AddNewProposal ....
func AddNewProposal(c *gin.Context) {
	successResp = map[string]string{}
	//var jc models.FullJobDb

	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Add New University Proposal")
	defer cancel()
	defer close(jobdb)
	var up models.UniversityProposal
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: fmt.Errorf("Cannot able to find the valid headers"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	binding.Validator = &defaultValidator{}

	var err error
	err = c.ShouldBindWith(&up, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))

	if err == nil {
		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}
		form, _ := c.MultipartForm()

		go func() {
			select {
			case insertJobChan := <-up.Insert(ID.(string), form):
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
		c.JSON(http.StatusOK, DelHCResp{"Proposal Added"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3UNVPJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetUnvProposal ...
func GetUnvProposal(c *gin.Context) {
	successResp = map[string]string{}

	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get Jobs Created")
	var up models.UniversityProposal
	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	go func() {
		select {
		case insertJobChan := <-up.GetUnvProposal(ID.(string)):
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

	fmt.Printf("\n GET : %+v\n", up)

	c.JSON(http.StatusOK, up)
	return
}

// DeleteUnvProposalByID ...
func DeleteUnvProposalByID(c *gin.Context, deleteItem string) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Delete University "+deleteItem+" ")
	defer cancel()

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	compareName := strings.ToLower(deleteItem[:1]) + deleteItem[1:] + "ID"
	deleteID := c.Param(compareName)
	deleteQuery := deleteItem + "_UNV_PRP_DEL"
	if deleteID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find %s in params", compareName), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	err := models.DeleteUnvProposalByID(ID.(string), deleteID, deleteQuery)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "INTERNAL ERROR", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, DelHCResp{deleteItem + " deleted"})
	c.Abort()
	return
}

// DeleteBranchByID ...
func DeleteBranchByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "Branch")
}

// DeleteProgramByID ...
func DeleteProgramByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "Program")
}

// DeleteAccredationsByID ...
func DeleteAccredationsByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "Accredations")
}

// DeleteCoesByID ...
func DeleteCoesByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "Coes")
}

// DeleteRankingByID ...
func DeleteRankingByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "Ranking")
}

// DeleteTieupsByID ...
func DeleteTieupsByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "Tieups")
}

// DeleteSpecialOfferingsByID ...
func DeleteSpecialOfferingsByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "SpecialOfferings")
}

// DeleteOtherInfoByID ...
func DeleteOtherInfoByID(c *gin.Context) {
	DeleteUnvProposalByID(c, "OtherInfo")
}
