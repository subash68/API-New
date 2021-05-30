package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s6_Token/models"
)

// AddTokenAllocationForID ...
func AddTokenAllocationForID(c *gin.Context) {

	jobdb := make(chan models.DbModelError, 1)
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Add Token Allocation")

	defer cancel()
	fmt.Printf("\n%v\n", c.Request.Body)
	var tokenAlloc models.TokenAllocationModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&tokenAlloc, binding.Form)
	if err == nil {
		fmt.Printf("\n%+v\n", tokenAlloc)
		go func() {
			select {
			case insertJobChan := <-tokenAlloc.AllocateTokensToID():
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
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.TokenDbResp{"Tokens Added"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetTokenAllocationsOfID ...
func GetTokenAllocationsOfID(c *gin.Context) {
	ctx, ID, _, _ := getFuncReq(c, "Get All Allocations")

	var tokenAllocs models.AllocatedTokens
	err := tokenAllocs.GetAllocateTokensToID(ID)

	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6TKN003", ErrTyp: "INTERNAL SERVER ERROR", Err: fmt.Errorf("Failed to get Token Allocation details %s", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, tokenAllocs.AllocatedTokens)
	return
}

// GetPaymentTxOfID ...
func GetPaymentTxOfID(c *gin.Context) {
	ctx, ID, userType, _ := getFuncReq(c, "Get All Payment transactions")

	var tokenAllocs models.TxTokens
	err := tokenAllocs.GetTransactionsOfTokensOfID(ID, userType)

	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6TKN003", ErrTyp: "INTERNAL SERVER ERROR", Err: fmt.Errorf("Failed to get Token Allocation details %s", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, tokenAllocs.AllocatedTokens)
	return
}
